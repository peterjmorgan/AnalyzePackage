package main

// TODO: import phylum token and project ID from env

import (
	"encoding/json"
	"fmt"
	"github.com/peterjmorgan/go-phylum"
	log "github.com/sirupsen/logrus"
	"net/http"
)

type AnalyzeClient struct {
	Client *phylum.PhylumClient
}

func (a *AnalyzeClient) SendErrorResponse(responseCode int, message string, writer *http.ResponseWriter) {
	(*writer).WriteHeader(responseCode)
	responseMessage := make(map[string]string)
	responseMessage["message"] = message
	jsonResponse, err := json.Marshal(responseMessage)
	if err != nil {
		log.Fatalf("Failed to marshal JSON: %v\n", err)
	}
	_, err = (*writer).Write(jsonResponse)
	if err != nil {
		log.Fatalf("Failed to write response to ResponseWriter: %v\n", err)
	}
}

func (a *AnalyzeClient) AnalyzePackage(writer http.ResponseWriter, req *http.Request) {
	query := req.URL.Query()
	pkgName := query.Get("name")
	pkgVersion := query.Get("version")
	pkgEcosystem := query.Get("ecosystem")

	writer.Header().Set("Content-Type", "application/json")
	if pkgName == "" || pkgVersion == "" || pkgEcosystem == "" {
		a.SendErrorResponse(422, "Missing required GET parameters", &writer)
		return
	}

	packages := make([]phylum.PackageDescriptor, 0)
	packages = append(packages, phylum.PackageDescriptor{
		Name:    pkgName,
		Version: pkgVersion,
		Type:    phylum.PackageType(pkgEcosystem),
	})

	analyzeJobID, err := a.Client.AnalyzeParsedPackages(pkgEcosystem, "afd14d21-eac8-4618-ae83-49d816903bd3", &packages)
	if err != nil {
		log.Errorf("AnalyzeParsedPackages failed: %v\n", err)
		errorMessage := fmt.Sprintf("go-phylum: AnalyzeParsedPackages failed: %v\n", err)
		a.SendErrorResponse(500, errorMessage, &writer)
		return
	}
	packageResponse, _, err := a.Client.GetJobVerbose(analyzeJobID)
	if err != nil {
		fmt.Printf("GetJobVerbose failed: %v\n", err)
		errorMessage := fmt.Sprintf("go-phylum: GetJobVerbose failed: %v\n", err)
		a.SendErrorResponse(500, errorMessage, &writer)
		return
	}
	if packageResponse.Status == "complete" {
		log.Infof("Complete response for %v:%v@%v\n", pkgEcosystem, pkgName, pkgVersion)
		err = json.NewEncoder(writer).Encode(packageResponse)
		if err != nil {
			log.Fatalf("Failed to write Completed response: %v\n", err)
			return
		}
	} else {
		log.Infof("Incomplete response for %v:%v@%v\n", pkgEcosystem, pkgName, pkgVersion)
		err = json.NewEncoder(writer).Encode(&map[string]string{"status": "incomplete"})
		if err != nil {
			log.Fatalf("Failed to write Incompleted response: %v\n", err)
			return
		}
	}
}

func main() {
	a := &AnalyzeClient{Client: phylum.NewClient()}
	http.HandleFunc("/", a.AnalyzePackage)
	log.Fatal(http.ListenAndServe(":3000", nil))
}

# AnalyzePackage
A tiny webservice to submit individual open-source packages to Phylum for analysis via HTTP. 

If the package has already been analyzed by Phylum, the JSON response will be returned. If the package has not been analyzed by Phylum, it is submitted for analysis and a JSON response indicating an `incomplete` status is returned.

**Strictly for internal use only**

Do not put this on the public Internet. This application does not use TLS.

## Overview
AnalyzePackage listens on `0.0.0.0:3000/tcp` for HTTP Get requests. 

To submit a package for analysis, the following `curl` command illustrates the required GET parameters:

`curl http://127.0.0.1:3000/?name=a_cool_package_name&version=1.0.0&ecosystem=npm`

The `ecosystem` parameter can be any of Phylum's 5 supported package registries:
* npm
* pypi
* rubygems
* maven
* nuget

## Requirements
* [Phylum CLI](https://github.com/phylum-dev/cli) installed and authenticated
* A Phylum Project ID. This is easily created:
  1. `phylum project create <project_name>` This will create a .phylum_project file in the current working directory
  2. Extract the `id` field (a GUID) for the projectID: `cat .phylum_project | grep 'id' | cut -d' ' -f2`
  3. Use the resulting GUID for the `-projectID` flag to the application.

## Quick start
1. Clone this repository: `git clone https://github.com/peterjmorgan/AnalyzePackage`
2. Build AnalyzePackage: `go build`
3. Run AnalyzePackage: `./AnalyzePackage -projectID=$PHYLUM_PROJECT_ID`

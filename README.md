# AnalyzePackage
A tiny webservice to submit individual packages to Phylum for analysis via HTTP

## Overview
AnalyzePackage listens on 0.0.0.0:3000/tcp for HTTP Get requests. 

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

## Quick start
1. Clone this repository: `git clone https://github.com/peterjmorgan/AnalyzePackage`
2. Build AnalyzePackage: `go build -o AnalyzePackage`
3. Run AnalyzePackage: `./AnalyzePackage`

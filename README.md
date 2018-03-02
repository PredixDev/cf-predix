[![Build Status](https://travis-ci.org/sks/cf-predix.svg?branch=master)](https://travis-ci.org/sks/cf-predix)

# Predix Plugin
This plugin sets the CF target API to a Predix API

## Usage
```
$ cf predix

1. US West
2. US East
Choose the PoP to set> 1
API endpoint: <URL>

Email>

Password>
```

## Installation
#### If re-installing or updating, first run
```
$ cf uninstall-plugin 'Predix'
```
#### Install using URL
Lookup the URL for the latest binary for your OS from [Releases](https://github.com/PredixDev/cf-predix/releases)

#### Install using File
Download the latest binary for your OS from [Releases](https://github.com/PredixDev/cf-predix/releases)
```
cf install-plugin <PATH TO FILE>
```
#### Install from Source (Should have [Go](https://golang.org/dl/) installed)
```
go get -u github.com/cloudfoundry/cli
go get github.com/PredixDev/cf-predix
cd $GOPATH/src/github.com/PredixDev/cf-predix
go build
cf install-plugin predix
```

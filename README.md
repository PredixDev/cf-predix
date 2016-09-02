# Predix Plugin
This plugin sets the CF target API to a Predix API

## Usage
```
$ cf predix

1. Basic
2. Select
Choose the API to set> 1
Setting api endpoint to <URL>
OK

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
Lookup the URL for the latest binary for your OS from https://github.com/PredixDev/cf-predix/releases
```
$ cf install-plugin https://github.com/PredixDev/cf-predix/releases/download/<VERSION>/predix_<OS>
e.g.
$ cf install-plugin https://github.com/PredixDev/cf-predix/releases/download/<VERSION>/predix_linux64
$ cf install-plugin https://github.com/PredixDev/cf-predix/releases/download/<VERSION>/predix_osx
$ cf install-plugin https://github.com/PredixDev/cf-predix/releases/download/<VERSION>/predix_win64.exe
```
#### Install using File
Download the latest binary for your OS from https://github.com/PredixDev/cf-predix/releases
```
$ cf install-plugin <PATH TO FILE>
```
#### Install from Source (Should have [Go](http://golang.org/dl/) installed)
```
$ go get github.com/cloudfoundry/cli
$ go get github.com/PredixDev/cf-predix
$ cd $GOPATH/src/github.com/PredixDev/cf-predix
$ go build
$ cf install-plugin predix
```

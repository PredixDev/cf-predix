#!/bin/bash

rm -rf bin
mkdir bin
GOOS=darwin GOARCH=amd64 go build -o "bin/osx/predix"
GOOS=linux GOARCH=amd64 go build -o "bin/linux64/predix"
GOOS=windows GOARCH=amd64 go build -o "bin/win64/predix.exe"
GOOS=windows GOARCH=386 go build -o "bin/win32/predix.exe"


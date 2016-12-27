#!/bin/bash -exu

rm -rf bin
mkdir -p bin
GOOS=darwin GOARCH=amd64 go build -o "bin/predix_darwin_64"
GOOS=linux GOARCH=amd64 go build -o "bin/predix_linux_64"
GOOS=windows GOARCH=amd64 go build -o "bin/predix_win_64.exe"
GOOS=windows GOARCH=386 go build -o "bin/predix_win_32.exe"

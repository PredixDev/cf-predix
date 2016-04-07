#!/bin/bash

if [[ "$(which gox)X" == "X" ]]; then
    echo "Please install gox. https://github.com/mitchellh/gox#readme"
    exit 1
fi

rm -rf bin
mkdir bin
gox -os darwin -arch amd64 --output="bin/osx/predix"
gox -os linux -arch amd64 --output="bin/linux64/predix"
gox -os windows -arch amd64 --output="bin/win64/predix"

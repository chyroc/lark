#!/bin/bash

set -e

module_name=$(cat go.mod | grep module | cut -d ' ' -f 2-2)
echo "module_name is $module_name"

export LOCAL_TEST=true

gotestsum --rerun-fails --packages="./..." -- -coverprofile=coverage.txt -coverpkg=./... -parallel 1 -p 1 -count=1 -gcflags=-l -v ./...

# go tool cover -func=coverage.txt -o coverage.txt

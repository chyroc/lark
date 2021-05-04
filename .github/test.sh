#!/bin/bash

set -e

module_name=$(cat go.mod | grep module | cut -d ' ' -f 2-2)
echo "module_name is $module_name"

GO111MODULE=on echo 'mode: atomic' > coverage.txt && \
  go list ./... | \
    xargs -n1 -I{} sh -c 'LOCAL_TEST=true go test -covermode=atomic -coverprofile=coverage.tmp -coverpkg=./... -parallel 1 -p 1 -count=1 -gcflags=-l {} && tail -n +2 coverage.tmp >> coverage.txt' && \
  rm coverage.tmp

# go tool cover -func=coverage.txt -o coverage.txt

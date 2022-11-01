#!/bin/bash

cd /go/src/app
go get .
curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s -- -b $(go env GOPATH)/bin
export PATH=$PATH:$(go env GOPATH)/bin
export air_wd=/go/src/app
air
# go run .
# ./shorter-url

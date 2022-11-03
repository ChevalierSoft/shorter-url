#!/bin/bash

cd /go/src/app
go get .
go install github.com/swaggo/swag/cmd/swag	# install swagger
swag init
curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s -- -b $(go env GOPATH)/bin	# install air
export PATH=$PATH:$(go env GOPATH)/bin
export air_wd=/go/src/app
air

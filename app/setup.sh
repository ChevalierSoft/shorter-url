#!/bin/bash

export PATH=$PATH:/go/bin
cd /go/src/app

go get -u github.com/swaggo/swag/cmd/swag
go install github.com/swaggo/swag/cmd/swag	# install swagger
swag init

curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s -- -b $(go env GOPATH)/bin	# install air
export air_wd=/go/src/app
air

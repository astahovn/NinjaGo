#!/usr/bin/env bash
export GOPATH=`pwd`

#go get github.com/gin-gonic/gin
#go get github.com/lib/pq
#go get -u github.com/jinzhu/gorm
#go install routes
#export PATH=$PATH:$(go env GOPATH)/bin

go run src/server/main.go
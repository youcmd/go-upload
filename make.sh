#!/bin/bash
export CGO_ENABLED=0
go build
GOOS=windows GOARCH=amd64 go build

#!/bin/bash
export CGO_ENABLED=0
go build -o go-ul
GOOS=windows GOARCH=amd64 go build -o go-ul.exe

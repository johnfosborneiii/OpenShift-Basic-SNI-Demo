#!/bin/bash

GOPATH=$PWD go build server.go

docker build -t pweil/tls-server .

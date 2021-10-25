#!/usr/bin/env bash 
set -xe

# install packages and dependencies
cat tools.go | grep _ | awk -F'"' '{print $2}' | xargs -tI % go install %\n

# build command
GOARCH=amd64 GOOS=linux go build -o bin/application application.go
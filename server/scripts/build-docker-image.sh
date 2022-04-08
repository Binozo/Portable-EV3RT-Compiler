#!/bin/bash

cd ..
echo "Building server executable"
cd "cmd/main"
env GOOS=linux GOARCH=amd64 go build -o main

echo "Building docker image"
cd "../../"
docker build -t binozoworks/portable-ev3-compiler:latest .
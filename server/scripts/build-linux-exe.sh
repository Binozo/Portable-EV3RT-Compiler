#!/bin/bash

cd ..
echo "Building server executable"
cd "cmd/main"
env GOOS=linux GOARCH=amd64 go build -o main

mv main ../../bin/main
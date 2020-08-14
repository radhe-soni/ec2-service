#! /bin/sh
rm -r build/
mkdir build
GOOS=linux go build -o build/ec2Service
GOOS=linux go build -ldflags="-s -w" -o build/ec2Service-trimmed
GOOS=windows go build -o build/ec2Service.exe
GOOS=windows go build -ldflags="-s -w" -o build/ec2Service-trimmed.exe
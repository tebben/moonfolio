#!/bin/bash

mkdir -p ../bin
mkdir -p ../bin/win64
mkdir -p ../bin/linux64
mkdir -p ../bin/darwin64
mkdir -p ../bin/linuxARMv7

export GOOS=windows
export GOARCH=amd64
go build -o ../bin/win64/moonfolio.exe ../../main.go
cp ../../config.json ../bin/win64/config.json
echo "Built application for Windows/amd64"

export GOOS=linux
export GOARCH=amd64
go build -o ../bin/linux64/moonfolio ../../main.go
cp ../../config.json ../bin/linux64/config.json
echo "Built application for Linux/amd64"

export GOOS=darwin
export GOARCH=amd64
go build -o ../bin/darwin64/moonfolio ../../main.go
cp ../../config.json ../bin/darwin64/config.json
echo "Built application for Darwin/amd64"

export GOOS=linux
export GOARCH=arm
export GOARM=7
go build -o ../bin/linuxARMv7/moonfolio ../../main.go
cp ../../config.json ../bin/linuxARMv7/config.json
echo "Built application for Linux/ARMv7"
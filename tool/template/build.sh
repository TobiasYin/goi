#!/bin/bash
cd {{PROJECT_PATH}}
rm -rf output 2>  /dev/null
mkdir output  2> /dev/null
cp -r asset output/ 2> /dev/null
cp index.html output 2> /dev/null
go mod tidy 2> /dev/null
go get github.com/TobiasYin/goi@master 2> /dev/null
GOARCH=wasm GOOS=js go build -o output/main.wasm main.go
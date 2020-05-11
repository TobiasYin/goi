#!/bin/bash
export GOARCH=wasm
export GOOS=js
cd {{PROJECT_PATH}}
rm -rf output 2>  /dev/null
mkdir output  2> /dev/null
cp -r asset output/ 2> /dev/null
cp index.html output 2> /dev/null
go get github.com/TobiasYin/goi@latest 2> /dev/null
go build -o output/main.wasm main.go
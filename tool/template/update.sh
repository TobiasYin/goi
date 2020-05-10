#!/bin/bash
export GOARCH=wasm
export GOOS=js
cd {{PROJECT_PATH}}
if [ ! -d  "output" ]; then
  echo "Create Static File"
  mkdir output  2> /dev/null
  cp -r asset output/ 2> /dev/null
  cp index.html output 2> /dev/null
else
  echo "Static Exist, Compile Only"
fi
go get github.com/TobiasYin/goi@latest 2> /dev/null
GOARCH=wasm GOOS=js go build -o output/main.wasm main.go

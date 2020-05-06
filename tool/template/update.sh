#!/bin/bash
cd {{PROJECT_PATH}}
if [ ! -d  "output" ]; then
  echo "Create Static File"
  mkdir output  2> /dev/null
  cp -r assert output/ 2> /dev/null
  cp index.html output 2> /dev/null
else
  echo "Static Exist, Compile Only"
fi
go mod tidy 2> /dev/null
go get -d github.com/TobiasYin/go_web_ui 2> /dev/null
GOARCH=wasm GOOS=js go build -o output/main.wasm main.go

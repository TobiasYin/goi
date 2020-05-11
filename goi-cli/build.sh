#!/usr/bin/env bash

python3 load_template.py
GOARCH=amd64 GOOS=darwin go build -o goi_tool main.go
GOARCH=amd64 GOOS=windows go build -o goi_tool.exe main.go
GOARCH=amd64 GOOS=linux go build -o goi_tool.out main.go
rm inline/init.go
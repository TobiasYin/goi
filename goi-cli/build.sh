#!/usr/bin/env bash

python3 load_template.py
GOARCH=amd64 GOOS=darwin go build -o goi-cli main.go
GOARCH=amd64 GOOS=windows go build -o goi-cli.exe main.go
GOARCH=amd64 GOOS=linux go build -o goi-cli.out main.go
#!/usr/bin/env bash

python3 load_template.py
go build -o go_web_ui_tool main.go
GOARCH=amd64 GOOS=windows go build -o go_web_ui_tool.exe main.go
GOARCH=amd64 GOOS=linux go build -o go_web_ui_tool.out main.go
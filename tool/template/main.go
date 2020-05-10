package main

import (
	"github.com/TobiasYin/goi/logs"
	"github.com/TobiasYin/goi/node"
	pages "github.com/TobiasYin/goi/tool/template/page"
	_ "github.com/TobiasYin/goi/tool/template/urls"
)

func main() {
	logs.SetLogLevel(logs.Debug)
	node.SetMaxFPS(100)
	node.NewApp(pages.GetMainPage())
}

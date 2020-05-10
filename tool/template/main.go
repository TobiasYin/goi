package main

import (
	"github.com/TobiasYin/goi"
	"github.com/TobiasYin/goi/logs"
	pages "github.com/TobiasYin/goi/tool/template/page"
	_ "github.com/TobiasYin/goi/tool/template/urls"
)

func main() {
	logs.SetLogLevel(logs.Debug)
	goi.SetMaxFPS(100)
	goi.NewApp(pages.GetMainPage())
}

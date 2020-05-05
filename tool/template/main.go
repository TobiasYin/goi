package main

import (
	"github.com/TobiasYin/go_web_ui/node"
	pages "github.com/TobiasYin/go_web_ui/tool/template/page"
	_ "github.com/TobiasYin/go_web_ui/tool/template/urls"
)

func main() {
	node.NewApp(pages.GetMainPage())
}

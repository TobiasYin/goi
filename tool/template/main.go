package main

import (
	"github.com/TobiasYin/go_web_ui/node"
	pages "github.com/TobiasYin/go_web_ui/tool/template/page"
)

func main() {
	node.NewApp(pages.GetMainPage())
}

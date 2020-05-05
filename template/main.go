package main

import (
	pages "github.com/TobiasYin/go_web_ui/example/page"
	"github.com/TobiasYin/go_web_ui/node"
)

func main() {
	node.RegisterRoute("/image", pages.IntoImage)
	node.RegisterRoute("/image/new", pages.NewRelative)
	node.NewApp(pages.GetMainPage())
}

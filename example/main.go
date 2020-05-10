package main

import (
	pages "github.com/TobiasYin/goi/example/page"
	"github.com/TobiasYin/goi/node"
)

func main() {
	node.RegisterRoute("/image", pages.IntoImage)
	node.RegisterRoute("/image/new", pages.NewRelative)
	node.NewApp(pages.GetMainPage())
}

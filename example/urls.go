package main

import (
	pages "github.com/TobiasYin/go_web_ui/example/page"
	"github.com/TobiasYin/go_web_ui/node"
)

func init() {
	node.RegisterRoute("/image", pages.IntoImage)
}

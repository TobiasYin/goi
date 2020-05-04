package main

import (
	pages "github.com/TobiasYin/go_web_ui/example/page"
	"github.com/TobiasYin/go_web_ui/node"
)

func main() {
	c := make(chan struct{})
	node.RegisterRoute("/image", pages.IntoImage)
	node.NewApp(pages.GetMainPage())
	<-c
}

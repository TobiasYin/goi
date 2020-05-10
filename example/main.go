package main

import (
	"github.com/TobiasYin/goi"
	pages "github.com/TobiasYin/goi/example/page"
)

func main() {
	goi.RegisterRoute("/image", pages.IntoImage)
	goi.RegisterRoute("/image/new", pages.NewRelative)
	goi.NewApp(pages.GetMainPage())
}

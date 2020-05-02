package node

import (
	"fmt"
	"github.com/TobiasYin/go_web_ui/dom"
	"time"
)

func NewApp(page *Page) {
	stack.Add(page)
	FlashApp()
}

func FlashApp() {
	start := time.Now()
	app := dom.Dom.GetElementById("app")
	children := app.GetChildren()
	for _, child := range children {
		app.RemoveChild(child)
	}
	app.AppendChild(stack.Top().pack())
	end := time.Now()
	fmt.Printf("Re Render Page, Using: %v\n", end.Sub(start))
}

func PushToPage(page *Page) {
	stack.Add(page)
	FlashApp()
}

func BackToLastPage() {
	stack.Pop()
	FlashApp()
}

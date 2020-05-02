package node

import "github.com/TobiasYin/go_web_ui/dom"

func NewApp(page *Page) {
	stack.Add(page)
	FlashApp()
}

func FlashApp() {
	app := dom.Dom.GetElementById("app")
	children := app.GetChildren()
	for _, child := range children{
		app.RemoveChild(child)
	}
	app.AppendChild(stack.Top().pack())
}

func PushToPage(page *Page)  {
	stack.Add(page)
	FlashApp()
}

func BackToLastPage() {
	stack.Pop()
	FlashApp()
}

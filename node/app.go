package node

import (
	"fmt"
	"github.com/TobiasYin/go_web_ui/vdom"
	"time"
)

func NewApp(page *Page) {
	stack.Add(page)
	FlashApp()
}

func FlashApp() {
	start := time.Now()
	top := stack.Top()
	d := top.pack()
	vdom.MergeTwoTree(&d, top.oldDom)
	top.oldDom = &d
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

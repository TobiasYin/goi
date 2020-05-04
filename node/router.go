package node

import (
	"fmt"
	dom "github.com/TobiasYin/go_web_ui/vdom"
)

type pageStack struct {
	stack []*Page
	size  int
}

var (
	router map[string]PathPage
	stack  pageStack
)

func init() {
	router = make(map[string]PathPage)
}

func RegisterRoute(path string, page PathPage) {
	if len(path) == 0{
		path = "/"
	}else if path[0] != '/'{
		path = "/" + path
	}
	router[path] = page
}

func (p *pageStack) Top() *Page {
	return p.stack[p.size-1]
}

func (p *pageStack) Pop() *Page {
	p.size--
	top := p.stack[p.size]
	p.stack[p.size] = nil
	p.stack = p.stack[0:p.size]
	return top
}

func (p *pageStack) Add(page *Page) {
	p.stack = append(p.stack, page)
	p.size++
}

func (p *pageStack) pack() dom.JsDomElement {
	return p.Top().pack()
}


func PushToPage(page *Page) {
	stack.Add(page)
	if page.oldDom != nil {
		dom.Display(page.oldDom)
	}
	FlashApp()
}

func BackToLastPage() {
	stack.Pop()
	dom.Display(stack.Top().oldDom)
	FlashApp()
}

func PushByPath(path string, arg map[string]interface{}) error {
	page, ok := router[path]
	if !ok {
		return fmt.Errorf("unkonw page")
	}
	PushToPage(page(arg).GetPage())
	return nil
}

type PageGetter interface {
	GetPage() *Page
}

type PathPage func(map[string]interface{}) PageGetter
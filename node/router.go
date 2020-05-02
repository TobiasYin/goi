package node

import "github.com/TobiasYin/go_web_ui/dom"

type pageStack struct {
	stack []*Page
	size  int
}

var (
	router map[string]*Page
	stack  pageStack
)

func init() {
	router = make(map[string]*Page)
}

func RegisterRoute(path string, page *Page) {
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

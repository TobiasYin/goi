package node

import (
	"context"
	"fmt"
	"github.com/TobiasYin/go_web_ui/dom"
	"runtime"
)

type StateArea interface {
	SetState(f func())
	setNode(node Node)
	setStateToFather()
	doSetState()
	getContext() context.Context
}

type Context struct {
	Context context.Context
	GetNode func() Node
	node    Node
}

func NewContext(area StateArea) *Context {
	newCtx := context.WithValue(area.getContext(), "father", area)
	return &Context{
		Context: newCtx,
	}
}

func (c *Context) SetState(f func()) {
	f()
	c.doSetState()
}

func (c *Context) doSetState() {
	c.setNode(c.GetNode())
	c.setStateToFather()
}

func (c *Context) setNode(node Node) {
	c.node = node
}

func (c *Context) getContext() context.Context {
	return c.Context
}

func (c *Context) setStateToFather() {
	father, ok := c.Context.Value("father").(StateArea)
	if ok {
		father.doSetState()
	}
}

func (c Context) pack() dom.JsDomElement {
	if c.node == nil {
		c.node = c.GetNode()

	}
	return c.node.pack()
}

type Page struct {
	Context
}

func NewPage() *Page {
	return &Page{Context{Context: context.Background()}}
}

func (p *Page) SetState(f func()) {
	f()
	p.doSetState()
}

func (p *Page) doSetState() {
	p.setNode(p.GetNode())
	FlashApp()
}

var contexts map[string]*Context

func init() {
	contexts = make(map[string]*Context)
}

func ContextKeepWrapperWithKey(f func(StateArea) *Context, key string) func(StateArea) *Context {
	return func(c StateArea) *Context {
		if res, ok := contexts[key]; ok {
			return res
		}
		res := f(c)
		contexts[key] = res
		return res
	}
}

func ContextKeepWrapper(f func(StateArea) *Context) func(StateArea) *Context {
	funcName, file, line, _ := runtime.Caller(0)
	key := fmt.Sprintf("%v,%v,%d", funcName, file, line)
	return ContextKeepWrapperWithKey(f, key)
}

package node

import (
	"context"
	"fmt"
	"github.com/TobiasYin/go_web_ui/dom"
	"runtime"
)

type ComponentConstruct func(StateArea) *Context

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
	newCtx = context.WithValue(newCtx, "fatherState", make(map[string]*Context))
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

func (c *Context) StatefulChild(f ComponentConstruct) Node {
	return ContextKeepWrapper(*c, f)(c)
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
	ctx := context.Background()
	ctx = context.WithValue(ctx, "fatherState", make(map[string]*Context))

	return &Page{Context{Context: ctx}}
}

func (p *Page) SetState(f func()) {
	f()
	p.doSetState()
}

func (p *Page) doSetState() {
	p.setNode(p.GetNode())
	FlashApp()
}

func (p *Page) StatefulChild(f ComponentConstruct) Node {
	return ContextKeepWrapper(p.Context, f)(p)
}

func ContextKeepWrapperWithKey(father Context, f ComponentConstruct, key string) ComponentConstruct {
	contexts, ok := father.Context.Value("fatherState").(map[string]*Context)
	if !ok {
		return f
	}
	return func(c StateArea) *Context {
		if res, ok := contexts[key]; ok {
			return res
		}
		res := f(c)
		contexts[key] = res
		return res
	}
}

func ContextKeepWrapper(father Context, f ComponentConstruct) ComponentConstruct {
	funcName, file, line, _ := runtime.Caller(2)
	key := fmt.Sprintf("%v,%v,%d", funcName, file, line)
	return ContextKeepWrapperWithKey(father, f, key)
}

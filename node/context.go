package node

import (
	"context"
	"fmt"
	"github.com/TobiasYin/go_web_ui/dom"
	"runtime"
)

type ComponentCreator func(StateArea) *Context
type ComponentConstructor func(*Context) Node
type Component func() ComponentConstructor

type StateArea interface {
	SetState(f func())
	setNode(node Node)
	setStateToFather()
	doSetState()
	getContext() context.Context
}

type Context struct {
	Context context.Context
	GetNode func(*Context) Node
	node    Node
	isPage  bool
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
	if c.isPage {
		c.doPageSetState()
		return
	}
	c.setNode(c.GetNode(c))
	c.setStateToFather()
}

func (c *Context) StatefulChild(f Component) Node {
	return ContextKeepWrapper(c, ComponentConstructWrapper(f))(c)
}

func (c *Context) StatelessChild(f ComponentConstructor) Node {
	return f(c)
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
		c.node = c.GetNode(&c)

	}
	return c.node.pack()
}

func (c *Context) doPageSetState() {
	c.setNode(c.GetNode(c))
	FlashApp()
}

type Page struct {
	Context
}

func NewPageEmpty() *Page {
	ctx := context.Background()
	ctx = context.WithValue(ctx, "fatherState", make(map[string]*Context))
	page := Page{Context{Context: ctx, isPage: true}}
	return &page
}

func NewPage(getNode ComponentConstructor) *Page {
	page := NewPageEmpty()
	page.GetNode = getNode
	return page
}

func ContextKeepWrapperWithKey(father *Context, f ComponentCreator, key string) ComponentCreator {
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

func ContextKeepWrapper(father *Context, f ComponentCreator) ComponentCreator {
	funcName, file, line, _ := runtime.Caller(2)
	key := fmt.Sprintf("%v,%v,%d", funcName, file, line)
	return ContextKeepWrapperWithKey(father, f, key)
}

func ComponentConstructWrapper(f Component) ComponentCreator {
	return func(area StateArea) *Context {
		newCtx := NewContext(area)
		c := f()
		newCtx.GetNode = c
		return newCtx
	}
}

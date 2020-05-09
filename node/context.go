package node

import (
	"context"
	"fmt"
	"github.com/TobiasYin/go_web_ui/logs"
	dom "github.com/TobiasYin/go_web_ui/vdom"
	"runtime"
)

type ComponentCreator func(StateArea) *Context
type ComponentConstructor func(*Context) Widget
type ComponentFunc func() ComponentConstructor

type StateArea interface {
	SetState(f func())
	setNode(node Widget)
	setStateToFather()
	doSetState()
	getContext() Context
}

type StatefulWidget interface {
	Widget
	GetConstructor() ComponentConstructor
	GetKey() string
}

type StatelessWidget interface {
	Widget
	GetNode(*Context) Widget
}

func PackStateful(sf StatefulWidget, ctx Context) Node {
	return ContextKeepWrapperWithKey(&ctx, ComponentConstructWrapper(sf.GetConstructor), sf.GetKey())(&ctx)
}

func PackStateless(sl StatelessWidget, ctx Context) Node {
	return sl.GetNode(&ctx).Pack(ctx)
}

type Context struct {
	Context context.Context
	GetNode func(*Context) Widget
	node    Widget
	isPage  bool
	oldTree *dom.JsDomElement
	tree    *dom.JsDomElement
}

func NewContext(area StateArea) *Context {
	newCtx := context.WithValue(area.getContext().Context, "father", area)
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
	c.oldTree = c.tree
	c.setNode(c.GetNode(c))
	c.pack()
	rerenderTree(c.tree, c.oldTree)
}

func (c *Context) Pack(ctx Context) Node {
	return c.GetNode(c).Pack(*c)
}

func (c *Context) StatefulChild(sc StatefulWidget) Node {
	return ContextKeepWrapper(c, ComponentConstructWrapper(sc.GetConstructor))(c)
}

func (c *Context) StatelessChild(sc StatelessWidget) Node {
	return sc.GetNode(c).Pack(*c)
}

func (c *Context) setNode(node Widget) {
	c.node = node
}

func (c *Context) getContext() Context {
	return *c
}

func (c *Context) setStateToFather() {
	father, ok := c.Context.Value("father").(StateArea)
	if ok {
		father.doSetState()
	}
}

func (c *Context) pack() dom.JsDomElement {
	if c.node == nil {
		c.node = c.GetNode(c)
	}
	tree := c.node.Pack(*c).pack()
	c.tree = &tree
	return tree
}

func (c *Context) doPageSetState() {
	c.setNode(c.GetNode(c))
	FlashApp()
}

type Page struct {
	path string
	Context
	Title string
}

func NewPageEmpty() *Page {
	ctx := context.Background()
	ctx = context.WithValue(ctx, "fatherState", make(map[string]*Context))
	page := Page{Context: Context{Context: ctx, isPage: true}}
	return &page
}

func NewPage(title string, getNode ComponentConstructor) *Page {
	page := NewPageEmpty()
	page.GetNode = getNode
	page.Title = title
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
	funcName, file, line, _ := runtime.Caller(3)
	key := fmt.Sprintf("%v,%v,%d", funcName, file, line)
	return ContextKeepWrapperWithKey(father, f, key)
}

//根据调用位置生成一个key，请勿在循环中使用！！！因为调用位置相同，循环中可使用更有代表性的key。
func GenerateKeyWithCallLine() string {
	funcName, file, line, _ := runtime.Caller(1)
	key := fmt.Sprintf("%v,%v,%d", funcName, file, line)
	return key
}

func ComponentConstructWrapper(f ComponentFunc) ComponentCreator {
	return func(area StateArea) *Context {
		newCtx := NewContext(area)
		c := f()
		newCtx.GetNode = c
		return newCtx
	}
}

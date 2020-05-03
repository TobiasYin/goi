package node

import (
	"github.com/TobiasYin/go_web_ui/dom"
	"strings"
)

type Node interface {
	pack() dom.JsDomElement
}

type packAble interface {
	getParam() Params
	packChildAble
}

type packChildAble interface {
	getChildren() []Node
}

func packChildren(able packChildAble, ele *dom.JsDomElement) {
	for _, c := range able.getChildren() {
		if c == nil{
			continue
		}
		//Pack
		v := c.pack()
		ele.AppendChild(v)
	}
}

func pack(able packAble, name string) dom.JsDomElement {
	p := able.getParam()
	e := p.packWithName(name)
	packChildren(able, &e)
	return e
}

type Params struct {
	Style       Style
	Class       []string
	Value       string
	Placeholder string
	OnClick     EventCallBack
	OnInput     EventCallBack
	OnChange    EventCallBack
	OnFocus     EventCallBack
	OnBlur      EventCallBack
	OnKeyDown   EventCallBack
	OnKeyUp     EventCallBack
}

func (e Params) getParam() Params {
	return e
}

func (e Params) packWithName(name string) dom.JsDomElement {
	ele := dom.Dom.CreateElement(name)
	style := e.Style.packStyle()
	if style != "" {
		ele.Set("style", style)
	}
	if len(e.Class) != 0 {
		ele.Set("className", strings.Join(e.Class, " "))
	}
	if e.OnClick != nil {
		ele.Set("onclick", wrapEventCallBack(e.OnClick))
	}
	if e.OnInput != nil {
		ele.Set("oninput", wrapEventCallBack(e.OnInput))
	}
	if e.OnChange != nil {
		ele.Set("onchange", wrapEventCallBack(e.OnChange))
	}
	if e.OnFocus != nil {
		ele.Set("onfocus", wrapEventCallBack(e.OnFocus))
	}
	if e.OnBlur != nil {
		ele.Set("onblur", wrapEventCallBack(e.OnBlur))
	}
	if e.OnKeyDown != nil {
		ele.Set("onkeydown", wrapEventCallBack(e.OnKeyDown))
	}
	if e.OnKeyUp != nil {
		ele.Set("onkeyup", wrapEventCallBack(e.OnKeyUp))
	}
	return ele
}

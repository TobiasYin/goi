package node

import (
	"github.com/TobiasYin/go_web_ui/dom"
	"strings"
)

type Node interface {
	pack() dom.JsDomElement
}

type Element struct {
	Style       *Style
	Class       []string
	Children    []Node
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

type SingleChildElement struct {
	Style       *Style
	Class       []string
	Child       Node
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

func (e Element) pack() dom.JsDomElement {
	return e.packWithName("div")
}

func (e Element) packWithName(name string) dom.JsDomElement {
	ele := dom.Dom.CreateElement(name)
	if e.Style != nil {
		ele.Set("style", e.Style.packStyle())
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
	for _, c := range e.Children {
		//Pack
		v := c.pack()
		ele.AppendChild(v)
	}
	return ele
}

func (e SingleChildElement) pack() dom.JsDomElement {
	return e.packWithName("div")
}

func (e SingleChildElement) packWithName(name string) dom.JsDomElement {
	ele := dom.Dom.CreateElement(name)
	if e.Style != nil {
		ele.Set("style", e.Style.packStyle())
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
	if e.Child != nil {
		ele.AppendChild(e.Child.pack())
	}
	return ele
}


//func test() {
//	page := Element{
//		Children: []Node{
//			Element{
//				Children: []Node{
//					Text{
//						Content: "aaa",
//					},
//					Text{
//						Content: "bbb",
//					},
//				},
//			},
//		},
//	}
//}

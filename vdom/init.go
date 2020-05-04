package vdom

import (
	"github.com/TobiasYin/go_web_ui/dom"
	"syscall/js"
)

var Root = dom.Dom.GetElementById("app")
var Dom = Document{}

func equal(e1 JsDomElement, e2 JsDomElement) bool {
	if e1.Tag != e2.Tag {
		return false
	}
	if e1.Content != e2.Content {
		return false
	}
	if len(e1.Value) != len(e2.Value) {
		return false
	}
	if len(e1.Children) != len(e2.Children) {
		return false
	}
	for k, v := range e1.Value {
		if _, ok := v.v.(js.Func); ok {
			continue
		}
		if e2.Value[k] != v {
			return false
		}
	}
	return true
}

type Value struct {
	v interface{}
}

func (v Value) String() string {
	if res, ok := v.v.(string); ok {
		return res
	}
	return ""
}

type JsDomElement struct {
	Tag      string
	Content  string
	Value    map[string]Value
	Children []*JsDomElement
	RealDom  *dom.JsDomElement
}

func NewJsDomElement(tag string) JsDomElement {
	return JsDomElement{Tag: tag, Value: make(map[string]Value)}
}

func (e JsDomElement) RemoveChild(c JsDomElement) {
	for i := 0; i < len(e.Children); i++ {
		if equal(c, *e.Children[i]) {
			e.Children = append(e.Children[:i], e.Children[i+1:]...)
			return
		}
	}
}

func (e JsDomElement) RemoveChildAt(index int) {
	e.Children = append(e.Children[:index], e.Children[index+1:]...)
}

func (e JsDomElement) Get(key string) Value {
	res, _ := e.Value[key]
	return res
}

func (e *JsDomElement) Set(key string, value interface{}) {
	e.Value[key] = Value{value}
}

func (e JsDomElement) GetValue() string {
	v, ok := e.Value["value"]
	res := ""
	if ok {
		res = v.String()
	}
	return res
}

func (e *JsDomElement) SetValue(v interface{}) {
	e.Set("value", v)
}

func (e *JsDomElement) AppendChild(n JsDomElement) {
	e.Children = append(e.Children, &n)
}

type Document struct {
	JsDomElement
}

func (d Document) CreateElement(name string) JsDomElement {
	return NewJsDomElement(name)
}
func (d Document) CreateTextNode(name string) JsDomElement {
	return JsDomElement{
		Tag:     "text",
		Content: name,
	}
}

func (e *JsDomElement) GetRealDom() dom.JsDomElement {
	if e.RealDom != nil {
		return *e.RealDom
	}
	if e.Tag == "text" {
		res := dom.Dom.CreateTextNode(e.Content)
		e.RealDom = &res
		return res
	}
	ele := dom.Dom.CreateElement(e.Tag)
	for k, v := range e.Value {
		ele.Set(k, v.v)
	}
	for _, i := range e.Children {
		ele.AppendChild(i.GetRealDom())
	}
	e.RealDom = &ele
	return ele
}

func MergeTwoTree(newTree *JsDomElement, oldTree *JsDomElement) {
	rebuild := func() {
		Display(newTree)
	}
	if oldTree == nil {
		rebuild()
	} else {
		if !equal(*newTree, *oldTree) {
			rebuild()
			return
		}
		mergeTwoTree(newTree, oldTree)
	}
}

func Display(tree *JsDomElement)  {
	children := Root.GetChildren()
	for _, child := range children {
		Root.RemoveChild(child)
	}
	Root.AppendChild(tree.GetRealDom())
}

func mergeTwoTree(newTree *JsDomElement, oldTree *JsDomElement) {
	length := len(newTree.Children)
	for k, v := range newTree.Value {
		if f, ok := v.v.(EventCallBack); ok {
			oldTree.RealDom.Set(k, WrapEventCallBack(f))
		}
	}
	for i, c := range newTree.Children {
		oc := oldTree.Children[i]
		if !equal(*oc, *c) {
			oldTree.RealDom.RemoveChild(oc.GetRealDom())
			if i == length-1 {
				oldTree.RealDom.AppendChild(c.GetRealDom())
			} else {
				oldTree.RealDom.InsertBefore(c.GetRealDom(), oldTree.Children[i+1].GetRealDom())
			}
		} else {
			c.RealDom = oc.RealDom
			mergeTwoTree(c, oc)
		}
	}
}

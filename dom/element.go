package dom

import "syscall/js"

type JsDomElement struct {
	Value
}

func NewJsDomElement(v js.Value) JsDomElement {
	return JsDomElement{
		Value{
			v,
		},
	}
}

func (e JsDomElement) GetElementById(name string) JsDomElement {
	return NewJsDomElement(e.call(name))
}

func (e JsDomElement) getElements(name string) []JsDomElement {
	res := e.callWithTwoSkip(name)
	es := make([]JsDomElement, res.Length())
	for i := 0; i < res.Length(); i++ {
		es[i] = NewJsDomElement(res.Index(i))
	}
	return es
}

func (e JsDomElement) GetElementsByClassName(name string) []JsDomElement {
	return e.getElements(name)
}

func (e JsDomElement) GetElementsByName(name string) []JsDomElement {
	return e.getElements(name)
}
func (e JsDomElement) GetElementsByTagName(name string) []JsDomElement {
	return e.getElements(name)
}
func (e JsDomElement) GetElementsByTagNameNS(name string) []JsDomElement {
	return e.getElements(name)
}

func (e JsDomElement) RemoveChild(c JsDomElement) {
	e.call(c.Value)
}

func (e JsDomElement) GetChildren() []JsDomElement {
	res := e.Get("children")
	es := make([]JsDomElement, res.Length())
	for i := 0; i < res.Length(); i++ {
		es[i] = NewJsDomElement(res.Index(i))
	}
	return es
}

func (e JsDomElement) GetValue() js.Value {
	return e.Get("value")
}

func (e JsDomElement) SetValue(v interface{}) {
	e.Set("value", v)
}

func (e JsDomElement) AppendChild(n JsDomElement) {
	e.call(n.Value)
}

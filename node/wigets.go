package node

import "github.com/TobiasYin/go_web_ui/dom"

type Div struct {
	Params
	Children []Node
}

type Input struct {
	Params
}

type Button struct {
	Params
	Child Node
}

type BR struct{}

func (br BR) pack() dom.JsDomElement {
	return dom.Dom.CreateElement("br")
}

func (d Div) getChildren() []Node {
	return d.Children
}

func (d Div) pack() dom.JsDomElement {
	return pack(d, "div")
}

func (i Input) getChildren() []Node {
	return []Node{}
}

func (i Input) pack() dom.JsDomElement {
	return pack(i, "input")
}

func (b Button) getChildren() []Node {
	return []Node{b.Child}
}

func (b Button) pack() dom.JsDomElement {
	return pack(b, "button")
}

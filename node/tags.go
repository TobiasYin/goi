package node

import "github.com/TobiasYin/go_web_ui/dom"

type Div Element

type Input SingleChildElement

type Button SingleChildElement

func (d Div) pack() dom.JsDomElement {
	return Element(d).packWithName("div")
}

func (i Input) pack() dom.JsDomElement {
	return SingleChildElement(i).packWithName("input")
}

func (b Button) pack() dom.JsDomElement {
	return SingleChildElement(b).packWithName("button")
}

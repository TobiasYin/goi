package node

import dom "github.com/TobiasYin/go_web_ui/vdom"

type Image struct {
	Params
	Src string
}

func (i Image) pack() dom.JsDomElement {
	e := pack(i, "img")
	e.Set("src", i.Src)
	return e
}

func (i Image) getChildren() []Node {
	return []Node{}
}

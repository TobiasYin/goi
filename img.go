package goi

import dom "github.com/TobiasYin/goi/vdom"

type Image struct {
	Params
	Src string
	_context
}

func (i Image) pack() dom.JsDomElement {
	e := pack(i, "img", i.getContext())
	e.Set("src", i.Src)
	return e
}

func (i Image) getChildren() []Widget {
	return []Widget{}
}

func (i Image) Pack(ctx Context) Node {
	i._context.Context = ctx
	return i
}

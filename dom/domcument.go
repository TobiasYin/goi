package dom

type Document struct {
	JsDomElement
}

func (d Document) CreateElement(name string) JsDomElement {
	return NewJsDomElement(d.call(name))
}
func (d Document) CreateTextNode(name string) JsDomElement {
	return NewJsDomElement(d.call(name))
}

func (d Document) SetTitle(title string) {
	d.Set("title", title)
}

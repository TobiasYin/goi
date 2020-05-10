package component

import (
	"fmt"
	"github.com/TobiasYin/goi/logs"
	"github.com/TobiasYin/goi/node"
	"github.com/TobiasYin/goi/node/color"
	"strconv"
)

type StatelessDemo struct {
	Value string
}

func (sc StatelessDemo) GetWidget(context *node.Context) node.Widget {
	return node.Block{
		Children: []node.Widget{
			node.Text{
				Content: sc.Value + " Stateless",
			},
			node.BR{},
		},
	}
}

func (sc StatelessDemo) Pack(context node.Context) node.Node {
	return node.PackStateless(sc, context)
}

type StatefulDemo struct {
	Key   string
	Value string
	Child node.Widget
	Size  int
}

func (sc StatefulDemo) Pack(context node.Context) node.Node {
	return node.PackStateful(sc, context)
}

func (sc StatefulDemo) GetKey() string {
	if sc.Key != "" {
		return sc.Key
	}
	if sc.Value != "" {
		return sc.Value
	}
	return strconv.Itoa(sc.Size)
}

func (sc StatefulDemo) GetConstructor() node.ComponentConstructor {
	size := sc.Size
	if size == 0 {
		size = 15
	}
	return func(this *node.Context) node.Widget {
		return node.Block{
			Children: []node.Widget{
				node.Text{
					Content: "Text ComponentFunc " + sc.Value,
					TextStyle: node.TextStyle{
						Color:      color.RoyalBlue,
						FontSize:   size,
						FontWeight: node.FontWeight900,
					},
				},
				node.BR{},
				node.Text{
					Content: fmt.Sprintf("size: %d", size),
				},
				node.Button{
					Child: node.Text{
						Content: "add",
					},
					Params: node.Params{
						OnClick: func(e node.Event) {
							this.SetState(func() {
								size += 1
								logs.Printf("Push Button, size:%v\n", size)
							})
						},
					},
				},
				sc.Child,
			},
		}
	}
}

type Item struct {
	Title   string
	Content string
	Image   string
}

func (i Item) GetWidget(context *node.Context) node.Widget {
	return node.Row{
		Alignment: node.Center,
		Children: []node.Widget{
			node.Image{
				Src: i.Image,
				Params: node.Params{
					Style: node.Style{
						Height: node.Size{
							Mode:  node.SizeModePx,
							Value: 90,
						},
						Width: node.Size{
							Mode:  node.SizeModePx,
							Value: 90,
						},
					},
				},
			},
			node.Margin{
				Width: 15,
			},
			node.Column{
				Children: []node.Widget{
					node.Inline{
						Children: []node.Widget{
							node.Text{
								Content: i.Title,
								TextStyle: node.TextStyle{
									FontSize:   26,
									FontWeight: node.FontWeight700,
								},
							},
						},
					},
					node.Inline{
						Children: []node.Widget{
							node.Text{
								Content: i.Content,
							},
						},
					},
				},
			},
		},
	}
}

func (i Item) Pack(context node.Context) node.Node {
	return node.PackStateless(i, context)
}

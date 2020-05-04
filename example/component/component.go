package component

import (
	"fmt"
	"github.com/TobiasYin/go_web_ui/node"
	"github.com/TobiasYin/go_web_ui/node/color"
	"github.com/TobiasYin/go_web_ui/vdom"
)

type StatelessDemo struct {
	Value string
}

func (sc StatelessDemo) GetNode(context *node.Context) node.Widget {
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
	Value string
	Child node.Widget
	Size  int
}

func (sc StatefulDemo) Pack(context node.Context) node.Node {
	return node.PackStateful(sc, context)
}

func (sc StatefulDemo) GetKey() string {
	return sc.Value
}

func (sc StatefulDemo) GetConstructor() node.ComponentConstructor {
	size := sc.Size
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
						OnClick: func(e vdom.Event) {
							this.SetState(func() {
								size += 1
								fmt.Printf("Push Button, size:%v\n", size)
							})
						},
					},
				},
				sc.Child,
			},
		}
	}
}

package component

import (
	"fmt"
	"github.com/TobiasYin/goi"
	"github.com/TobiasYin/goi/color"
	"strconv"
)

type StatelessDemo struct {
	Value string
}

func (sc StatelessDemo) GetWidget(context *goi.Context) goi.Widget {
	return goi.Block{
		Children: []goi.Widget{
			goi.Text{
				Content: sc.Value + " Stateless",
			},
			goi.BR{},
		},
	}
}

func (sc StatelessDemo) Pack(context goi.Context) goi.Node {
	return goi.PackStateless(sc, context)
}

type StatefulDemo struct {
	Key   string
	Value string
	Child goi.Widget
	Size  int
}

func (sc StatefulDemo) Pack(context goi.Context) goi.Node {
	return goi.PackStateful(sc, context)
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

func (sc StatefulDemo) GetConstructor() goi.ComponentConstructor {
	size := sc.Size
	return func(this *goi.Context) goi.Widget {
		return goi.Block{
			Children: []goi.Widget{
				goi.Text{
					Content: "Text ComponentFunc " + sc.Value,
					TextStyle: goi.TextStyle{
						Color:      color.RoyalBlue,
						FontSize:   size,
						FontWeight: goi.FontWeight900,
					},
				},
				goi.BR{},
				goi.Text{
					Content: fmt.Sprintf("size: %d", size),
				},
				goi.Button{
					Child: goi.Text{
						Content: "add",
					},
					Params: goi.Params{
						OnClick: func(e goi.Event) {
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

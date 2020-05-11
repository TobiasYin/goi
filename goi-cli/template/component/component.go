package component

import (
	"fmt"
	"github.com/TobiasYin/goi"
	"github.com/TobiasYin/goi/color"
	"github.com/TobiasYin/goi/logs"
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
	if size == 0 {
		size = 15
	}
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

func (i Item) GetWidget(context *goi.Context) goi.Widget {
	return goi.Row{
		Alignment: goi.Center,
		Children: []goi.Widget{
			goi.Image{
				Src: i.Image,
				Params: goi.Params{
					Style: goi.Style{
						Height: goi.Size{
							Mode:  goi.SizeModePx,
							Value: 90,
						},
						Width: goi.Size{
							Mode:  goi.SizeModePx,
							Value: 90,
						},
					},
				},
			},
			goi.Margin{
				Width: 15,
			},
			goi.Column{
				Children: []goi.Widget{
					goi.Inline{
						Children: []goi.Widget{
							goi.Text{
								Content: i.Title,
								TextStyle: goi.TextStyle{
									FontSize:   26,
									FontWeight: goi.FontWeight700,
								},
							},
						},
					},
					goi.Inline{
						Children: []goi.Widget{
							goi.Text{
								Content: i.Content,
							},
						},
					},
				},
			},
		},
	}
}

func (i Item) Pack(context goi.Context) goi.Node {
	return goi.PackStateless(i, context)
}

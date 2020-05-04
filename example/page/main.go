package page

import (
	"fmt"
	"github.com/TobiasYin/go_web_ui/example/component"
	"github.com/TobiasYin/go_web_ui/node"
	"github.com/TobiasYin/go_web_ui/node/color"
	"github.com/TobiasYin/go_web_ui/vdom"
)

func GetMainPage() *node.Page {
	size := 22
	return node.NewPage(func(this *node.Context) node.Widget {
		return node.Column{
			Alignment: node.Center,
			Children: []node.Widget{
				node.Row{
					Expand:    true,
					Alignment: node.Center,
					Children: []node.Widget{
						node.Expanded{
							Flex: 1,
							Child: node.Inline{
								Params: node.Params{
									Style: node.Style{
										Height: node.Size{
											Mode:  node.SizeModePx,
											Value: 50,
										},
									},
								},
								Children: []node.Widget{
									node.Text{
										Content: "Hello",
									},
								},
							},
						},
						node.Margin{
							Left: 10,
						},
						node.Inline{
							Children: []node.Widget{
								node.Text{
									Content: "World",
								},
							},
						},
					},
				},
				node.Inline{
					Children: []node.Widget{
						node.Text{
							Content: "aaa",
							TextStyle: node.TextStyle{
								FontSize: size,
								Color:    color.Red,
							},
						},
						node.Text{
							Content: "bbb",
						},
					},
				},
				node.Button{
					Child: node.Text{
						Content: "Click to add 1",
					},
					Params: node.Params{
						OnClick: func(e vdom.Event) {
							fmt.Println("Hello Callback")
							this.SetState(func() {
								size++
							})
						},
					},
				},
				node.Button{
					Child: node.Text{
						Content: "Reset",
					},
					Params: node.Params{
						OnClick: func(e vdom.Event) {
							fmt.Println("Hello Callback")
							this.SetState(func() {
								size = 22
							})
						},
					},
				},
				component.StatefulDemo{
					Key:   node.GenerateKeyWithCallLine(),
					Value: "Hello ! Stateful First ",
					Size:  10,
					Child: node.Text{Content: "I'm child"},
				},
				node.Border{
					Child: node.Button{
						Child: node.Text{
							Content: "To new Page with out keep",
						},
						Params: node.Params{
							OnClick: func(e vdom.Event) {
								_ = node.PushByPathWithPathParams("/image?title=pushbypage1&t2=1")
							},
						},
					},
					Width: 2,
					Color: color.Red,
					Type:  node.BorderTypeSolid,
				},
				node.Border{
					Child: node.Button{
						Child: node.Text{
							Content: "To new Page",
						},
						Params: node.Params{
							OnClick: func(e vdom.Event) {
								_ = node.PushByPathKeepState("/image?title=pushbypage1&t2=1")
							},
						},
					},
					Width: 2,
					Color: color.Red,
					Type:  node.BorderTypeSolid,
				},
				node.Margin{
					Width: 10,
					Child: component.StatefulDemo{
						Key:   node.GenerateKeyWithCallLine(),
						Value: "Hello ! Stateful ",
						Size:  23,
						Child: node.Text{Content: "I'm child"},
					},
				},
				component.StatelessDemo{
					Value: "Tobias",
				},
				component.StatefulDemo{
					Key:   node.GenerateKeyWithCallLine(),
					Value: "Hello ! Stateful AGAIN ",
					Size:  33,
					Child: node.Text{Content: "I'm child"},
				},
			},
		}
	})
}

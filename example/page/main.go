package page

import (
	"fmt"
	"github.com/TobiasYin/goi"
	"github.com/TobiasYin/goi/color"
	"github.com/TobiasYin/goi/example/component"
)

func GetMainPage() *goi.Page {
	size := 22
	return goi.NewPage("Main", func(this *goi.Context) goi.Widget {
		return goi.Column{
			Alignment: goi.Center,
			Children: []goi.Widget{
				goi.Row{
					Expand:    true,
					Alignment: goi.Center,
					Children: []goi.Widget{
						goi.Expanded{
							Flex: 1,
							Child: goi.Inline{
								Params: goi.Params{
									Style: goi.Style{
										Height: goi.Size{
											Mode:  goi.SizeModePx,
											Value: 50,
										},
									},
								},
								Children: []goi.Widget{
									goi.Text{
										Content: "Hello",
									},
								},
							},
						},
						goi.Margin{
							Left: 10,
						},
						goi.Inline{
							Children: []goi.Widget{
								goi.Text{
									Content: "World",
								},
							},
						},
					},
				},
				goi.Inline{
					Children: []goi.Widget{
						goi.Text{
							Content: "aaa",
							TextStyle: goi.TextStyle{
								FontSize: size,
								Color:    color.Red,
							},
						},
						goi.Text{
							Content: "bbb",
						},
					},
				},
				goi.Button{
					Child: goi.Text{
						Content: "Click to add 1",
					},
					Params: goi.Params{
						OnClick: func(e goi.Event) {
							fmt.Println("Hello Callback")
							this.SetState(func() {
								size++
							})
						},
					},
				},
				goi.Button{
					Child: goi.Text{
						Content: "Reset",
					},
					Params: goi.Params{
						OnClick: func(e goi.Event) {
							fmt.Println("Hello Callback")
							this.SetState(func() {
								size = 22
							})
						},
					},
				},
				component.StatefulDemo{
					Key:   goi.GenerateKeyWithCallLine(),
					Value: "Hello ! Stateful First ",
					Size:  10,
					Child: goi.Text{Content: "I'm child"},
				},
				goi.Border{
					Child: goi.Button{
						Child: goi.Text{
							Content: "To new Page with out keep",
						},
						Params: goi.Params{
							OnClick: func(e goi.Event) {
								_ = goi.PushByPathWithPathParams("/image?title=pushbypage1&t2=1")
							},
						},
					},
					Width: 2,
					Color: color.Red,
					Type:  goi.BorderTypeSolid,
				},
				goi.Border{
					Child: goi.Button{
						Child: goi.Text{
							Content: "To new Page",
						},
						Params: goi.Params{
							OnClick: func(e goi.Event) {
								_ = goi.PushByPathKeepState("/image?title=pushbypage1&t2=1")
							},
						},
					},
					Width: 2,
					Color: color.Red,
					Type:  goi.BorderTypeSolid,
				},
				goi.Margin{
					Width: 10,
					Child: component.StatefulDemo{
						Key:   goi.GenerateKeyWithCallLine(),
						Value: "Hello ! Stateful ",
						Size:  23,
						Child: goi.Text{Content: "I'm child"},
					},
				},
				component.StatelessDemo{
					Value: "Tobias",
				},
				component.StatefulDemo{
					Key:   goi.GenerateKeyWithCallLine(),
					Value: "Hello ! Stateful AGAIN ",
					Size:  33,
					Child: goi.Text{Content: "I'm child"},
				},
			},
		}
	})
}

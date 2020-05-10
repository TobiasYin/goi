package page

import (
	"github.com/TobiasYin/goi/node"
	"github.com/TobiasYin/goi/tool/template/component"
)

func GetMainPage() *node.Page {
	return node.NewPage("Main", func(this *node.Context) node.Widget {
		return node.Column{
			Alignment: node.Center,
			Children: []node.Widget{
				node.Link{
					Child: node.Text{Content: "Image page"},
					Href: "/image?title=hello!",
				},
				node.Button{
					Child: node.Text{Content: "To new Page"},
					Params: node.Params{
						OnClick: func(e node.Event) {
							node.PushByPath("/image", map[string]interface{}{
								"title": "base64 test",
							})
						},
					},
				},
				node.Text{Content: "Hello"},
				component.StatefulDemo{
					Key: node.GenerateKeyWithCallLine(),
					Value: " Hello ",
					Size: 22,
				},
				component.StatelessDemo{
					Value: "World",
				},
				node.Link{
					Child: node.Text{Content: "Go To List Page"},
					Href: "/list",
				},
			},
		}
	})
}

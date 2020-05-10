package page

import (
	"github.com/TobiasYin/goi"
	"github.com/TobiasYin/goi/tool/template/component"
)

func GetMainPage() *goi.Page {
	return goi.NewPage("Main", func(this *goi.Context) goi.Widget {
		return goi.Column{
			Alignment: goi.Center,
			Children: []goi.Widget{
				goi.Link{
					Child: goi.Text{Content: "Image page"},
					Href:  "/image?title=hello!",
				},
				goi.Button{
					Child: goi.Text{Content: "To new Page"},
					Params: goi.Params{
						OnClick: func(e goi.Event) {
							goi.PushByPath("/image", map[string]interface{}{
								"title": "base64 test",
							})
						},
					},
				},
				goi.Text{Content: "Hello"},
				component.StatefulDemo{
					Key:   goi.GenerateKeyWithCallLine(),
					Value: " Hello ",
					Size:  22,
				},
				component.StatelessDemo{
					Value: "World",
				},
				goi.Link{
					Child: goi.Text{Content: "Go To List Page"},
					Href:  "/list",
				},
			},
		}
	})
}

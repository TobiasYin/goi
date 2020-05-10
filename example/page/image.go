package page

import (
	"github.com/TobiasYin/goi"
	"github.com/TobiasYin/goi/example/component"
)

type imagePage struct {
	title string
}

func (i imagePage)GetPage() *goi.Page {
	return goi.NewPage("Image", func(this *goi.Context) goi.Widget {
		return goi.Column{
			Alignment: goi.Center,
			Children: []goi.Widget{
				goi.Text{
					Content: "Page Title: " + i.title,
				},
				goi.Text{
					Content: "New Page",
				},
				goi.Link{
					Child: goi.Text{Content: "new"},
					Href:  "new",
				},
				goi.Link{
					Child: goi.Text{Content: "baidu"},
					Href:  "http://www.baidu.com",
				},
				component.StatefulDemo{
					Key:   goi.GenerateKeyWithCallLine(),
					Value: "In Page",
				},
				goi.BR{},
				goi.Image{
					Src: "/example.png",
					Params: goi.Params{
						Style: goi.Style{
							Height: goi.Size{
								Value: 100,
							},
							Width: goi.Size{
								Value: 100,
							},
						},
					},
				},
				goi.BR{},
				goi.Button{
					Child: goi.Text{
						Content: "back",
					},
					Params: goi.Params{
						OnClick: func(e goi.Event) {
							goi.BackToLastPage()
						},
					},
				},
			},
		}
	})
}

func IntoImage(m map[string]interface{}) goi.PageGetter {
	n, ok := m["title"]
	title := ""
	if ok {
		title, _ = n.(string)
	}
	return imagePage{title: title}
}
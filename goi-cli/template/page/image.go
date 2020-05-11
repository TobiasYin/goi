package page

import (
	"github.com/TobiasYin/goi"
	"github.com/TobiasYin/goi/goi-cli/template/component"
)

type imagePage struct {
	title string
}

func (i imagePage) GetPage() *goi.Page {
	imageWidth := 100
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
					Child: goi.Text{Content: "baidu"},
					Href:  "http://www.baidu.com",
				},
				component.StatefulDemo{
					Key:   goi.GenerateKeyWithCallLine(),
					Value: "In Page",
				},
				goi.BR{},
				goi.Image{
					Src: "asset/image/example.png",
					Params: goi.Params{
						Style: goi.Style{
							Height: goi.Size{
								Value: imageWidth,
							},
							Width: goi.Size{
								Value: imageWidth,
							},
						},
					},
				},
				goi.Row{
					Children: []goi.Widget{
						goi.Button{
							Child: goi.Text{
								Content: "Enlarge Image",
							},
							Params: goi.Params{
								OnClick: func(e goi.Event) {
									this.SetState(func() {
										imageWidth += 5
									})
								},
							},
						},
						goi.Button{
							Child: goi.Text{
								Content: "Smaller Image",
							},
							Params: goi.Params{
								OnClick: func(e goi.Event) {
									this.SetState(func() {
										imageWidth -= 5
									})
								},
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

func NewImagePage(m map[string]interface{}) goi.PageGetter {
	n, ok := m["title"]
	title := ""
	if ok {
		title, _ = n.(string)
	}
	return imagePage{title: title}
}
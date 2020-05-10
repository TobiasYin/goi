package page

import (
	"github.com/TobiasYin/goi/node"
	"github.com/TobiasYin/goi/tool/template/component"
)

type imagePage struct {
	title string
}

func (i imagePage) GetPage() *node.Page {
	imageWidth := 100
	return node.NewPage("Image", func(this *node.Context) node.Widget {
		return node.Column{
			Alignment: node.Center,
			Children: []node.Widget{
				node.Text{
					Content: "Page Title: " + i.title,
				},
				node.Text{
					Content: "New Page",
				},
				node.Link{
					Child: node.Text{Content: "baidu"},
					Href:  "http://www.baidu.com",
				},
				component.StatefulDemo{
					Key:   node.GenerateKeyWithCallLine(),
					Value: "In Page",
				},
				node.BR{},
				node.Image{
					Src: "asset/image/example.png",
					Params: node.Params{
						Style: node.Style{
							Height: node.Size{
								Value: imageWidth,
							},
							Width: node.Size{
								Value: imageWidth,
							},
						},
					},
				},
				node.Row{
					Children: []node.Widget{
						node.Button{
							Child: node.Text{
								Content: "Enlarge Image",
							},
							Params: node.Params{
								OnClick: func(e node.Event) {
									this.SetState(func() {
										imageWidth += 5
									})
								},
							},
						},
						node.Button{
							Child: node.Text{
								Content: "Smaller Image",
							},
							Params: node.Params{
								OnClick: func(e node.Event) {
									this.SetState(func() {
										imageWidth -= 5
									})
								},
							},
						},
					},
				},
				node.BR{},
				node.Button{
					Child: node.Text{
						Content: "back",
					},
					Params: node.Params{
						OnClick: func(e node.Event) {
							node.BackToLastPage()
						},
					},
				},
			},
		}
	})
}

func NewImagePage(m map[string]interface{}) node.PageGetter {
	n, ok := m["title"]
	title := ""
	if ok {
		title, _ = n.(string)
	}
	return imagePage{title: title}
}
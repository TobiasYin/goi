package page

import (
	"fmt"
	"github.com/TobiasYin/go_web_ui/example/component"
	"github.com/TobiasYin/go_web_ui/node"
	"github.com/TobiasYin/go_web_ui/vdom"
)

type imagePage struct {
	title string
}

func (i imagePage)GetPage() *node.Page {
	return node.NewPage(func(this *node.Context) node.Widget {
		return node.Column{
			Alignment: node.Right,
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
					Key: node.GenerateKeyWithCallLine(),
					Value: "In Page",
				},
				node.BR{},
				node.Image{
					Src: "/example.png",
					Params: node.Params{
						Style: node.Style{
							Height: node.Size{
								Value: 100,
							},
							Width: node.Size{
								Value: 100,
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
						OnClick: func(e vdom.Event) {
							node.BackToLastPage()
						},
					},
				},
			},
		}
	})
}

func IntoImage(m map[string]interface{}) node.PageGetter {
	n, ok := m["title"]
	x, _ := m["t2"]
	fmt.Println("t2", x)
	title := ""
	if ok {
		title, _ = n.(string)
	}
	return imagePage{title: title}
}
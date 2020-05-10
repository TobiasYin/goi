package page

import (
	"github.com/TobiasYin/goi/node"
)

type relativePage struct {
}

func (r relativePage) GetPage() *node.Page {
	return node.NewPage("Relative", func(this *node.Context) node.Widget {
		return node.Column{
			Alignment: node.Center,
			Children: []node.Widget{
				node.Block{
					Children: []node.Widget{
						node.Text{Content: "welcome to new page"},
					},
				},
			},
		}
	})
}

func NewRelative(m map[string]interface{}) node.PageGetter {
	return relativePage{}
}

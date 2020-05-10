package page

import (
	"github.com/TobiasYin/goi"
)

type relativePage struct {
}

func (r relativePage) GetPage() *goi.Page {
	return goi.NewPage("Relative", func(this *goi.Context) goi.Widget {
		return goi.Column{
			Alignment: goi.Center,
			Children: []goi.Widget{
				goi.Block{
					Children: []goi.Widget{
						goi.Text{Content: "welcome to new page"},
					},
				},
			},
		}
	})
}

func NewRelative(m map[string]interface{}) goi.PageGetter {
	return relativePage{}
}

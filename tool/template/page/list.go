package page

import (
	"fmt"
	"github.com/TobiasYin/goi"
	"github.com/TobiasYin/goi/tool/template/component"
	"math/rand"
	"strings"
)

type listPage struct {
}

func (i listPage) GetPage() *goi.Page {
	var itemList []goi.Widget
	itemList = append(itemList, component.StatefulDemo{
		Value: "List",
	})
	for i:=0; i< 100; i++{
		var content strings.Builder
		for j := 0; j < rand.Intn(10) + 10; j++{
			content.WriteString(fmt.Sprintf("Content %d! ", i))
		}
		itemList = append(itemList, goi.Margin{
			Child: component.Item{
				Title: fmt.Sprintf("Title %d", i),
				Content: content.String(),
				Image: "asset/image/example.png",
			},
			Width: 20,
		})
	}
	return goi.NewPage("List", func(this *goi.Context) goi.Widget {
		return goi.Column{
			Children: itemList,
		}
	})
}

func NewListPage(m map[string]interface{}) goi.PageGetter {
	return listPage{}
}

package page

import (
	"fmt"
	"github.com/TobiasYin/go_web_ui/node"
	"github.com/TobiasYin/go_web_ui/tool/template/component"
	"math/rand"
	"strings"
)

type listPage struct {
}

func (i listPage) GetPage() *node.Page {
	var itemList []node.Widget
	itemList = append(itemList, component.StatefulDemo{
		Value: "List",
	})
	for i:=0; i< 100; i++{
		var content strings.Builder
		for j := 0; j < rand.Intn(10) + 10; j++{
			content.WriteString(fmt.Sprintf("Content %d! ", i))
		}
		itemList = append(itemList, node.Margin{
			Child: component.Item{
				Title: fmt.Sprintf("Title %d", i),
				Content: content.String(),
				Image: "assert/image/example.png",
			},
			Width: 20,
		})
	}
	return node.NewPage("List", func(this *node.Context) node.Widget {
		return node.Column{
			Children: itemList,
		}
	})
}

func NewListPage(m map[string]interface{}) node.PageGetter {
	return listPage{}
}

package urls

import (
	"github.com/TobiasYin/goi/node"
	pages "github.com/TobiasYin/goi/tool/template/page"
)

func init() {
	node.RegisterRoute("/image", pages.NewImagePage)
	node.RegisterRoute("/list", pages.NewListPage)
}

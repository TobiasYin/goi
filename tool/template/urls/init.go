package urls

import (
	"github.com/TobiasYin/go_web_ui/node"
	pages "github.com/TobiasYin/go_web_ui/tool/template/page"
)

func init() {
	node.RegisterRoute("/image", pages.NewImagePage)
	node.RegisterRoute("/list", pages.NewListPage)
}

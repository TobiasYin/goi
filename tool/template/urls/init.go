package urls

import (
	"github.com/TobiasYin/goi"
	pages "github.com/TobiasYin/goi/tool/template/page"
)

func init() {
	goi.RegisterRoute("/image", pages.NewImagePage)
	goi.RegisterRoute("/list", pages.NewListPage)
}

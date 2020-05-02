package dom

import (
	"syscall/js"
)

var (
	Dom Document
	Win Window
)

func init() {
	Dom = Document{NewJsDomElement(js.Global().Get("document"))}
	Win = Window{
		Value{
			js.Global(),
		},
	}
}

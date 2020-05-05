package node

import (
	"fmt"
	"github.com/TobiasYin/go_web_ui/vdom"
	"time"
)

func init() {
	allowRerender <- 1
	go renderLoop()
}

var (
	allowRerender = make(chan int, 1)
	needRerender  = true
)

type main struct {
	page *Page
}

func (m main) GetPage() *Page {
	return m.page
}
func NewApp(page *Page) {
	RegisterRoute("/", func(m map[string]interface{}) PageGetter {
		return main{page}
	})
	page.path = "/"
	PushToPage(page)
	initPush()
	//在这里等待，防止wasm退出。
	c := make(chan struct{})
	<-c
}

func renderLoop() {
	for {
		time.Sleep(time.Millisecond * 50)
		if !needRerender {
			continue
		}
		select {
		case <-allowRerender:
			go func() {
				needRerender = false
				rerender()
				allowRerender <- 1
			}()
		default:
		}
	}
}

func FlashApp() {
	needRerender = true
}

func rerender() {
	start := time.Now()
	top := stack.Top()
	d := top.pack()
	vdom.MergeTwoTree(&d, top.oldDom)
	top.oldDom = &d
	end := time.Now()
	fmt.Printf("Re Render Page, Using: %v\n", end.Sub(start))
}

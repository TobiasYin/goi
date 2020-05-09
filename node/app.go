package node

import (
	"github.com/TobiasYin/go_web_ui/logs"
	"github.com/TobiasYin/go_web_ui/vdom"
	"time"
)

func init() {
	allowRerender <- 1
	updateList = make(map[*Context]struct{})
	SetMaxFPS(30)
	go renderLoop()
}

var (
	allowRerender  = make(chan int, 1)
	needRerender   = true
	mainPage       main
	frameTime      time.Duration
	updateList     map[*Context]struct{}
	needRenderPage = true
	nilStruct      = struct{}{}
)

type main struct {
	page *Page
}

func panicCatch() {
	if r := recover(); r != nil {
		logs.Errorf("panic found!, catch: %v\n", r)
	}
}

func recoverReverse() {
	if r := recover(); r != nil {
		logs.Errorf("panic found! try to recover! %v\n", r)
		page := mainPage.page
		if stack.size == 0 {
			go newApp(page)
			return
		}
		newStack := pageStack{}
		for stack.size > 0 {
			newStack.Add(stack.Pop())
		}
		go newApp(newStack.Pop())
		for newStack.size > 0 {
			top := newStack.Pop()
			if page.path != "" {
				_ = PushByPathWithPathParams(top.path)
			} else {
				PushToPage(top)
			}
		}
	}
}

func (m main) GetPage() *Page {
	return m.page
}

func newApp(page *Page) {
	defer recoverReverse()
	mainPage = main{page}
	RegisterRoute("/", func(m map[string]interface{}) PageGetter {
		return mainPage
	})
	page.path = "/"
	PushToPage(page)
	initPush()
}

func NewApp(page *Page) {
	defer recoverReverse()
	newApp(page)
	//在这里等待，防止wasm退出。
	c := make(chan struct{})
	<-c
}

// 设置每秒最大帧率
func SetMaxFPS(fps int) {
	if fps > 144 {
		fps = 144
	}
	if fps < 10 {
		fps = 10
	}
	frameTime = time.Second / time.Duration(fps)
}

func renderLoop() {
	defer recoverReverse()
	for {
		time.Sleep(frameTime)
		if !needRerender {
			continue
		}
		select {
		case <-allowRerender:
			go func() {
				defer recoverReverse()
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
	needRenderPage = true
}

func addRerenderContext(ctx *Context) {
	updateList[ctx] = nilStruct
	needRerender = true
}

func rerender() {
	clock := logs.Clock{Hint: "Re Render Batch"}
	clock.Start()
	if needRenderPage {
		rerenderPage()
		needRenderPage = false
	}
	if len(updateList) > 0 {
		// TODO 遍历Map是很慢的操作，考虑更换成List。
		for k, _ := range updateList {
			k.doSetState()
		}
		updateList = make(map[*Context]struct{})
	}
	clock.End()
}

func rerenderPage() {
	clock := logs.Clock{Hint: "Re Render Page"}
	clock.Start()
	top := stack.Top()
	d := top.pack()
	vdom.MergeTwoTree(&d, top.oldTree)
	top.oldTree = &d
	clock.End()
}

func rerenderTree(newTree *vdom.JsDomElement, oldTree *vdom.JsDomElement) {
	clock := logs.Clock{Hint: "Re Render Tree"}
	clock.Start()
	vdom.MergeTwoContext(newTree, oldTree)
	clock.End()
}

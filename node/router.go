package node

import (
	"fmt"
	rdom "github.com/TobiasYin/go_web_ui/dom"
	dom "github.com/TobiasYin/go_web_ui/vdom"
	"strings"
)

type pageStack struct {
	stack []*Page
	size  int
}

var (
	router        map[string]NewPathPage
	stack         pageStack
	keepStatePage map[string]*Page
	pushByCode    map[string]bool
)

func pushWithHash(path string)  {
	splitPath := strings.Split(path, "?")
	arg := make(map[string]interface{})
	if len(splitPath) > 1 {
		args := strings.Split(splitPath[1], "&")
		path = splitPath[0]
		for _, v := range args {
			if v == "" {
				continue
			}
			value := ""
			a := strings.Split(v, "=")
			key := a[0]
			if len(a) > 1 {
				value = a[1]
			}
			arg[key] = value
		}
	}
	page, ok := router[path]
	if !ok {
		rdom.Win.SetHash("/")
		return
	}
	p := page(arg).GetPage()
	p.path = path
	PushToPage(p)
}

func listenHash(event rdom.Event) {
	path := ""
	hash := rdom.Win.GetHash()
	if hash == "" {
		path = "/"
	} else {
		path = hash[1:]
	}
	if v := pushByCode[path]; v {
		pushByCode[path] = false
		return
	}
	pushWithHash(path)
}

func initPush()  {
	path := ""
	hash := rdom.Win.GetHash()
	if hash == "" {
		path = "/"
	} else {
		path = hash[1:]
	}
	pushWithHash(path)
}

func init() {
	router = make(map[string]NewPathPage)
	keepStatePage = make(map[string]*Page)
	pushByCode = make(map[string]bool)
	rdom.Win.AddEventListener("hashchange", listenHash)
}

func RegisterRoute(path string, page NewPathPage) {
	if len(path) == 0 {
		path = "/"
	} else if path[0] != '/' {
		path = "/" + path
	}
	path = strings.Split(path, "?")[0]
	router[path] = page
}

func (p *pageStack) Top() *Page {
	return p.stack[p.size-1]
}

func (p *pageStack) Pop() *Page {
	p.size--
	top := p.stack[p.size]
	p.stack[p.size] = nil
	p.stack = p.stack[0:p.size]
	return top
}

func (p *pageStack) Add(page *Page) {
	p.stack = append(p.stack, page)
	p.size++
}

func (p *pageStack) pack() dom.JsDomElement {
	return p.Top().pack()
}

func PushToPage(page *Page) {
	stack.Add(page)
	if page.oldDom != nil {
		dom.Display(page.oldDom)
	}
	FlashApp()
}

func BackToLastPage() {
	stack.Pop()
	top := stack.Top()
	setHash(top.path)
	dom.Display(top.oldDom)
	FlashApp()
}

func PushByPathWithPathParams(path string) error {
	return PushByPath(path, map[string]interface{}{})
}

//这种页面全局唯一，不接受传参，因为参数仅在第一次有效, 可重新封装New方法
func PushByPathKeepState(path string) error {
	if len(path) == 0 {
		path = "/"
	} else if path[0] != '/' {
		path = "/" + path
	}
	path = strings.Split(path, "?")[0]
	pageGetter, ok := router[path]
	if !ok {
		return fmt.Errorf("unkonw page")
	}
	var page *Page = nil
	v, ok := keepStatePage[path]
	if ok {
		page = v
	} else {
		page = pageGetter(map[string]interface{}{}).GetPage()
		keepStatePage[path] = page
	}
	setHash(path)
	page.path = path
	PushToPage(page)
	return nil
}

func PushByPath(path string, arg map[string]interface{}) error {
	splitPath := strings.Split(path, "?")
	if len(splitPath) > 1 {
		args := strings.Split(splitPath[1], "&")
		path = splitPath[0]
		for _, v := range args {
			if v == "" {
				continue
			}
			value := ""
			a := strings.Split(v, "=")
			key := a[0]
			if len(a) > 1 {
				value = a[1]
			}
			arg[key] = value
		}
	}
	page, ok := router[path]
	if !ok {
		return fmt.Errorf("unkonw page")
	}
	var params strings.Builder
	for k, v := range arg {
		params.WriteString(fmt.Sprintf("%s=%v&", k, v))
	}
	paramsStr := params.String()
	if paramsStr[0] > 0 {
		paramsStr = paramsStr[:len(paramsStr)-1]
	}
	hash := fmt.Sprintf("%s?%s", path, paramsStr)
	setHash(hash)
	p := page(arg).GetPage()
	p.path = hash
	PushToPage(p)
	return nil
}

type PageGetter interface {
	GetPage() *Page
}

func setHash(hash string) {
	pushByCode[hash] = true
	rdom.Win.SetHash(hash)
}

func GetNowPath() string {
	path:=stack.Top().path
	path = strings.Split(path, "?")[0]
	if path[len(path) - 1] != '/' {
		path = path + "/"
	}
	return path
}

type NewPathPage func(map[string]interface{}) PageGetter

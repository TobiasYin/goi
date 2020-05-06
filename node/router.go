package node

import (
	"encoding/base64"
	"encoding/json"
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
	encoding      *base64.Encoding
)

const (
	param       = "param"
	encodingStr = "uVPKwOaXhW5nkCF3bLem7s1cfZy9ldJRqQvYHot6z+Sp-GD4riBjxI2MUN8EATg0"
	encodingPad = '_'
)


func init() {
	router = make(map[string]NewPathPage)
	keepStatePage = make(map[string]*Page)
	pushByCode = make(map[string]bool)
	rdom.Win.AddEventListener("hashchange", WrapEventCallBack(listenHash))
	encoding = base64.NewEncoding(encodingStr).WithPadding(encodingPad)
}

func encodeParams(v map[string]interface{}) string {
	res, err := json.Marshal(v)
	if err != nil {
		res = []byte("{}")
	}
	return encoding.EncodeToString(res)
}

func decodeParams(p string) map[string]interface{} {
	str, err := encoding.DecodeString(p)
	if err != nil {
		str = []byte("{}")
	}
	res := make(map[string]interface{})
	err = json.Unmarshal(str, &res)
	if err != nil {
		return make(map[string]interface{})
	}
	return res
}

func mergeArg(arg *map[string]interface{}, path *string) {
	splitPath := strings.Split(*path, "?")
	if len(splitPath) > 1 {
		args := strings.Split(splitPath[1], "&")
		*path = splitPath[0]
		for _, v := range args {
			if v == "" {
				continue
			}
			a := strings.Split(v, "=")
			key := a[0]
			if key == param {
				newArg := decodeParams(a[1])
				for k, v2 := range newArg {
					if _, ok := (*arg)[k]; !ok {
						(*arg)[k] = v2
					}
				}
			} else {
				if key == "" {
					continue
				}
				value := ""
				if len(a) > 1 {
					value = a[1]
				}
				if _, ok := (*arg)[key]; !ok {
					(*arg)[key] = value
				}
			}

		}
	}
}

func getUrlByPathAndArg(path string, arg map[string]interface{}) string {
	return fmt.Sprintf("%s?%s=%s", path, param, encodeParams(arg))
}

func pushWithHash(path string) {
	arg := make(map[string]interface{})
	mergeArg(&arg, &path)
	page, ok := router[path]
	if !ok {
		rdom.Win.SetHash("/")
		return
	}
	p := page(arg).GetPage()
	p.path = getUrlByPathAndArg(path, arg)
	PushToPage(p)
}

func listenHash(event Event) {
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

func initPush() {
	path := ""
	hash := rdom.Win.GetHash()
	if hash == "" {
		path = "/"
	} else {
		path = hash[1:]
	}
	if path == "/" {
		return
	}
	pushWithHash(path)
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
	if page.Title != "" {
		rdom.Dom.SetTitle(page.Title)
	}
	stack.Add(page)
	if page.oldDom != nil {
		dom.Display(page.oldDom)
	}
	FlashApp()
}

func BackToLastPage() {
	if stack.size == 0 {
		_ = PushByPath("/", map[string]interface{}{})
		return
	}
	stack.Pop()
	top := stack.Top()
	if top.Title != "" {
		rdom.Dom.SetTitle(top.Title)
	}
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
	mergeArg(&arg, &path)
	page, ok := router[path]
	if !ok {
		return fmt.Errorf("unkonw page")
	}
	hash := getUrlByPathAndArg(path, arg)
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
	path := stack.Top().path
	path = strings.Split(path, "?")[0]
	if path[len(path)-1] != '/' {
		path = path + "/"
	}
	return path
}

type NewPathPage func(map[string]interface{}) PageGetter

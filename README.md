# Go WebAssembly UI Kit

## 简介
本项目旨在用go语言声明式的编写响应式的web app。
项目使用类似flutter的语法，帮助开发人员快速的建构web应用。

## 示例

```go
package main

import (
	"fmt"
	"github.com/TobiasYin/go_web_ui/node"
	"github.com/TobiasYin/go_web_ui/node/color"
)

func main() {
	c := make(chan struct{})
	size := 22
	page2 := node.NewPage(func(this *node.Context) node.Node {
		return node.Div{
			Children: []node.Node{
				node.Text{
					Content: "New Page",
				},
				node.BR{},
				node.Image{Src: "/example.png"},
				node.BR{},
				node.Button{
					Child: node.Text{
						Content: "back",
					},
					Params: node.Params{
						OnClick: func(e node.Event) {
							node.BackToLastPage()
						},
					},
				},
			},
		}
	})
	page := node.NewPage(func(this *node.Context) node.Node {
		return node.Div{
			Children: []node.Node{
				node.Div{
					Children: []node.Node{
						node.Text{
							Content: "aaa",
							TextStyle: &node.TextStyle{
								FontSize:  size,
								FontColor: color.Red,
							},
						},
						node.Text{
							Content: "bbb",
						},
					},
				},
				node.Button{
					Child: node.Text{
						Content: "Click to add 1",
					},
					Params: node.Params{
						OnClick: func(e node.Event) {
							fmt.Println("Hello Callback")
							this.SetState(func() {
								size++
							})
						},
					},
				},
				node.Button{
					Child: node.Text{
						Content: "Reset",
					},
					Params: node.Params{
						OnClick: func(e node.Event) {
							fmt.Println("Hello Callback")
							this.SetState(func() {
								size = 22
							})
						},
					},
				},
				this.StatefulChild(Demo),
				node.Button{
					Child: node.Text{
						Content: "To new Page",
					},
					Params: node.Params{
						OnClick: func(e node.Event) {
							node.PushToPage(page2)
						},
					},
				},
				this.StatefulChild(Demo),
				this.StatefulChild(ComponentWithPara("hello")),
				this.StatefulChild(ComponentWithPara("daisy: ")),
				this.StatelessChild(StatelessDemo),
			},
		}
	})
	node.NewApp(page)
	<-c
}

func ComponentWithPara(aaa string) node.Component {
	return func() node.ComponentConstructor {
		hello := 0
		return func(context *node.Context) node.Node {
			return node.Div{
				Children: []node.Node{
					node.Text{
						Content: aaa,
					},
					node.Text{
						Content: fmt.Sprintf("value: %d", hello),
					},
					node.Button{
						Child: node.Text{
							Content: "increase",
						},
						Params: node.Params{
							OnClick: func(e node.Event) {
								context.SetState(func() {
									hello++
								})
							},
						},
					},
				},
			}
		}
	}
}

func StatelessDemo(context *node.Context) node.Node {
	return node.Div{
		Children: []node.Node{
			node.Text{
				Content: "Stateless\n",
			},
		},
	}
}

func Demo() node.ComponentConstructor {
	size := 22
	return func(this *node.Context) node.Node {
		return node.Div{
			Children: []node.Node{
				node.Text{
					Content: "Text Component",
					TextStyle: &node.TextStyle{
						FontColor:  color.RoyalBlue,
						FontSize:   size,
						FontWeight: node.FontWeight900,
					},
				},
				node.BR{},
				node.Text{
					Content: fmt.Sprintf("size: %d", size),
				},
				node.Button{
					Child: node.Text{
						Content: "add",
					},
					Params: node.Params{
						OnClick: func(e node.Event) {
							this.SetState(func() {
								size += 1
								fmt.Printf("Push Button, size:%v\n", size)
							})
						},
					},
				},
			},
		}
	}
}
```

接下来将分别讲解Demo的多个部分的含义。

```go
func main() {
    c := make(chan struct{})
	size := 22
	page2 := node.NewPage(func(this *node.Context) node.Node {
		return node.Div{
			Children: []node.Node{
				node.Text{
					Content: "New Page",
				},
				node.BR{},
				node.Image{Src: "/example.png"},
				node.BR{},
				node.Button{
					Child: node.Text{
						Content: "back",
					},
					Params: node.Params{
						OnClick: func(e node.Event) {
							node.BackToLastPage()
						},
					},
				},
			},
		}
	})
	page := node.NewPage(func(this *node.Context) node.Node {
		return node.Div{
			Children: []node.Node{
				node.Div{
					Children: []node.Node{
						node.Text{
							Content: "aaa",
							TextStyle: &node.TextStyle{
								FontSize:  size,
								FontColor: color.Red,
							},
						},
						node.Text{
							Content: "bbb",
						},
					},
				},
				node.Button{
					Child: node.Text{
						Content: "Click to add 1",
					},
					Params: node.Params{
						OnClick: func(e node.Event) {
							fmt.Println("Hello Callback")
							this.SetState(func() {
								size++
							})
						},
					},
				},
				node.Button{
					Child: node.Text{
						Content: "Reset",
					},
					Params: node.Params{
						OnClick: func(e node.Event) {
							fmt.Println("Hello Callback")
							this.SetState(func() {
								size = 22
							})
						},
					},
				},
				this.StatefulChild(Demo),
				node.Button{
					Child: node.Text{
						Content: "To new Page",
					},
					Params: node.Params{
						OnClick: func(e node.Event) {
							node.PushToPage(page2)
						},
					},
				},
				this.StatefulChild(Demo),
				this.StatefulChild(ComponentWithPara("hello")),
				this.StatefulChild(ComponentWithPara("daisy: ")),
				this.StatelessChild(StatelessDemo),
			},
		}
	})
	node.NewApp(page)
	<-c
}
```
c用于保证go主函数在网页关闭前不退出。
接下来使用NewPage方法声明了两个页面。
### 组件
在声明页面的代码中，我们可看见StatefulChild，StatelessChild这样的方法，是使用子组件的方法。

组件使用函数式的方法创建，组件的状态通过闭包来存储。

要制作一个无状态的组件，需要实现这样的函数，返回一个node即可。
```go
func StatelessDemo(context *node.Context) node.Node {
	return node.Div{
		Children: []node.Node{
			node.Text{
				Content: "Stateless\n",
			},
		},
	}
}
```

要实现一个有状态无参数的组件，则略微复杂一点，以为状态使用闭包来保存，所以返回的对象其实四个函数。

```go
func Demo() node.ComponentConstructor {
	size := 22
	return func(this *node.Context) node.Node {
		return node.Div{
			Children: []node.Node{
				node.Text{
					Content: "Text Component",
					TextStyle: &node.TextStyle{
						FontColor:  color.RoyalBlue,
						FontSize:   size,
						FontWeight: node.FontWeight900,
					},
				},
				node.Button{
					Child: node.Text{
						Content: "add",
					},
					Params: node.Params{
						OnClick: func(e node.Event) {
							this.SetState(func() {
								size += 1
							})
						},
					},
				},
			},
		}
	}
}
```
这里的size就是通过闭包保存的状态，在刷新组件时，状态将会被保留而不会重置。

除此之外，我们还可以实现有参数有状态的组件，制作这样的组件只需比无状态的组件多一层闭包来保存参数即可。
```go
func ComponentWithPara(aaa string) node.Component {
	return func() node.ComponentConstructor {
		hello := 0
		return func(context *node.Context) node.Node {
			return node.Div{
				Children: []node.Node{
					node.Text{
						Content: aaa,
					},
					node.Text{
						Content: fmt.Sprintf("value: %d", hello),
					},
					node.Button{
						Child: node.Text{
							Content: "increase",
						},
						Params: node.Params{
							OnClick: func(e node.Event) {
								context.SetState(func() {
									hello++
								})
							},
						},
					},
				},
			}
		}
	}
}
```

最后，使用`node.NewApp(page)`就可以把页面注册到根上了！

## 运行
目前暂未提供cli工具，拷贝example文件夹，使用`GOARCH=wasm GOOS=js go build -o main.wasm main.go`命令编译main.go，即可生成main.wasm文件，使用配套的服务端工具，即可运行本框架的程序。

这里提供一个简单的go server，用来运行。

```go
package main

import (
	"flag"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

var (
	listen = flag.String("listen", ":8080", "listen address")
	dir    = flag.String("dir", "./frontend", "directory to serve")
)

func Exists(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}

func IsDir(path string) bool {
	s, err := os.Stat(path)
	if err != nil {
		return false
	}
	return s.IsDir()
}

func main() {
	flag.Parse()
	log.Printf("listening on %q...", *listen)
	server := http.FileServer(http.Dir(*dir))
	log.Fatal(http.ListenAndServe(*listen, http.HandlerFunc(func(resp http.ResponseWriter, req *http.Request) {
		if strings.HasSuffix(req.URL.Path, ".wasm") {
			resp.Header().Set("content-type", "application/wasm")
		}
		path := *dir + req.RequestURI
		if Exists(path) && !IsDir(path) {
			server.ServeHTTP(resp, req)
		} else {
			defer func() { _ = req.Body.Close() }()
			res, err := ioutil.ReadFile(*dir + "/index.html")
			if err != nil {
				resp.WriteHeader(http.StatusInternalServerError)
				_, _ = resp.Write([]byte("internal error"))
				return
			}
			resp.WriteHeader(http.StatusOK)
			_, _ = resp.Write(res)
		}
	})))
}
```

编译词程序，使用--dir参数制定刚刚的example目录，即可运行！
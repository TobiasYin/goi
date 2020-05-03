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
		return node.Column{
			Alignment: node.Right,
			Children: []node.Node{
				node.Text{
					Content: "New Page",
				},
				node.Link{
					Child: node.Text{Content: "baidu"},
					Href:  "http://www.baidu.com",
				},
				node.BR{},
				node.Image{
					Src: "/example.png",
					Params: node.Params{
						Style: node.Style{
							Height: node.Size{
								Value: 100,
							},
							Width: node.Size{
								Value: 100,
							},
						},
					},
				},
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
		return node.Column{
			Alignment: node.Center,
			Children: []node.Node{
				node.Row{
					Expand:    true,
					Alignment: node.Center,
					Children: []node.Node{
						node.Expanded{
							Flex: 1,
							Child: node.Inline{
								Params: node.Params{
									Style: node.Style{
										Height: node.Size{
											Mode:  node.SizeModePx,
											Value: 50,
										},
									},
								},
								Children: []node.Node{
									node.Text{
										Content: "Hello",
									},
								},
							},
						},
						node.Padding{
							Left: 10,
						},
						node.Inline{
							Children: []node.Node{
								node.Text{
									Content: "World",
								},
							},
						},
					},
				},
				node.Inline{
					Children: []node.Node{
						node.Text{
							Content: "aaa",
							TextStyle: node.TextStyle{
								FontSize: size,
								Color:    color.Red,
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
				node.Border{
					Child: node.Button{
						Child: node.Text{
							Content: "To new Page",
						},
						Params: node.Params{
							OnClick: func(e node.Event) {
								node.PushToPage(page2)
							},
						},
					},
					Width: 2,
					Color: color.Red,
					Type:  node.BorderTypeSolid,
				},
				node.Padding{
					Width: 10,
					Child: this.StatefulChild(Demo),
				},
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
			return node.Block{
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
	return node.Block{
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
		return node.Block{
			Children: []node.Node{
				node.Text{
					Content: "Text Component",
					TextStyle: node.TextStyle{
						Color:      color.RoyalBlue,
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

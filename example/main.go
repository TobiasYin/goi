package main

import (
	"fmt"
	"github.com/TobiasYin/go_web_ui/node"
)


func main() {
	c := make(chan struct{})
	size := 22
	page := node.NewPage()
	page2 := node.NewPage()
	page2.GetNode = func() node.Node {
		return node.Div{
			Children: []node.Node{
				node.Text{
					Content: "New Page",
				},
				node.Button{
					Child: node.Text{
						Content: "back",
					},
					OnClick: func(e node.Event) {
						node.BackToLastPage()
					},
				},
			},
		}
	}
	page.GetNode = func() node.Node {
		return node.Div{
			Children: []node.Node{
				node.Div{
					Children: []node.Node{
						node.Text{
							Content: "aaa",
							TextStyle: &node.TextStyle{
								FontSize:  size,
								FontColor: node.Red,
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
					OnClick: func(e node.Event) {
						fmt.Println("Hello Callback")
						page.SetState(func() {
							size++
						})
					},
				},
				node.Button{
					Child: node.Text{
						Content: "Reset",
					},
					OnClick: func(e node.Event) {
						fmt.Println("Hello Callback")
						page.SetState(func() {
							size = 22
						})
					},
				},
				page.StatefulChild(Demo),
				node.Button{
					Child: node.Text{
						Content: "To new Page",
					},
					OnClick: func(e node.Event) {
						node.PushToPage(page2)
					},
				},
				page.StatefulChild(Demo),
			},
		}
	}
	node.NewApp(page)
	<-c
}

func Demo(area node.StateArea) *node.Context {
	size := 22
	newCtx := node.NewContext(area)
	newCtx.GetNode = func() node.Node {
		return node.Div{
			Children: []node.Node{
				node.Text{
					Content: "Text Component",
					TextStyle: &node.TextStyle{
						FontColor:  node.RoyalBlue,
						FontSize:   size,
						FontWeight: node.FontWeight900,
					},
				},
				node.Button{
					Child: node.Text{
						Content: "add",
					},
					OnClick: func(e node.Event) {
						newCtx.SetState(func() {
							size += 1
							fmt.Printf("Push Button, size:%v\n", size)
						})
					},
				},
			},
		}
	}
	return newCtx
}

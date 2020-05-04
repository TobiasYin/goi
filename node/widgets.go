package node

import (
	"fmt"
	"github.com/TobiasYin/go_web_ui/node/color"
	dom "github.com/TobiasYin/go_web_ui/vdom"
	"strings"
)

type Inline struct {
	Params
	Children []Node
}

func (d Inline) getChildren() []Node {
	return d.Children
}

func (d Inline) pack() dom.JsDomElement {
	ele := pack(d, "div")
	ele.Set("className", ele.Get("className").String()+" go-ui-inline-div")
	return ele
}

type Block struct {
	Params
	Children []Node
}

func (d Block) getChildren() []Node {
	return d.Children
}

func (d Block) pack() dom.JsDomElement {
	return pack(d, "div")
}

type P struct {
	Params
	Children []Node
}

func (p P) getChildren() []Node {
	return p.Children
}

func (p P) pack() dom.JsDomElement {
	return pack(p, "p")
}

type Link struct {
	Params
	Child Node
	Href  string
}

func (l Link) getChildren() []Node {
	return []Node{l.Child}
}

func (l Link) pack() dom.JsDomElement {
	ele := pack(l, "a")
	ele.Set("href", l.Href)
	return ele
}

type Input struct {
	Params
}

func (i Input) getChildren() []Node {
	return []Node{}
}

func (i Input) pack() dom.JsDomElement {
	return pack(i, "input")
}

type Button struct {
	Params
	Child Node
}

func (b Button) getChildren() []Node {
	return []Node{b.Child}
}

func (b Button) pack() dom.JsDomElement {
	return pack(b, "button")
}

type BR struct{}

func (br BR) pack() dom.JsDomElement {
	return dom.Dom.CreateElement("br")
}

type Border struct {
	Child       Node
	Width       int
	Left        int
	Right       int
	Bottom      int
	Top         int
	Type        BorderType
	LeftType    BorderType
	RightType   BorderType
	BottomType  BorderType
	TopType     BorderType
	Color       color.Color
	LeftColor   color.Color
	RightColor  color.Color
	TopColor    color.Color
	BottomColor color.Color
}
type border struct {
	Width int
	Type  BorderType
	Color color.Color
}

var BorderNil = Border{
	Color: color.White,
	Width: 0,
}

func (b Border) packStyle() string {
	borders := make([]border, 4)
	for i := 0; i < 4; i++ {
		borders[i].Width = b.Width
		borders[i].Type = b.Type
		borders[i].Color = b.Color
	}
	if b.Left != 0 {
		borders[0].Width = b.Left
	}
	if b.Right != 0 {
		borders[1].Width = b.Right
	}
	if b.Top != 0 {
		borders[2].Width = b.Top
	}
	if b.Bottom != 0 {
		borders[3].Width = b.Bottom
	}

	if b.LeftColor != color.ColorNil {
		borders[0].Color = b.LeftColor
	}
	if b.RightColor != color.ColorNil {
		borders[1].Color = b.RightColor
	}
	if b.TopColor != color.ColorNil {
		borders[2].Color = b.TopColor
	}
	if b.BottomColor != color.ColorNil {
		borders[3].Color = b.BottomColor
	}

	if b.LeftType != "" {
		borders[0].Type = b.LeftType
	}
	if b.RightType != "" {
		borders[1].Type = b.RightType
	}
	if b.TopType != "" {
		borders[2].Type = b.TopType
	}
	if b.BottomType != "" {
		borders[3].Type = b.BottomType
	}
	var style strings.Builder
	style.WriteString("display:inline-block;")
	prefix := []string{"left", "right", "top", "bottom"}
	for i, v := range borders {
		if v.Width == 0 && v.Type == "" && v.Color == color.ColorNil {
			continue
		}
		style.WriteString("border-")
		style.WriteString(prefix[i])
		style.WriteString(":")
		if v.Width != 0 {
			style.WriteString(fmt.Sprintf("%dpx ", v.Width))
		}
		if v.Type != "" {
			style.WriteString(string(v.Type))
			style.WriteString(" ")
		}
		if v.Color != color.ColorNil {
			style.WriteString(v.Color.String())
		}
		style.WriteString(";")
	}
	return style.String()
}

func (b Border) pack() dom.JsDomElement {
	ele := dom.Dom.CreateElement("div")

	style := b.packStyle()
	if len(style) != 0 {
		ele.Set("style", style)
	}
	packChildren(b, &ele)
	return ele
}

func (b Border) getChildren() []Node {
	return []Node{b.Child}
}

type BorderType string

const (
	BorderTypeNone   BorderType = "none"
	BorderTypeHidden BorderType = "hidden"
	BorderTypeDotted BorderType = "dotted"
	BorderTypeDashed BorderType = "dashed"
	BorderTypeSolid  BorderType = "solid"
	BorderTypeDouble BorderType = "double"
	BorderTypeGroove BorderType = "groove"
	BorderTypeRidge  BorderType = "ridge"
	BorderTypeInset  BorderType = "inset"
	BorderTypeOutset BorderType = "outset"
)

type Margin struct {
	Child  Node
	Width  int
	Left   int
	Right  int
	Top    int
	Bottom int
}

func (m Margin) pack() dom.JsDomElement {
	ele := dom.Dom.CreateElement("div")

	margin := make([]int, 4)
	for i := 0; i < 4; i++ {
		margin[i] = m.Width
	}
	if m.Left != 0 {
		margin[0] = m.Left
	}
	if m.Right != 0 {
		margin[1] = m.Right
	}
	if m.Top != 0 {
		margin[2] = m.Top
	}
	if m.Bottom != 0 {
		margin[3] = m.Bottom
	}

	var style strings.Builder
	style.WriteString("display:inline-block;")
	prefix := []string{"left", "right", "top", "bottom"}
	for i, v := range margin {
		if v == 0 {
			continue
		}
		style.WriteString(fmt.Sprintf("padding-%s: %dpx;", prefix[i], v))
	}
	if style.Len() != 0 {
		ele.Set("style", style.String())
	}
	packChildren(m, &ele)
	return ele
}

func (m Margin) getChildren() []Node {
	return []Node{m.Child}
}

type Column struct {
	Children  []Node
	Alignment Position
}

func (c Column) getChildren() []Node {
	alignment := c.Alignment
	if alignment == "" {
		alignment = Left
	}
	res := make([]Node, len(c.Children))
	for i, child := range c.Children {
		res[i] = Block{
			Params: Params{
				Style: Style{
					TextAlign: alignment,
				},
			},
			Children: []Node{
				child,
			},
		}
	}
	return res
}

func (c Column) pack() dom.JsDomElement {
	ele := dom.Dom.CreateElement("div")
	packChildren(c, &ele)
	ele.Set("className", ele.Get("className").String()+" go-ui-column")
	return ele
}

type Row struct {
	Alignment Position
	Expand    bool
	Children  []Node
}

func (r Row) getChildren() []Node {
	return r.Children
}

func (r Row) pack() dom.JsDomElement {
	ele := dom.Dom.CreateElement("div")
	ele.Set("className", ele.Get("className").String()+" go-ui-row")
	style := "align-items:" + string(r.Alignment) + ";"
	if r.Expand {
		style += "display:flex;"
	}
	ele.Set("style", style)
	packChildren(r, &ele)
	return ele
}

type Expanded struct {
	Flex  int
	Child Node
}

func (e Expanded) getChildren() []Node {

	return []Node{e.Child}
}

func (e Expanded) pack() dom.JsDomElement {
	return Inline{
		Params: Params{
			Style: Style{
				FlexGrow: e.Flex,
			},
		},
		Children: []Node{
			e.Child,
		},
	}.pack()
}

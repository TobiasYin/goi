package node

import (
	"fmt"
	"github.com/TobiasYin/go_web_ui/node/color"
	dom "github.com/TobiasYin/go_web_ui/vdom"
	"strings"
)

type _context struct {
	Context
}

func (c _context) getContext() Context {
	return c.Context
}

type Inline struct {
	Params
	Children []Widget
	_context
}

func (d Inline) getChildren() []Widget {
	return d.Children
}

func (d Inline) pack() dom.JsDomElement {
	ele := pack(d, "div", d.getContext())
	ele.Set("className", ele.Get("className").String()+" go-ui-inline-div")
	return ele
}

func (d Inline) Pack(ctx Context) Node {
	d._context.Context = ctx
	return d
}

type Block struct {
	Params
	_context
	Children []Widget
}

func (b Block) getChildren() []Widget {
	return b.Children
}

func (b Block) pack() dom.JsDomElement {
	return pack(b, "div", b.getContext())
}

func (b Block) Pack(ctx Context) Node {
	b._context.Context = ctx
	return b
}

type P struct {
	Params
	_context
	Children []Widget
}

func (p P) getChildren() []Widget {
	return p.Children
}

func (p P) pack() dom.JsDomElement {
	return pack(p, "p", p.Context)
}
func (p P) Pack(ctx Context) Node {
	p._context.Context = ctx
	return p
}

type Link struct {
	Params
	_context
	Child Widget
	Href  string
}

func (l Link) getChildren() []Widget {
	return []Widget{l.Child}
}

func (l Link) pack() dom.JsDomElement {
	ele := pack(l, "a", l.getContext())
	href := l.Href
	if !strings.Contains(href, "://") {
		if href[0] == '/' {
			href = "#" + href
		} else {
			href = "#" + GetNowPath() + href
		}
	}
	ele.Set("href", href)
	return ele
}

func (l Link) Pack(ctx Context) Node {
	l._context.Context = ctx
	return l
}

type Input struct {
	_context
	Params
}

func (i Input) getChildren() []Widget {
	return []Widget{}
}

func (i Input) pack() dom.JsDomElement {
	return pack(i, "input", i.getContext())
}
func (i Input) Pack(ctx Context) Node {
	i._context.Context = ctx
	return i
}

type Button struct {
	Params
	_context
	Child Widget
}

func (b Button) getChildren() []Widget {
	return []Widget{b.Child}
}

func (b Button) pack() dom.JsDomElement {
	return pack(b, "button", b.getContext())
}
func (b Button) Pack(ctx Context) Node {
	b._context.Context = ctx
	return b
}

type BR struct {
	_context
}

func (br BR) pack() dom.JsDomElement {
	return dom.Dom.CreateElement("br")
}
func (br BR) Pack(ctx Context) Node {
	br._context.Context = ctx
	return br
}

type Border struct {
	Child       Widget
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
	_context
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
	if style.Len() != 0{
		style.WriteString("display:inline-block;")
	}
	return style.String()
}

func (b Border) pack() dom.JsDomElement {
	ele := dom.Dom.CreateElement("div")

	style := b.packStyle()
	if len(style) != 0 {
		ele.Set("style", style)
	}
	packChildren(b, &ele, b.getContext())
	return ele
}

func (b Border) getChildren() []Widget {
	return []Widget{b.Child}
}
func (b Border) Pack(ctx Context) Node {
	b._context.Context = ctx
	return b
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
	Child  Widget
	Width  int
	Left   int
	Right  int
	Top    int
	Bottom int
	_context
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
	prefix := []string{"left", "right", "top", "bottom"}
	for i, v := range margin {
		if v == 0 {
			continue
		}
		style.WriteString(fmt.Sprintf("padding-%s: %dpx;", prefix[i], v))
	}
	if style.Len() != 0 {
		style.WriteString("display:inline-block;")
		ele.Set("style", style.String())
	}
	packChildren(m, &ele, m.getContext())
	return ele
}

func (m Margin) getChildren() []Widget {
	return []Widget{m.Child}
}

func (m Margin) Pack(ctx Context) Node {
	m._context.Context = ctx
	return m
}

type Column struct {
	Children  []Widget
	Alignment Position
	_context
}

func (c Column) getChildren() []Widget {
	alignment := c.Alignment
	if alignment == "" {
		alignment = Left
	}
	res := make([]Widget, len(c.Children))
	for i, child := range c.Children {
		res[i] = Block{
			Params: Params{
				Style: Style{
					TextAlign: alignment,
				},
			},
			Children: []Widget{
				child,
			},
		}
	}
	return res
}

func (c Column) pack() dom.JsDomElement {
	ele := dom.Dom.CreateElement("div")
	packChildren(c, &ele, c.getContext())
	ele.Set("className", ele.Get("className").String()+" go-ui-column")
	return ele
}
func (c Column) Pack(ctx Context) Node {
	c._context.Context = ctx
	return c
}

type Row struct {
	Alignment Position
	Expand    bool
	Children  []Widget
	_context
}

func (r Row) getChildren() []Widget {
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
	packChildren(r, &ele, r.getContext())
	return ele
}
func (r Row) Pack(ctx Context) Node {
	r._context.Context = ctx
	return r
}

type Expanded struct {
	Flex  int
	Child Widget
	_context
}

func (e Expanded) getChildren() []Widget {

	return []Widget{e.Child}
}

func (e Expanded) pack() dom.JsDomElement {
	return Inline{
		Params: Params{
			Style: Style{
				FlexGrow: e.Flex,
			},
		},
		Children: []Widget{
			e.Child,
		},
	}.pack()
}

func (e Expanded) Pack(ctx Context) Node {
	e._context.Context = ctx
	return e
}

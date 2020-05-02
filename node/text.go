package node

import (
	"fmt"
	"github.com/TobiasYin/go_web_ui/dom"
	"github.com/TobiasYin/go_web_ui/node/color"
	"strconv"
	"strings"
)

type Text struct {
	Content   string
	TextStyle *TextStyle
}

func (t Text) pack() dom.JsDomElement {
	node := dom.Dom.CreateTextNode(t.Content)
	father := dom.Dom.CreateElement("span")
	father.AppendChild(node)
	if t.TextStyle != nil {
		father.Set("style", t.TextStyle.packStyle())
	}
	return father
}

type FontWeight int

const (
	FontWeightLighter = FontWeight(100)
	FontWeightLight   = FontWeight(200)
	FontWeightNormal  = FontWeight(400)
	FontWeightBold    = FontWeight(700)
	FontWeightBolder  = FontWeight(900)
	_                 = FontWeight(iota * 100)
	FontWeight100
	FontWeight200
	FontWeight300
	FontWeight400
	FontWeight500
	FontWeight600
	FontWeight700
	FontWeight800
	FontWeight900
)

type TextStyle struct {
	FontSize   int
	FontWeight FontWeight
	FontColor  color.Color
	FontStyle  FontStyle
	FontFamily string
	TextDecoration
}

type TextDecoration string

const (
	TextDecorationUnderline   TextDecoration = "underline"
	TextDecorationLineThrough TextDecoration = "line-through"
	TextDecorationOverline    TextDecoration = "overline"
)

type FontStyle string

const (
	FontStyleItalic FontStyle = "italic"
	FontStyleNormal FontStyle = "normal"
)

func (s TextStyle) packStyle() string {
	res := make(map[string]string)
	if s.FontWeight != 0 {
		res["font-weight"] = strconv.Itoa(int(s.FontWeight))
	}
	if s.FontSize != 0 {
		res["font-size"] = fmt.Sprintf("%dpx", s.FontSize)
	}
	res["color"] = s.FontColor.ToHex()
	if s.FontFamily != "" {
		res["font-family"] = s.FontFamily
	}
	if s.FontStyle != "" {
		res["font-style"] = string(s.FontStyle)
	}

	var r strings.Builder
	for k, v := range res {
		r.WriteString(fmt.Sprintf("%s:%s;", k, v))
	}
	return r.String()
}
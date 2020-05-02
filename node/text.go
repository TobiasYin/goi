package node

import (
	"fmt"
	"github.com/TobiasYin/go_web_ui/dom"
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
	FontColor  Color
}

func (s TextStyle) packStyle() string {
	res := make(map[string]string)
	if s.FontWeight != 0 {
		res["font-weight"] = strconv.Itoa(int(s.FontWeight))
	}
	if s.FontSize != 0 {
		res["font-size"] = fmt.Sprintf("%dpx", s.FontSize)
	}
	res["color"] = s.FontColor.ToHex()

	var r strings.Builder
	for k, v := range res {
		r.WriteString(fmt.Sprintf("%s:%s;", k, v))
	}
	fmt.Println(r.String())
	return r.String()
}

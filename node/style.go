package node

import (
	"fmt"
	"github.com/TobiasYin/go_web_ui/node/color"
	"strings"
)

type Style struct {
	Height                 Size
	Width                  Size
	TextAlign              Position
	TextIndent             int
	Color                  color.Color
	BackgroundColor        color.Color
	BackgroundImage        string
	BackgroundRepeat       BackgroundRepeat
	BackgroundPositionByPx int
	BackgroundPosition     Position
	Display                Display
	BoxShadow              []BoxShadow
	FlexGrow               int
	// TODO 填充CSS style属性
}

func (s Style) packStyle() string {
	res := make(map[string]string)
	if s.Height.Mode != SizeModeDefault || s.Height.Value != 0 {
		res["height"] = s.Height.String()
	}
	if s.Width.Mode != SizeModeDefault || s.Width.Value != 0 {
		res["width"] = s.Width.String()
	}
	if s.TextAlign != "" {
		res["text-align"] = string(s.TextAlign)
	}
	if s.TextIndent != 0 {
		res["text-indent"] = fmt.Sprintf("%dem", s.TextIndent)
	}
	if s.Color != color.Black {
		res["color"] = s.Color.String()
	}
	if s.BackgroundColor != color.Black {
		res["background-color"] = s.BackgroundColor.String()
	}
	if s.BackgroundImage != "" {
		res["background-image"] = s.BackgroundImage
	}
	if s.BackgroundRepeat != "" {
		res["background-repeat"] = string(s.BackgroundRepeat)
	}
	if s.BackgroundPositionByPx != 0 {
		res["background-position-by-px"] = fmt.Sprintf("%dpx", s.BackgroundPositionByPx)
	}
	if s.BackgroundPosition != "" {
		res["background-position"] = string(s.BackgroundPosition)
	}
	if s.Display != "" {
		res["display"] = string(s.Display)
	}
	if s.FlexGrow != 0 {
		res["flex-grow"] = fmt.Sprintf("%d", s.FlexGrow)
	}
	if len(s.BoxShadow) != 0 {
		var b strings.Builder
		for _, bs := range s.BoxShadow {
			b.WriteString(bs.String())
		}
		res["box-shadow"] = b.String()
	}
	var r strings.Builder
	for k, v := range res {
		r.WriteString(fmt.Sprintf("%s:%s;", k, v))
	}
	return r.String()
}

type Size struct {
	Mode  SizeMode
	Value int
}

func (s Size) String() string {
	if s.Mode == SizeModeDefault && s.Value != 0 {
		s.Mode = SizeModePx
	}
	if s.Mode == SizeModeDefault || s.Mode == SizeModeAuto {
		return string(SizeModeAuto)
	}
	return fmt.Sprintf("%d%v", s.Value, s.Mode)
}

type SizeMode string

const (
	SizeModeDefault SizeMode = ""
	SizeModeAuto    SizeMode = "auto"
	SizeModePx      SizeMode = "px"
	SizeModePercent SizeMode = "%"
)

type BoxShadow struct {
	XOffset      int
	YOffset      int
	BlurRadius   int
	SpreadRadius int
	Color        color.Color
}

func (b BoxShadow) String() string {
	return fmt.Sprintf("%d %d %d %d %s ", b.XOffset, b.YOffset, b.BlurRadius, b.SpreadRadius, b.Color.String())
}

type Position string

const (
	Right  Position = "right"
	Left   Position = "left"
	Center Position = "center"
	Top    Position = "top"
	Bottom Position = "bottom"
)

type BackgroundRepeat string

const (
	BackgroundRepeatRepeat   BackgroundRepeat = "repeat"
	BackgroundRepeatNoRepeat BackgroundRepeat = "no-repeat"
	BackgroundRepeatRepeatX  BackgroundRepeat = "repeat-x"
	BackgroundRepeatRepeatY  BackgroundRepeat = "repeat-y"
)

type Display string

const (
	DisplayInline      Display = "inline"
	DisplayBlock       Display = "block"
	DisplayInlineBlock Display = "inline-block"
)

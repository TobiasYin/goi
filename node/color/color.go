package color

import (
	"fmt"
)

type Color struct {
	r byte
	g byte
	b byte
	a float32
}

func New(r int, g int, b int) Color {
	return Color{
		byte(r),
		byte(g),
		byte(b),
		0,
	}
}
func NewWithAlpha(r int, g int, b int, a int) Color {
	alpha := float32(a) / 255
	if alpha < 0 {
		alpha = 0
	}
	if alpha > 1 {
		alpha = 1
	}
	return Color{
		byte(r),
		byte(g),
		byte(b),
		1 - alpha,
	}
}

func (c Color) ToHex() string {
	return fmt.Sprintf("#%.2x%.2x%.2x", c.r, c.g, c.b)
}

func (c Color) ToRGBA() string {
	if c.a != 0 {
		return fmt.Sprintf("rgba(%d, %d, %d, %v)", c.r, c.g, c.b, 1-c.a)
	} else {
		return fmt.Sprintf("rgb(%d, %d, %d)", c.r, c.g, c.b)
	}
}

func (c Color) String() string {
	return c.ToRGBA()
}

var (
	LightPink            = New(255, 182, 193)
	Pink                 = New(255, 192, 203)
	Crimson              = New(220, 20, 60)
	LavenderBlush        = New(255, 240, 245)
	PaleVioletRed        = New(219, 112, 147)
	HotPink              = New(255, 105, 180)
	DeepPink             = New(255, 20, 147)
	MediumVioletRed      = New(199, 21, 133)
	Orchid               = New(218, 112, 214)
	Thistle              = New(216, 191, 216)
	plum                 = New(221, 160, 221)
	Violet               = New(238, 130, 238)
	Magenta              = New(255, 0, 255)
	Fuchsia              = New(255, 0, 255)
	DarkMagenta          = New(139, 0, 139)
	Purple               = New(128, 0, 128)
	MediumOrchid         = New(186, 85, 211)
	DarkVoilet           = New(148, 0, 211)
	DarkOrchid           = New(153, 50, 204)
	Indigo               = New(75, 0, 130)
	BlueViolet           = New(138, 43, 226)
	MediumPurple         = New(147, 112, 219)
	MediumSlateBlue      = New(123, 104, 238)
	SlateBlue            = New(106, 90, 205)
	DarkSlateBlue        = New(72, 61, 139)
	Lavender             = New(230, 230, 250)
	GhostWhite           = New(248, 248, 255)
	Blue                 = New(0, 0, 255)
	MediumBlue           = New(0, 0, 205)
	MidnightBlue         = New(25, 25, 112)
	DarkBlue             = New(0, 0, 139)
	Navy                 = New(0, 0, 128)
	RoyalBlue            = New(65, 105, 225)
	CornflowerBlue       = New(100, 149, 237)
	LightSteelBlue       = New(176, 196, 222)
	LightSlateGray       = New(119, 136, 153)
	SlateGray            = New(112, 128, 144)
	DoderBlue            = New(30, 144, 255)
	AliceBlue            = New(240, 248, 255)
	SteelBlue            = New(70, 130, 180)
	LightSkyBlue         = New(135, 206, 250)
	SkyBlue              = New(135, 206, 235)
	DeepSkyBlue          = New(0, 191, 255)
	LightBLue            = New(173, 216, 230)
	PowDerBlue           = New(176, 224, 230)
	CadetBlue            = New(95, 158, 160)
	Azure                = New(240, 255, 255)
	LightCyan            = New(225, 255, 255)
	PaleTurquoise        = New(175, 238, 238)
	Cyan                 = New(0, 255, 255)
	Aqua                 = New(0, 255, 255)
	DarkTurquoise        = New(0, 206, 209)
	DarkSlateGray        = New(47, 79, 79)
	DarkCyan             = New(0, 139, 139)
	Teal                 = New(0, 128, 128)
	MediumTurquoise      = New(72, 209, 204)
	LightSeaGreen        = New(32, 178, 170)
	Turquoise            = New(64, 224, 208)
	Auqamarin            = New(127, 255, 170)
	MediumAquamarine     = New(0, 250, 154)
	MediumSpringGreen    = New(245, 255, 250)
	MintCream            = New(0, 255, 127)
	SpringGreen          = New(60, 179, 113)
	SeaGreen             = New(46, 139, 87)
	Honeydew             = New(240, 255, 240)
	LightGreen           = New(144, 238, 144)
	PaleGreen            = New(152, 251, 152)
	DarkSeaGreen         = New(143, 188, 143)
	LimeGreen            = New(50, 205, 50)
	Lime                 = New(0, 255, 0)
	ForestGreen          = New(34, 139, 34)
	Green                = New(0, 128, 0)
	DarkGreen            = New(0, 100, 0)
	Chartreuse           = New(127, 255, 0)
	LawnGreen            = New(124, 252, 0)
	GreenYellow          = New(173, 255, 47)
	OliveDrab            = New(85, 107, 47)
	Beige                = New(107, 142, 35)
	LightGoldenrodYellow = New(250, 250, 210)
	Ivory                = New(255, 255, 240)
	LightYellow          = New(255, 255, 224)
	Yellow               = New(255, 255, 0)
	Olive                = New(128, 128, 0)
	DarkKhaki            = New(189, 183, 107)
	LemonChiffon         = New(255, 250, 205)
	PaleGodenrod         = New(238, 232, 170)
	Khaki                = New(240, 230, 140)
	Gold                 = New(255, 215, 0)
	Cornislk             = New(255, 248, 220)
	GoldEnrod            = New(218, 165, 32)
	FloralWhite          = New(255, 250, 240)
	OldLace              = New(253, 245, 230)
	Wheat                = New(245, 222, 179)
	Moccasin             = New(255, 228, 181)
	Orange               = New(255, 165, 0)
	PapayaWhip           = New(255, 239, 213)
	BlanchedAlmond       = New(255, 235, 205)
	NavajoWhite          = New(255, 222, 173)
	AntiqueWhite         = New(250, 235, 215)
	Tan                  = New(210, 180, 140)
	BrulyWood            = New(222, 184, 135)
	Bisque               = New(255, 228, 196)
	DarkOrange           = New(255, 140, 0)
	Linen                = New(250, 240, 230)
	Peru                 = New(205, 133, 63)
	PeachPuff            = New(255, 218, 185)
	SandyBrown           = New(244, 164, 96)
	Chocolate            = New(210, 105, 30)
	SaddleBrown          = New(139, 69, 19)
	SeaShell             = New(255, 245, 238)
	Sienna               = New(160, 82, 45)
	LightSalmon          = New(255, 160, 122)
	Coral                = New(255, 127, 80)
	OrangeRed            = New(255, 69, 0)
	DarkSalmon           = New(233, 150, 122)
	Tomato               = New(255, 99, 71)
	MistyRose            = New(255, 228, 225)
	Salmon               = New(250, 128, 114)
	Snow                 = New(255, 250, 250)
	LightCoral           = New(240, 128, 128)
	RosyBrown            = New(188, 143, 143)
	IndianRed            = New(205, 92, 92)
	Red                  = New(255, 0, 0)
	Brown                = New(165, 42, 42)
	FireBrick            = New(178, 34, 34)
	DarkRed              = New(139, 0, 0)
	Maroon               = New(128, 0, 0)
	White                = New(255, 255, 255)
	WhiteSmoke           = New(245, 245, 245)
	Gainsboro            = New(220, 220, 220)
	LightGrey            = New(211, 211, 211)
	Silver               = New(192, 192, 192)
	DarkGray             = New(169, 169, 169)
	Gray                 = New(128, 128, 128)
	DimGray              = New(105, 105, 105)
	Black                = New(0, 0, 0)
)

package termui

// Color is an integer in the range -1 to 255.
type Color int

// ColorDefault = clear
const ColorDefault = -1

// Copied from termbox. Attributes that can be bitwise OR'ed with a color.
const (
	AttrBold Color = 1 << (iota + 9)
	AttrUnderline
	AttrReverse
)

// Theme is assigned to the current theme.
var Theme = DefaultTheme

// DefaultTheme implements a generic set of colors to use by default.
var DefaultTheme = Colorscheme{
	Fg: 7,
	Bg: -1,

	LabelFg:  7,
	LabelBg:  -1,
	BorderFg: 6,
	BorderBg: -1,

	Sparkline:   4,
	LineGraph:   0,
	TableCursor: 4,
	GaugeColor:  7,
}

// A Colorscheme represents the current look-and-feel of the dashboard.
type Colorscheme struct {
	Fg Color
	Bg Color

	LabelFg  Color
	LabelBg  Color
	BorderFg Color
	BorderBg Color

	Sparkline   Color
	LineGraph   Color
	TableCursor Color
	GaugeColor  Color
}

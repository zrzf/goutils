package color

import (
	"fmt"
)

const Esc string = "\033["

const (
	End = Esc + string('0'+iota/10%10) + string('0'+iota/1%10) + "m"
	Bold
	Faint
	Italic
	Underline
	SlowBlink
	RapidBlink
	Invert
	Hide
	Strike


	//----------------------Fonts----------
	Font0		// Default font
	Font1		// Alternative fonts
	Font2
	Font3
	Font4
	Font5
	Font6
	Font7
	Font8
	Font9
	FontGothic
	//-------------------------------------


	DoubleUnderline
	NormalIntensity
	NoItalic
	NoUnderlined
	NoBlinking
	Spacing
	NoReversed
	Reveal
	NoStrike


	//----------------------Foregrounds----
	Black
	Red
	Green
	Yellow
	Blue
	Magenta
	Cyan
	White
	_	//RGB
	DefaultColor
	//-------------------------------------


	//----------------------Backgrounds----
	BBlack
	BRed
	BGreen
	BYellow
	BBlue
	BMagenta
	BCyan
	BWhite
	_	//BRGB
	DefaultBColor
	//-------------------------------------


	NoSpacing
	Frame
	Encircle
	Overline
	NoFrame
	NoOverline
	_	//URGB
	DefaultUnderlineColor
	IdeogramUnderline
	IdeogramDoubleUnderline
	IdeogramOverline
	IdeogramDoubleOverline
	IdeogramStressMarking
	ResetIdeogram
)

const (
	Superscript = Esc + string('0'+(iota+73)/10%10) + string('0'+(iota+73)/1%10) + "m"
	Subscript
)

const (
	//--------------Bright Foregrounds-----
	BrightBlack = Esc + string('0'+(iota+90)/100%10) + string('0'+(iota+90)/10%10) + string('0'+(iota+90)/1%10) + "m"
	BrightRed
	BrightGreen
	BrightYellow
	BrightBlue
	BrightMagenta
	BrightCyan
	BrightWhite
	_
	_
	//-------------------------------------

	//--------------Bright Backgrounds-----
	BBrightBlack
	BBrightRed
	BBrightGreen
	BBrightYellow
	BBrightBlue
	BBrightMagenta
	BBrightCyan
	BBrightWhite
	//-------------------------------------
)


// Aliases
const (
	Reset = End
	Normal = End
	Font10 = FontGothic
	Blackletter = FontGothic
	NoBold = DoubleUnderline
	NoBlackletter = NoItalic
	NoEncircle = NoFrame
)


func F8bit(n int) string {
	return fmt.Sprintf("%s38;5;%dm",Esc,n)
}

func B8bit(n int) string {
	return fmt.Sprintf("%s48;5;%dm",Esc,n)
}

func U8bit(n int) string {
	return fmt.Sprintf("%s58;5;%dm",Esc,n)
}

func RGB(r int, g int, b int) string {
	return fmt.Sprintf("%s38;2;%d;%d;%dm",Esc,r,g,b)
}

func BRGB(r int, g int, b int) string {
	return fmt.Sprintf("%s48;2;%d;%d;%dm",Esc,r,g,b)
}

func URGB(r int, g int, b int) string {
	return fmt.Sprintf("%s58;2;%d;%d;%dm",Esc,r,g,b)
}

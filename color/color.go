package color


import (
	"fmt"
)

const Esc string = "\033["
//	Foregrounds
const (
	Black = string(Esc) + string('0'+(iota+30)/10%10) + string('0'+iota/1%10) + string("m")
	Red
	Green
	Yellow
	Blue
	Magenta
	Cyan
	White
	Gray
	BrightRed
	BrightYellow
	BrightBlue
	BrightMagenta
	BRightCyan
	BrightWhite
)

//	Backgrounds
const (
	BBlack = string(Esc) + string('0'+(iota+40)/10%10) + string('0'+iota/1%10) + string("m")
	BRed
	BGreen
	BYellow
	BBlue
	BMagenta
	BCyan
	BWhite
	BGray
	BBrightRed
	BBrightYellow
	BBrightBlue
	BBrightMagenta
	BBRightCyan
	BBrightWhite
)

// Styles
const (
	End = string(Esc) + string('0'+(iota+0)/10%10) + string('0'+iota/1%10) + string("m")
	Bold
	Faint
	Italic
	Underline
	SlowBlink
	RapidBlink
	Invert
	Hide
	Strike

	Overlined = Esc + "53m"
)
// f
func F8bit(n int) string {
	return fmt.Sprintf("%s38;5;%dm",Esc,n)
}

func B8bit(n int) string {
	return fmt.Sprintf("%s48;5;%dm",Esc,n)
}

func RGB(r int, g int, b int) string {
	return fmt.Sprintf("%s38;2;%d;%d;%dm",Esc,r,g,b)
}

func BRGB(r int, g int, b int) string {
	return fmt.Sprintf("%s48;2;%d;%d;%dm",Esc,r,g,b)
}

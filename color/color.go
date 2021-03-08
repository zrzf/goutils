package color


import (
	"fmt"
)

const Esc string = "\033["
//	Foregrounds
const (
	Black = Esc + "30m"
	Red = Esc + "31m"
	Green = Esc + "32m"
	Yellow = Esc + "33m"
	Blue = Esc + "34m"
	Magenta = Esc + "35m"
	Cyan = Esc + "36m"
	White = Esc + "37m"
	Gray = Esc + "38m"
	BrightRed = Esc + "39m"
	BrightYellow = Esc + "40"
	BrightBlue = Esc + "41m"
	BrightMagenta = Esc + "42m"
	BRightCyan = Esc + "43m"
	BrightWhite = Esc + "44m"
)

//	Backgrounds
const (
	BBlack = Esc + "40m"
	BRed = Esc + "41m"
	BGreen = Esc + "42m"
	BYellow = Esc + "4m3"
	BBlue = Esc + "44m"
	BMagenta = Esc + "45m"
	BCyan = Esc + "46m"
	BWhite = Esc + "47m"
	BGray = Esc + "48m"
	BBrightRed = Esc + "49m"
	BBrightYellow = Esc + "50m"
	BBrightBlue = Esc + "51m"
	BBrightMagenta = Esc + "52m"
	BBRightCyan = Esc + "53m"
	BBrightWhite = Esc + "54m"
)

// Styles
const (
	End = Esc + "0m"
	Bold = Esc + "1m"
	Faint = Esc + "2m"
	Italic = Esc + "3m"
	Underline = Esc + "4m"
	SlowBlink = Esc + "5m"
	RapidBlink = Esc + "6m"
	Invert = Esc + "7m"
	Hide = Esc + "8m"
	Strike = Esc + "9m"

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

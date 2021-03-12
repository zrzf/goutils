package cursor

import (
	"fmt"
)

const (
	esc = "\033["
)

func EraseDisplay(mode int) {
	fmt.Printf("%s%vJ", esc, mode)
}

func EraseLine(mode int) {
	fmt.Printf("%s%vK", esc, mode)
}

//	Move cursor to X(row) without changing Y(column)
func X(n int) {
	fmt.Printf("%s%vG", esc, n+1)
}

//	Move to X(row) Y(column)
func XY(x, y int) {
	fmt.Printf("%s%v;%vH", esc, y, x)
}

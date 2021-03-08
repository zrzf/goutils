package strings

import (
	"strings"
)

func RemoveBetweenRune(text string, d1 rune, d2 rune, separator string) string {

	var ignore bool = false
	builder := strings.Builder{}
	//builder.Grow(len(text))
	for _,v := range text {
		switch v {
		case d1:
			ignore=true
			continue
		case d2:
			ignore=false
			builder.WriteString(separator)
			continue
		default:
			if !ignore{
				builder.WriteRune(v)
			}
		}

	}
	return builder.String()
}

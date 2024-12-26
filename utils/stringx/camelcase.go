package stringx

import (
	"unicode"
)

func CapitalizeFirstLetter(s string) string {
	if s == "" {
		return ""
	}

	aux := []rune(s)
	aux[0] = unicode.ToUpper(aux[0])
	return string(aux)
}

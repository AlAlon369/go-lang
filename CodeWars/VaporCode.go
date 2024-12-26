package CodeWars

import (
	"strings"
	"unicode"
)

func Vaporcode(s string) string {
	var sb strings.Builder

	for _, r := range s {
		if !unicode.IsSpace(r) {
			sb.WriteRune(unicode.ToUpper(r))
			sb.WriteString("  ")
		}
	}

	result := sb.String()
	if len(result) > 0 {
		return result[:len(result)-2]
	}
	return result
}

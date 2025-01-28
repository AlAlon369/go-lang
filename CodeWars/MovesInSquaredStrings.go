package CodeWars

import (
	"strings"
)

// VertMirror переворачивает каждую строку горизонтально
func VertMirror(s string) string {
	lines := strings.Split(s, "\n")
	for i, line := range lines {
		lines[i] = reverseString(line)
	}
	return strings.Join(lines, "\n")
}

// HorMirror переворачивает строки вертикально
func HorMirror(s string) string {
	lines := strings.Split(s, "\n")
	n := len(lines)
	for i := 0; i < n/2; i++ {
		lines[i], lines[n-i-1] = lines[n-i-1], lines[i]
	}
	return strings.Join(lines, "\n")
}

// reverseString переворачивает строку
func reverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

// Oper применяет функцию f к строке s
type FParam func(string) string

func Oper(f FParam, x string) string {
	return f(x)
}

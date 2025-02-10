package main

import (
	"Golang/CodeWars"
	"fmt"
)

// Тесты для проверки работы функции Triangle
func main() {
	tests := []struct {
		input    string
		expected rune
	}{
		{"RRGBRGBB", 'G'},
		{"RGB", 'B'},
		{"R", 'R'},
		{"GG", 'G'},
		{"GB", 'R'},
		{"BGR", 'G'},
	}

	for i, test := range tests {
		result := CodeWars.Triangle(test.input)
		fmt.Printf("Test %d: Input: %s, Expected: %c, Got: %c\n",
			i+1, test.input, test.expected, result)

		if result != test.expected {
			fmt.Printf("❌ Test %d failed!\n\n", i+1)
		} else {
			fmt.Printf("✅ Test %d passed!\n\n", i+1)
		}
	}
}

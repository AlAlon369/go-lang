package main

import (
	"Golang/CodeWars"
	"fmt"
)

func main() {
	tests := []struct {
		input    string
		expected rune
	}{
		{"a", 'a'},
		{"ab", 'a'},
		{"axyzxyz", 'x'},
		{"abcabc", 'a'},
		{"zzzabczzz", 'z'},
		{"abcdef", 'a'},
		{"babaa", 'a'},
	}

	for i, test := range tests {
		result := CodeWars.SolveString(test.input)
		fmt.Printf("Test %d: Input: %s, Expected: %c, Got: %c\n",
			i+1, test.input, test.expected, result)

		if result != test.expected {
			fmt.Printf("❌ Test %d failed!\n\n", i+1)
		} else {
			fmt.Printf("✅ Test %d passed!\n\n", i+1)
		}
	}
}

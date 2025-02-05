package main

import (
	"Golang/CodeWars"
	"fmt"
)

func main() {
	// Набор тестов
	tests := []struct {
		input    string
		expected int
	}{
		{"agrtertyfikfmroyrntbvsukldkfa", 6},
		{"erfaiekjudhyfimngukduo", 4},
		{"aeiou", 5},
		{"uoiea", 1},
		{"abcdeioua", 5},
		{"xyz", 0},
		{"aaeebbiioouu", 5},
	}

	// Запуск тестов
	for i, test := range tests {
		result := CodeWars.GetTheVowels(test.input)
		fmt.Printf("Test %d: Input: %q, Expected: %d, Got: %d\n",
			i+1, test.input, test.expected, result)

		if result != test.expected {
			fmt.Printf("❌ Test %d failed!\n\n", i+1)
		} else {
			fmt.Printf("✅ Test %d passed!\n\n", i+1)
		}
	}
}

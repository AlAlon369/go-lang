package main

import (
	"Golang/CodeWars"
	"fmt"
)

func main() {
	// Набор тестов для функции Strong
	tests := []struct {
		input    int
		expected string
	}{
		{1, "STRONG!!!!"},      // 1! = 1
		{145, "STRONG!!!!"},    // 1! + 4! + 5! = 145
		{123, "Not Strong !!"}, // 1! + 2! + 3! = 9 != 123
		{150, "Not Strong !!"}, // 1! + 5! + 0! = 122 != 150
		{2, "STRONG!!!!"},      // 2! = 2
		{40585, "STRONG!!!!"},  // 4! + 0! + 5! + 8! + 5! = 40585
	}

	// Запуск тестов
	for i, test := range tests {
		result := CodeWars.Strong(test.input)
		fmt.Printf("Test %d: Input: %d, Expected: %s, Got: %s\n", i+1, test.input, test.expected, result)
		if result != test.expected {
			fmt.Printf("❌ Test %d failed!\n\n", i+1)
		} else {
			fmt.Printf("✅ Test %d passed!\n\n", i+1)
		}
	}
}

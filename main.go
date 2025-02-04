package main

import (
	"Golang/CodeWars"
	"fmt"
)

func main() {
	// Набор тестов
	tests := []struct {
		input    string
		expected string
	}{
		{"our code", "edo cruo"},
		{"your code rocks", "skco redo cruoy"},
		{"codewars", "srawedoc"},
		{"a b c", "c b a"},
		{"ab cd ef", "fe dc ba"},
		{"   a   ", "   a   "}, // Тест с пробелами в начале и конце
	}

	// Запуск тестов
	for i, test := range tests {
		result := CodeWars.Solved(test.input)
		fmt.Printf("Test %d: Input: %q, Expected: %q, Got: %q\n",
			i+1, test.input, test.expected, result)

		if result != test.expected {
			fmt.Printf("❌ Test %d failed!\n\n", i+1)
		} else {
			fmt.Printf("✅ Test %d passed!\n\n", i+1)
		}
	}
}

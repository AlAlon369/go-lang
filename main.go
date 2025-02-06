package main

import (
	"Golang/CodeWars"
	"fmt"
)

func main() {
	// Набор тестов
	tests := []struct {
		a0       int
		a        int
		p        float64
		expected string
	}{
		{100, 101, 0.98, "2017-01-01"},
		{100, 150, 2.00, "2035-12-26"},
		{5000, 7000, 1.50, "2031-10-13"},
		{100, 200, 1.00, "2047-01-01"},
	}

	// Запуск тестов
	for i, test := range tests {
		result := CodeWars.DateNbDays(test.a0, test.a, test.p)
		fmt.Printf("Test %d: a0=%d, a=%d, p=%.2f, Expected: %s, Got: %s\n",
			i+1, test.a0, test.a, test.p, test.expected, result)

		if result != test.expected {
			fmt.Printf("❌ Test %d failed!\n\n", i+1)
		} else {
			fmt.Printf("✅ Test %d passed!\n\n", i+1)
		}
	}
}

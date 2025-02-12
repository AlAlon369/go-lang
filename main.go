package main

import (
	"Golang/CodeWars"
	"fmt"
)

func main() {
	tests := []struct {
		year     int
		expected bool
	}{
		{1600, true},  // Високосный (делится на 400)
		{1700, false}, // Не високосный (делится на 100, но не на 400)
		{2000, true},  // Високосный (делится на 400)
		{2024, true},  // Високосный (делится на 4, но не на 100)
		{2023, false}, // Не високосный (не делится на 4)
	}

	for i, test := range tests {
		result := CodeWars.IsLeapYear(test.year)
		fmt.Printf("Test %d: Year: %d, Expected: %t, Got: %t\n",
			i+1, test.year, test.expected, result)

		if result != test.expected {
			fmt.Printf("❌ Test %d failed!\n\n", i+1)
		} else {
			fmt.Printf("✅ Test %d passed!\n\n", i+1)
		}
	}
}

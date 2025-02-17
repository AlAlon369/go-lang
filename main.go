package main

import (
	"Golang/CodeWars"
	"fmt"
)

func main() {
	tests := []struct {
		digits   int
		expected int
	}{
		{1, 3}, // sqrt(3^2) = 3
		{3, 6}, // sqrt(3^2 + 1^2 + 4^2) = 5.099 -> 6
		{5, 8}, // sqrt(3^2 + 1^2 + 4^2 + 1^2 + 5^2) = 7.416 -> 8
	}

	for i, test := range tests {
		result := CodeWars.SquarePi(test.digits)
		fmt.Printf("Test %d: Digits: %d\nExpected: %d\nGot: %d\n",
			i+1, test.digits, test.expected, result)

		if result != test.expected {
			fmt.Printf("❌ Test %d failed!\n\n", i+1)
		} else {
			fmt.Printf("✅ Test %d passed!\n\n", i+1)
		}
	}
}

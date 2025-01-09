package main

import (
	"Golang/CodeWars"
	"fmt"
)

func main() {
	// Тестовые кейсы
	testCases := []struct {
		a, b, c, d string
		expected   int
	}{
		{"white", "black", "white", "black", 1},
		{"white", "black", "black", "white", 2},
		{"black", "black", "white", "white", 1},
		{"black", "white", "white", "black", 2},
		{"white", "white", "black", "black", 1},
	}

	// Прогон тестов
	for i, tc := range testCases {
		result := CodeWars.GuessHatColor(tc.a, tc.b, tc.c, tc.d)
		fmt.Printf("Test Case %d: a=%q, b=%q, c=%q, d=%q\n", i+1, tc.a, tc.b, tc.c, tc.d)
		fmt.Printf("  Expected: %d, Got: %d\n", tc.expected, result)

		if result == tc.expected {
			fmt.Println("  ✅ Test Passed")
		} else {
			fmt.Println("  ❌ Test Failed")
		}
		fmt.Println("-----------------------------")
	}
}

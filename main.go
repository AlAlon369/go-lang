package main

import (
	"Golang/CodeWars"
	"fmt"
)

func main() {
	tests := []struct {
		input    []int
		expected []int
	}{
		{[]int{}, []int{}},
		{[]int{1}, []int{1}},
		{[]int{2, 1}, []int{1, 2, 1}},
		{[]int{1, 3, 2}, []int{1, 2, 3, 2, 1}},
		{[]int{-8, 42, 18, 0, -16}, []int{-16, -8, 0, 18, 42, 18, 0, -8, -16}},
		{[]int{-3, 15, 8, -1, 7, -1}, []int{-3, -1, -1, 7, 8, 15, 8, 7, -1, -1, -3}},
		{[]int{-5, 10, 8, 10, 2, -3, 10}, []int{-5, -3, 2, 8, 10, 10, 10, 10, 10, 8, 2, -3, -5}},
	}

	for i, test := range tests {
		result := CodeWars.Mirror(test.input)
		fmt.Printf("Test %d: Input: %v\nExpected: %v\nGot: %v\n\n",
			i+1, test.input, test.expected, result)

		if fmt.Sprint(result) != fmt.Sprint(test.expected) {
			fmt.Printf("❌ Test %d failed!\n\n", i+1)
		} else {
			fmt.Printf("✅ Test %d passed!\n\n", i+1)
		}
	}
}

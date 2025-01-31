package main

import (
	"Golang/CodeWars"
	"fmt"
)

func main() {
	// Набор тестов для функции SolveTask
	testsSolveTask := []struct {
		arr      []int
		expected int
	}{
		{[]int{1, -1, 2, -2, 3}, 3},
		{[]int{-3, 1, 2, 3, -1, -4, -2}, -4},
		{[]int{1, -1, 2, -2, 3, 3}, 3},
		{[]int{-5, 5, 6, -6, 7}, 7},
		{[]int{8, -8, 9, -9, 10, 10}, 10},
		{[]int{1, -1, 2, -2, 3, -3, 4, -4, 5}, 5},
	}

	// Запуск тестов для функции SolveTask
	for i, test := range testsSolveTask {
		result := CodeWars.SolveTask(test.arr)
		fmt.Printf("Test %d (SolveTask): Input: arr=%v, Expected: %d, Got: %d\n",
			i+1, test.arr, test.expected, result)
		if result != test.expected {
			fmt.Printf("❌ Test %d failed!\n\n", i+1)
		} else {
			fmt.Printf("✅ Test %d passed!\n\n", i+1)
		}
	}
}

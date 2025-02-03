package main

import (
	"Golang/CodeWars"
	"fmt"
)

func main() {
	// Набор тестов для функции StantonMeasure
	tests := []struct {
		arr      []int
		expected int
	}{
		{[]int{1, 4, 3, 2, 1, 2, 3, 2}, 3},
		{[]int{1, 4, 1, 2, 11, 2, 3, 1}, 1},
		{[]int{1, 1, 1, 1}, 4},
		{[]int{2, 2, 2, 2}, 0}, // 1 нет в массиве -> n = 0, а 0 нет -> возвращаем 0
		{[]int{1, 1, 2, 2, 2, 3, 3, 3, 3}, 3},
	}

	// Запуск тестов
	for i, test := range tests {
		result := CodeWars.StantonMeasure(test.arr)
		fmt.Printf("Test %d: Input: %v, Expected: %d, Got: %d\n",
			i+1, test.arr, test.expected, result)

		if result != test.expected {
			fmt.Printf("❌ Test %d failed!\n\n", i+1)
		} else {
			fmt.Printf("✅ Test %d passed!\n\n", i+1)
		}
	}
}

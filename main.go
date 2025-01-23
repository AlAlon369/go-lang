package main

import (
	"Golang/CodeWars"
	"fmt"
)

func main() {
	// Набор тестов для проверки PairZeros
	tests := []struct {
		input    []uint64
		expected []uint64
	}{
		{[]uint64{0, 0, 0}, []uint64{0, 0}},       // Несколько нулей подряд
		{[]uint64{1, 0, 1, 0}, []uint64{1, 0, 1}}, // Чередование чисел и нулей
		{[]uint64{}, []uint64{}},                  // Пустой массив
		{[]uint64{1, 2, 3}, []uint64{1, 2, 3}},    // Без нулей
		{[]uint64{0, 0}, []uint64{0}},             // Только пара нулей
	}

	// Запуск тестов
	for i, test := range tests {
		result := CodeWars.PairZeros(test.input)
		fmt.Printf("Test %d: Input: %v, Expected: %v, Got: %v\n", i+1, test.input, test.expected, result)
		if fmt.Sprintf("%v", result) != fmt.Sprintf("%v", test.expected) {
			fmt.Printf("❌ Test %d failed!\n\n", i+1)
		} else {
			fmt.Printf("✅ Test %d passed!\n\n", i+1)
		}
	}
}

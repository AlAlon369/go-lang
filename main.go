package main

import (
	"Golang/CodeWars" // Убедитесь, что путь импорта совпадает с вашей структурой проекта
	"fmt"
)

func main() {
	// Тестовые кейсы
	testCases := []struct {
		divisor int
		bound   int
	}{
		{2, 7},    // Пример из задачи
		{10, 50},  // Пример из задачи
		{37, 200}, // Пример из задачи
		{3, 10},   // Дополнительный тест
	}

	// Прогон тестов
	for i, testCase := range testCases {
		result := CodeWars.MaxMultiple(testCase.divisor, testCase.bound)
		fmt.Printf("Test Case %d: Divisor=%d, Bound=%d\n", i+1, testCase.divisor, testCase.bound)
		fmt.Printf("  Result: %d\n", result)
		fmt.Println("-----------------------------")
	}
}

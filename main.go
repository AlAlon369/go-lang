package main

import (
	"Golang/CodeWars"
	"fmt"
)

func main() {
	// Тестовые кейсы
	testCases := [][]string{
		{"abode", "ABc", "xyzD"},     // Пример из задачи
		{"hello", "world", "GoLang"}, // Другие слова
		{"abc", "def", "ghi", "jkl"}, // Простые тесты
		{"", "ABCDEF", "qrstuv"},     // Пустая строка и строки разной длины
	}

	// Прогон тестов
	for i, testCase := range testCases {
		result := CodeWars.SolveIt(testCase) // Используем правильное имя функции
		fmt.Printf("Test Case %d: %v\n", i+1, testCase)
		fmt.Printf("  Result: %v\n", result)
		fmt.Println("-----------------------------")
	}
}

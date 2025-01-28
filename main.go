package main

import (
	"Golang/CodeWars"
	"fmt"
)

func main() {
	// Набор тестов для функции Strong
	testsStrong := []struct {
		input    int
		expected string
	}{
		{1, "STRONG!!!!"},      // 1! = 1
		{145, "STRONG!!!!"},    // 1! + 4! + 5! = 145
		{123, "Not Strong !!"}, // 1! + 2! + 3! = 9 != 123
		{150, "Not Strong !!"}, // 1! + 5! + 0! = 122 != 150
		{2, "STRONG!!!!"},      // 2! = 2
		{40585, "STRONG!!!!"},  // 4! + 0! + 5! + 8! + 5! = 40585
	}

	// Запуск тестов для функции Strong
	for i, test := range testsStrong {
		result := CodeWars.Strong(test.input)
		fmt.Printf("Test %d (Strong): Input: %d, Expected: %s, Got: %s\n", i+1, test.input, test.expected, result)
		if result != test.expected {
			fmt.Printf("❌ Test %d failed!\n\n", i+1)
		} else {
			fmt.Printf("✅ Test %d passed!\n\n", i+1)
		}
	}

	// Набор тестов для функции VertMirror и HorMirror
	testsStrings := []struct {
		input    string
		f        CodeWars.FParam
		expected string
	}{
		{"abcd\nefgh\nijkl\nmnop", CodeWars.VertMirror, "dcba\nhgfe\nlkji\nponm"},
		{"abcd\nefgh\nijkl\nmnop", CodeWars.HorMirror, "mnop\nijkl\nefgh\nabcd"},
	}

	// Запуск тестов для функций VertMirror и HorMirror
	for i, test := range testsStrings {
		result := CodeWars.Oper(test.f, test.input)
		fmt.Printf("Test %d (Strings): Input: %q, Expected: %q, Got: %q\n", i+1, test.input, test.expected, result)
		if result != test.expected {
			fmt.Printf("❌ Test %d failed!\n\n", i+1)
		} else {
			fmt.Printf("✅ Test %d passed!\n\n", i+1)
		}
	}
}

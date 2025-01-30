package main

import (
	"Golang/CodeWars"
	"fmt"
)

func main() {
	// Набор тестов для функции ContainAllRots
	testsContainAllRots := []struct {
		strng    string
		arr      []string
		expected bool
	}{
		{"bsjq", []string{"bsjq", "qbsj", "sjqb", "twZNsslC", "jqbs"}, true},
		{"Ajylvpy", []string{"Ajylvpy", "ylvpyAj", "jylvpyA", "lvpyAjy", "pyAjylv", "vpyAjyl", "ipywee"}, false},
		{"", []string{"anything"}, true},
		{"abc", []string{"abc", "cab", "bca"}, true},
		{"xyz", []string{"xzy", "yzx", "zxy"}, false},
	}

	// Запуск тестов для функции ContainAllRots
	for i, test := range testsContainAllRots {
		result := CodeWars.ContainAllRots(test.strng, test.arr)
		fmt.Printf("Test %d (ContainAllRots): Input: strng=%q, arr=%v, Expected: %v, Got: %v\n",
			i+1, test.strng, test.arr, test.expected, result)
		if result != test.expected {
			fmt.Printf("❌ Test %d failed!\n\n", i+1)
		} else {
			fmt.Printf("✅ Test %d passed!\n\n", i+1)
		}
	}
}

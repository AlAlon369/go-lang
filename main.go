package main

import (
	"Golang/CodeWars"
	"fmt"
)

func main() {
	// Набор тестов для функции Compose
	testsCompose := []struct {
		s1       string
		s2       string
		expected string
	}{
		{
			"abcd\nefgh\nijkl\nmnop",
			"qrst\nuvwx\nyz12\n3456",
			"a3456\nefyz1\nijkuv\nmnopq",
		},
		{
			"012\n345\n678",
			"9ab\ncde\nfgh",
			"0gh\n34f\n67e",
		},
		// Добавьте больше тестов по мере необходимости
	}

	// Запуск тестов для функции Compose
	for i, test := range testsCompose {
		result := CodeWars.Compose(test.s1, test.s2)
		fmt.Printf("Test %d (Compose): Input: s1=%q, s2=%q, Expected: %q, Got: %q\n", i+1, test.s1, test.s2, test.expected, result)
		if result != test.expected {
			fmt.Printf("❌ Test %d failed!\n\n", i+1)
		} else {
			fmt.Printf("✅ Test %d passed!\n\n", i+1)
		}
	}
}

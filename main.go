package main

import (
	"Golang/CodeWars" // Импортируем пакет CodeWars
	"fmt"
)

func main() {
	// Примеры для проверки функции EvenOrOdd
	testCases := []int{1, 3, 100, 1298734, 42, 7}

	for _, n := range testCases {
		result := CodeWars.EvenOrOdd(n)
		fmt.Printf("EvenOrOdd(%d) = %s\n", n, result)
	}
}

package main

import (
	"Golang/CodeWars" // Импортируем пакет CodeWars
	"fmt"
)

func main() {
	// Примеры для проверки функции Digitize
	testCases := []int{35231, 0, 1, 3, 100, 1298734, 42, 7}

	for _, n := range testCases {
		result := CodeWars.Digitize(n)
		fmt.Printf("Digitize(%d) = %v\n", n, result)
	}
}

package main

import (
	"Golang/CodeWars" // Импортируем пакет CodeWars
	"fmt"
)

func main() {
	// Примеры для проверки функции NthEven
	testCases := []int{1, 3, 100, 1298734}

	for _, n := range testCases {
		result := CodeWars.NthEven(n)
		fmt.Printf("NthEven(%d) = %d\n", n, result)
	}
}

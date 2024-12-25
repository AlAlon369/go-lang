package main

import (
	"Golang/CodeWars" // Импортируем пакет CodeWars
	"fmt"
)

func main() {
	// Примеры для проверки функции AmIWilson
	testCases := []int{5, 13, 563, 7, 1, 0, 10, 6}

	for _, n := range testCases {
		result := CodeWars.AmIWilson(n)
		fmt.Printf("AmIWilson(%d) = %v\n", n, result)
	}
}

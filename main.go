package main

import (
	"Golang/CodeWars"
	"fmt"
	"math"
)

func main() {
	// Тестовые кейсы
	testCases := []struct {
		a, b, c float64
	}{
		{7, 0.40e14, 8}, // Пример из задачи
		{1, 1e9, 1},     // Простой случай
		{2, 1e10, 3},    // Ещё один случай
	}

	// Прогон тестов
	for i, tc := range testCases {
		x2 := CodeWars.Quadratic(tc.a, tc.b, tc.c)
		fmt.Printf("Test Case %d: a=%.2f, b=%.2f, c=%.2f\n", i+1, tc.a, tc.b, tc.c)
		fmt.Printf("  Solution x2: %.12e\n", x2)

		// Проверяем, что |g(x2)| < 1e-12
		gx2 := tc.a*x2*x2 + tc.b*x2 + tc.c
		if math.Abs(gx2) < 1e-12 {
			fmt.Println("  ✅ abs(g(x2)) < 1e-12")
		} else {
			fmt.Printf("  ❌ abs(g(x2)) = %.12e >= 1e-12\n", math.Abs(gx2))
		}
		fmt.Println("-----------------------------")
	}
}

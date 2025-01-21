package main

import (
	"Golang/CodeWars"
	"fmt"
)

func main() {
	// Тестируем функцию OverTheRoad
	fmt.Println(CodeWars.OverTheRoad(1, 3))   // Ожидаем 6
	fmt.Println(CodeWars.OverTheRoad(3, 3))   // Ожидаем 4
	fmt.Println(CodeWars.OverTheRoad(2, 3))   // Ожидаем 5
	fmt.Println(CodeWars.OverTheRoad(3, 5))   // Ожидаем 8
	fmt.Println(CodeWars.OverTheRoad(10, 20)) // Ожидаем 31
}

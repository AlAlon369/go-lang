package main

import (
	"Golang/CodeWars"
	"fmt"
)

func main() {
	// Тестируем функцию Angle
	fmt.Println(CodeWars.Angle(3))  // Ожидаем 180 (треугольник)
	fmt.Println(CodeWars.Angle(4))  // Ожидаем 360 (четырехугольник)
	fmt.Println(CodeWars.Angle(5))  // Ожидаем 540 (пятиугольник)
	fmt.Println(CodeWars.Angle(10)) // Ожидаем 1440 (десятиугольник)
}

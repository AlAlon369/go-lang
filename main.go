package main

import (
	"Golang/CodeWars"
	"fmt"
)

func main() {
	// Создаем функцию addOne из пакета CodeWars
	addOne := CodeWars.Add(1)

	// Тестируем addOne
	fmt.Println(addOne(3))  // Ожидаем 4
	fmt.Println(addOne(10)) // Ожидаем 11

	// Пример с другим числом
	addFive := CodeWars.Add(5)
	fmt.Println(addFive(10)) // Ожидаем 15
	fmt.Println(addFive(20)) // Ожидаем 25
}

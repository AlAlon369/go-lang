package main

import (
	"Golang/CodeWars" // Импортируем пакет CodeWars
	"fmt"
)

func main() {
	arr := []int{1, 2, 3, 4, 5}
	n := 2

	// Вызов функции EachCons из пакета CodeWars
	result := CodeWars.EachCons(arr, n)
	fmt.Println("Результат функции EachCons:", result)
}

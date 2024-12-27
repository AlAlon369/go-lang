package main

import (
	"Golang/CodeWars"
	"fmt"
)

func main() {
	// Пример данных
	chores := []int{1, 5, 2, 8, 4, 9, 6, 4, 2, 2, 2, 9}

	// Вызываем функцию и получаем результат
	result := CodeWars.ChoreAssignment(chores)

	// Вывод результата
	fmt.Println(result) // Ожидаемый результат: [7, 8, 8, 10, 10, 11]
}

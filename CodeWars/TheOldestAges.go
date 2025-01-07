package CodeWars

import "sort"

// TwoOldestAges возвращает два наибольших возраста как массив [2]int.
func TwoOldestAges(ages []int) [2]int {
	if len(ages) < 2 {
		panic("The input array must have at least 2 elements")
	}

	// Сортируем массив по возрастанию
	sort.Ints(ages)

	// Возвращаем два последних элемента как [второй по старшинству, самый старший]
	return [2]int{ages[len(ages)-2], ages[len(ages)-1]}
}

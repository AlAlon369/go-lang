package CodeWars

import "sort"

// TwoOldestAges возвращает два наибольших возраста из массива ages.
func TwoOldestAges(ages []int) []int {
	if len(ages) < 2 {
		return nil // Проверяем, что массив содержит минимум 2 элемента
	}

	// Сортируем массив по возрастанию
	sort.Ints(ages)

	// Возвращаем два последних элемента как [второй по старшинству, самый старший]
	return []int{ages[len(ages)-2], ages[len(ages)-1]}
}

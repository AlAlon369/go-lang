package CodeWars

import "sort"

func Mirror(data []int) []int {
	if len(data) == 0 {
		return []int{}
	}

	// Копируем и сортируем слайс
	sorted := append([]int{}, data...)
	sort.Ints(sorted)

	// Создаем зеркальную версию
	mirrored := append(sorted, sorted[:len(sorted)-1]...)
	reverse(mirrored[len(sorted):]) // Разворачиваем вторую половину

	return mirrored
}

// reverse переворачивает переданный слайс на месте
func reverse(data []int) {
	for i, j := 0, len(data)-1; i < j; i, j = i+1, j-1 {
		data[i], data[j] = data[j], data[i]
	}
}

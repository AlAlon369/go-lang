package CodeWars

import (
	"strconv"
)

func Solve(st string, k int) int {
	n := len(st)
	if k == 0 {
		// Если нет запятых, вся строка — одно число
		num, _ := strconv.Atoi(st)
		return num
	}

	maxValue := 0

	// Генерируем разделения запятыми
	partitions := generatePartitions(n, k)
	for _, partition := range partitions {
		currentMax := 0
		start := 0
		for _, end := range partition {
			num, _ := strconv.Atoi(st[start:end])
			if num > currentMax {
				currentMax = num
			}
			start = end
		}
		// Проверяем последнюю часть строки
		lastNum, _ := strconv.Atoi(st[start:])
		if lastNum > currentMax {
			currentMax = lastNum
		}
		if currentMax > maxValue {
			maxValue = currentMax
		}
	}

	return maxValue
}

// Генерация возможных позиций для вставки запятых
func generatePartitions(n, k int) [][]int {
	var result [][]int
	var dfs func(start int, remaining int, path []int)

	dfs = func(start int, remaining int, path []int) {
		if remaining == 0 {
			result = append(result, append([]int(nil), path...))
			return
		}
		for i := start + 1; i < n; i++ {
			path = append(path, i)
			dfs(i, remaining-1, path)
			path = path[:len(path)-1]
		}
	}

	dfs(0, k, []int{})
	return result
}

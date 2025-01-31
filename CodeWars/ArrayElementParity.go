package CodeWars

func SolveTask(arr []int) int {
	// Создаем мапу для подсчета количества вхождений абсолютных значений чисел
	counts := make(map[int]int)

	// Проходимся по массиву и заполняем мапу
	for _, num := range arr {
		absNum := abs(num)
		counts[absNum]++
	}

	// Теперь ищем число, у которого абсолютное значение встречается только один раз
	for _, num := range arr {
		absNum := abs(num)
		if counts[absNum] == 1 {
			return num
		}
	}

	// Если такого числа нет (что теоретически невозможно по условиям задачи), возвращаем 0
	return 0
}

// Вспомогательная функция для вычисления модуля числа
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

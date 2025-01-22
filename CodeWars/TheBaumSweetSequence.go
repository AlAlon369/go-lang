package CodeWars

// BaumSweet генерирует последовательность Баум-Свита и отправляет значения в канал.
func BaumSweet(ch chan<- int) {
	defer close(ch) // Закрываем канал после завершения генерации.

	for n := 0; ; n++ {
		if isBaumSweet(n) {
			ch <- 1
		} else {
			ch <- 0
		}
	}
}

// isBaumSweet проверяет, соответствует ли число правилу Баум-Свита.
func isBaumSweet(n int) bool {
	if n == 0 {
		return true // Пустая строка соответствует b0 = 1.
	}

	zeroCount := 0
	for n > 0 {
		if n&1 == 0 { // Последний бит равен 0.
			zeroCount++
		} else { // Последний бит равен 1.
			if zeroCount%2 == 1 {
				return false // Найден блок нулей нечетной длины.
			}
			zeroCount = 0
		}
		n >>= 1 // Сдвигаем число вправо на 1 бит.
	}
	return zeroCount%2 == 0 // Проверяем последний блок нулей.
}

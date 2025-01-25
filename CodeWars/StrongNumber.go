package CodeWars

// Вспомогательная функция для вычисления факториала
func factorial(n int) int {
	if n == 0 || n == 1 {
		return 1
	}
	result := 1
	for i := 2; i <= n; i++ {
		result *= i
	}
	return result
}

// Проверка, является ли число "сильным"
func Strong(n int) string {
	original := n
	sum := 0

	// Разбиваем число на цифры и суммируем их факториалы
	for n > 0 {
		digit := n % 10
		sum += factorial(digit)
		n /= 10
	}

	// Сравниваем сумму факториалов с исходным числом
	if sum == original {
		return "STRONG!!!!"
	}
	return "Not Strong !!"
}

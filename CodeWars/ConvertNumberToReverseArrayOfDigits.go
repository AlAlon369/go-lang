package CodeWars

func Digitize(n int) []int {
	var result []int
	for n > 0 {
		result = append(result, n%10) // Добавляем последнюю цифру числа
		n /= 10                       // Убираем последнюю цифру
	}
	if len(result) == 0 { // Обрабатываем случай, когда n == 0
		return []int{0}
	}
	return result
}

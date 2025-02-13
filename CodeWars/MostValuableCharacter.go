package CodeWars

// SolveString находит самый ценный символ строки
func SolveString(s string) rune {
	first := make(map[rune]int) // Индексы первого появления символов
	last := make(map[rune]int)  // Индексы последнего появления символов

	for i, ch := range s {
		if _, exists := first[ch]; !exists {
			first[ch] = i
		}
		last[ch] = i
	}

	maxValue := -1
	var bestRune rune

	for ch, firstIdx := range first {
		value := last[ch] - firstIdx
		if value > maxValue || (value == maxValue && ch < bestRune) {
			maxValue = value
			bestRune = ch
		}
	}

	return bestRune
}

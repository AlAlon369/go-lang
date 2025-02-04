package CodeWars

func Solved(s string) string {
	runes := []rune(s)
	reversed := make([]rune, len(runes))

	// Сначала копируем пробелы на их исходные места
	for i, ch := range runes {
		if ch == ' ' {
			reversed[i] = ' '
		}
	}

	// Заполняем оставшиеся позиции символами из исходной строки, идя с конца
	j := len(runes) - 1
	for i := 0; i < len(runes); i++ {
		if runes[i] != ' ' {
			for reversed[j] == ' ' {
				j--
			}
			reversed[j] = runes[i]
			j--
		}
	}

	return string(reversed)
}

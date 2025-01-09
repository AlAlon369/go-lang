package CodeWars

func GuessHatColor(a, b, c, d string) int {
	// Человек 1 видит шляпы b и c
	if b == c {
		// Если цвета b и c одинаковы, человек 1 угадывает свой цвет
		return 1
	}

	// Если цвета b и c различны, человек 2 угадывает свой цвет
	return 2
}

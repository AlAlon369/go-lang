package CodeWars

// Triangle вычисляет конечный цвет треугольника
func Triangle(row string) rune {
	// Если строка состоит из одного символа, просто возвращаем его
	if len(row) == 1 {
		return rune(row[0])
	}

	// Карта для определения нового цвета
	colorMap := map[string]rune{
		"RR": 'R', "GG": 'G', "BB": 'B',
		"RG": 'B', "GR": 'B',
		"RB": 'G', "BR": 'G',
		"GB": 'R', "BG": 'R',
	}

	// Преобразуем строку в массив рун для удобства работы
	chars := []rune(row)

	// Постепенно уменьшаем строку
	for len(chars) > 1 {
		newRow := make([]rune, len(chars)-1)
		for i := 0; i < len(chars)-1; i++ {
			newRow[i] = colorMap[string(chars[i])+string(chars[i+1])]
		}
		chars = newRow
	}

	return chars[0]
}

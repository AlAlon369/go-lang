package CodeWars

func Solve(slice []string) []int {
	results := make([]int, len(slice))
	for i, word := range slice {
		count := 0
		for j, char := range word {
			position := j + 1                     // Позиция в слове (начинается с 1)
			alphabetPos := int(char|32) - 'a' + 1 // Позиция в алфавите (игнорируем регистр)
			if position == alphabetPos {
				count++
			}
		}
		results[i] = count
	}
	return results
}

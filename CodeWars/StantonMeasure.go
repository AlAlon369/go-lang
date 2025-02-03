package CodeWars

func StantonMeasure(arr []int) int {
	countOne := 0
	for _, num := range arr {
		if num == 1 {
			countOne++
		}
	}

	countN := 0
	for _, num := range arr {
		if num == countOne {
			countN++
		}
	}

	return countN
}

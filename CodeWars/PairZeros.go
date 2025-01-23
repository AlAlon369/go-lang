package CodeWars

func PairZeros(arr []uint64) []uint64 {
	if len(arr) == 0 {
		return []uint64{}
	}

	var result []uint64
	zeroCount := 0

	for _, num := range arr {
		if num == 0 {
			zeroCount++
			if zeroCount%2 != 0 {
				result = append(result, num)
			}
		} else {
			result = append(result, num)
		}
	}

	return result
}

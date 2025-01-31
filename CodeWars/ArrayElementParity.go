package CodeWars

func SolveTask(arr []int) int {
	counts := make(map[int]int)
	signs := make(map[int][2]bool)

	for _, num := range arr {
		absNum := abs(num)
		counts[absNum]++
		if num > 0 {
			signs[absNum] = [2]bool{true, signs[absNum][1]}
		} else {
			signs[absNum] = [2]bool{signs[absNum][0], true}
		}
	}

	for _, num := range arr {
		absNum := abs(num)
		if counts[absNum] == 1 || (counts[absNum] > 1 && (!signs[absNum][0] || !signs[absNum][1])) {
			return num
		}
	}

	return 0
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

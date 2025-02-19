package CodeWars

func CreateBox(width, length int) [][]int {
	box := make([][]int, length)
	for i := range box {
		box[i] = make([]int, width)
	}

	layers := (min(width, length) + 1) / 2

	for layer := 0; layer < layers; layer++ {
		value := layer + 1
		for i := layer; i < length-layer; i++ {
			for j := layer; j < width-layer; j++ {
				box[i][j] = value
			}
		}
	}

	return box
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

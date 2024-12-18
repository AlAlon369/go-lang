package CodeWars

// EachCons создает срезы длины n из массива arr
func EachCons(arr []int, n int) [][]int {
	var result [][]int
	for i := 0; i <= len(arr)-n; i++ {
		result = append(result, arr[i:i+n])
	}
	return result
}

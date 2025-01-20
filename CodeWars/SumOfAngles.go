package CodeWars

// Angle вычисляет сумму внутренних углов n-угольника
func Angle(n int) int {
	if n <= 2 {
		// У n-угольника с n <= 2 внутренние углы отсутствуют
		return 0
	}
	return (n - 2) * 180
}

package CodeWars

// MaxMultiple находит наибольшее число N, которое делится на divisor, меньше либо равно bound
func MaxMultiple(divisor, bound int) int {
	return (bound / divisor) * divisor
}

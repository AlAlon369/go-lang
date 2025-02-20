package CodeWars

func SolveFunc(s int, g int) []int {
	if s%g != 0 {
		return []int{-1, -1}
	}

	// Ищем минимальный x, кратный g, такой что gcd(x/g, (s-x)/g) == 1
	for x := g; x <= s/2; x += g {
		y := s - x
		if gcd(x/g, y/g) == 1 {
			return []int{x, y}
		}
	}

	return []int{-1, -1}
}

// НОД (алгоритм Евклида)
func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

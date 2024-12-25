package CodeWars

import "math/big"

// AmIWilson проверяет, является ли число Wilson-простым
func AmIWilson(n int) bool {
	// Проверяем, является ли число простым
	if !isPrime(n) {
		return false
	}

	// Проверяем условие Wilson-prime
	factorial := big.NewInt(0).MulRange(1, int64(n-1)) // (n-1)!
	nBig := big.NewInt(int64(n))
	leftSide := big.NewInt(0).Add(factorial, big.NewInt(1)) // (n-1)! + 1

	return big.NewInt(0).Mod(leftSide, nBig.Mul(nBig, nBig)).Cmp(big.NewInt(0)) == 0 // Проверяем делимость
}

// isPrime проверяет, является ли число простым
func isPrime(n int) bool {
	if n < 2 {
		return false
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

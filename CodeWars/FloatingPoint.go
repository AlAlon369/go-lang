package CodeWars

import "math"

// Quadratic решает уравнение ax^2 + bx + c = 0
// и возвращает меньший по абсолютной величине корень x2.
func Quadratic(a float64, b float64, c float64) float64 {
	// Вычисление дискриминанта
	discriminant := b*b - 4*a*c

	// Защита от численных ошибок (дискриминант должен быть >= 0)
	if discriminant < 0 {
		return math.NaN() // Возвращаем NaN, если нет реальных корней
	}

	// Корень дискриминанта
	sqrtD := math.Sqrt(discriminant)

	// Из-за больших значений b используем устойчивую формулу:
	// Выбираем корень, который меньше по абсолютной величине
	if b > 0 {
		x2 := (-2 * c) / (b + sqrtD)
		return x2
	}
	x2 := (-b - sqrtD) / (2 * a)
	return x2
}

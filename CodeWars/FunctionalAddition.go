package CodeWars

// Add возвращает функцию, которая добавляет переданное значение n к любому числу.
func Add(n int) func(int) int {
	return func(x int) int {
		return x + n
	}
}

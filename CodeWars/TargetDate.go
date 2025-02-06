package CodeWars

import (
	"time"
)

// DateNbDays вычисляет дату, когда сумма на счете достигнет или превысит a
func DateNbDays(a0 int, a int, p int) string {
	startDate := time.Date(2016, 1, 1, 0, 0, 0, 0, time.UTC)
	dailyRate := float64(p) / 36000.0
	currentAmount := float64(a0)
	days := 0

	for currentAmount < float64(a) {
		currentAmount += currentAmount * dailyRate
		days++
	}

	targetDate := startDate.AddDate(0, 0, days)
	return targetDate.Format("2006-01-02")
}

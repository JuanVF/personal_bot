package services

import "time"

func getMonthStartAndEnd(t time.Time) (startOfMonth, endOfMonth time.Time) {
	// Calculate the start of the current month
	startOfMonth = time.Date(t.Year(), t.Month(), 1, 0, 0, 0, 0, t.Location())

	// Calculate the start of the next month and subtract one day to get the end of the current month
	endOfMonth = time.Date(t.Year(), t.Month()+1, 1, 0, 0, 0, 0, t.Location()).AddDate(0, 0, -1)

	return startOfMonth, endOfMonth
}

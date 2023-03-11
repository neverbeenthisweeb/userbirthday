package common

import "time"

// GetBeginningOfToday gets today time at 00:00
func GetBeginningOfToday(t time.Time) time.Time {
	yy, mm, dd := t.Date()
	return time.Date(yy, mm, dd, 0, 0, 0, 0, t.Location())
}

// GetBeginningOfTomorrow gets tomorrow time at 00:00
func GetBeginningOfTomorrow(t time.Time) time.Time {
	yy, mm, dd := t.Add(24 * time.Hour).Date()
	return time.Date(yy, mm, dd, 0, 0, 0, 0, t.Location())
}

package utils

import "time"

// ToLocalTime handles the timezone conversion for database timestamps
// Database stores as "timestamp without time zone", but GORM reads it
// with current timezone offset, causing 7-hour difference.
// This function strips the timezone and reinterprets as local.
func ToLocalTime(t time.Time) time.Time {
	// Get the wall clock time (ignore timezone)
	year, month, day := t.Date()
	hour, min, sec := t.Clock()
	nsec := t.Nanosecond()

	// Reconstruct as local time
	localTime := time.Date(year, month, day, hour, min, sec, nsec, time.Local)

	// Add 7 hours to compensate for the GORM reading issue
	// GORM reads "18:07:00" from DB as "18:07:00+07:00" which is "11:07:00 UTC"
	// We need to add 7 hours back to get the correct local time
	return localTime.Add(7 * time.Hour)
}

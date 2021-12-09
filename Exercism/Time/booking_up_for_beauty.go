package booking

import (
	"fmt"
	"time"
)

// Schedule returns a time.Time from a string containing a date
func Schedule(date string) time.Time {
	const schedLayout = "1/2/2006 15:04:05" // specific date to use when parsing date (no repeating numbers for compiler identification)
	t, _ := time.Parse(schedLayout, date)
	return t
	panic("Please implement the Schedule function")
}

// HasPassed returns whether a date has passed
func HasPassed(date string) bool {
	const checkDateLayout = "January 2, 2006 15:04:05" //to match string to be inputted.
	t, _ := time.Parse(checkDateLayout, date)
	return t.Before(time.Now())
	panic("Please implement the HasPassed function")
}

// IsAfternoonAppointment returns whether a time is in the afternoon
func IsAfternoonAppointment(date string) bool {
	t, _ := time.Parse("Monday, January 2, 2006 15:04:05", date)
	if t.Hour() >= 12 && t.Hour() < 18 {
		return true
	} else {
		return false
	}
	panic("Please implement the IsAfternoonAppointment function")
}

// Description returns a formatted string of the appointment time
func Description(date string) string {
	t, _ := time.Parse("1/2/2006 15:04:05", date)
	return fmt.Sprintf("You have an appointment on %s, %s, at %d:%d.", t.Weekday(), t.Format("January 2, 2006"), t.Hour(), t.Minute())
	panic("Please implement the Description function")
}

// AnniversaryDate returns a Time with this year's anniversary
func AnniversaryDate() time.Time {
	anniv := "2021-09-15"
	t, _ := time.Parse("2006-01-02", anniv)
	return t
	panic("Please implement the AnniversaryDate function")
}

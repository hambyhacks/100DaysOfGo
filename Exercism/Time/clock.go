package clock

import (
	"fmt"
	"time"
)

// Define the Clock type here.
type Clock struct {
	hours   int
	minutes int
}

func New(h, m int) Clock {
	clock := time.Date(2006, time.January, 2, h, m, 0, 0, time.UTC)
	return Clock{hours: clock.Hour(), minutes: clock.Minute()}
}

func (c Clock) Add(m int) Clock {
	baseClock := time.Date(2006, time.January, 2, c.hours, c.minutes, 0, 0, time.UTC)
	newClock := baseClock.Add(time.Duration(m) * time.Minute)

	hours := newClock.Hour()
	mins := newClock.Minute()
	return New(hours, mins)
}

func (c Clock) Subtract(m int) Clock {
	baseClock := time.Date(2006, time.January, 2, c.hours, c.minutes, 0, 0, time.UTC)
	newClock := baseClock.Add(-time.Duration(m) * time.Minute)

	hours := newClock.Hour()
	mins := newClock.Minute()
	return New(hours, mins)
}

func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.hours, c.minutes)
}

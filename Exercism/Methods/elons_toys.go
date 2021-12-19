package elon

import (
	"strconv"
)

// gets the *memory address* of struct `Car`` and places it to variable `c`.
func (c *Car) Drive() Car {
	if c.battery < c.batteryDrain {
		return *c
	}
	c.battery -= c.batteryDrain
	c.speed -= c.distance
	c.distance = c.speed
	return *c // fetch the new values from modified `Car` struct.
}

func (c Car) DisplayDistance() string {
	return "Driven " + strconv.Itoa(c.distance) + " meters"
}

func (c Car) DisplayBattery() string {
	return "Battery at " + strconv.Itoa(c.battery) + "%"
}

func (c Car) CanFinish(trackDistance int) bool {
	return !((c.battery/c.batteryDrain)*c.speed < trackDistance)
}

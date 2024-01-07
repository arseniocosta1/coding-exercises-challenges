package elon

import (
	"strconv"
)

func (c *Car) Drive() {
	if c.battery >= c.batteryDrain {
		c.distance += c.speed
		c.battery -= c.batteryDrain
	}
}

func (c Car) DisplayDistance() string {
	return "Driven " + strconv.Itoa(c.distance) + " meters"
}

func (c Car) DisplayBattery() string {
	return "Battery at " + strconv.Itoa(c.battery) + "%"
}

func (c Car) CanFinish(trackDistance int) bool {
	return trackDistance <= c.battery/c.batteryDrain*c.speed
}

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.

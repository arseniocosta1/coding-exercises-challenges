package meteorology

import (
	"strconv"
)

type TemperatureUnit int

const (
	Celsius    TemperatureUnit = 0
	Fahrenheit TemperatureUnit = 1
)

func (t TemperatureUnit) String() string {
	switch t {
	case Celsius:
		return "°C"
	case Fahrenheit:
		return "°F"
	default:
		panic("invalid temperature unit")
	}
}

type Temperature struct {
	degree int
	unit   TemperatureUnit
}

func (t Temperature) String() string {
	return strconv.Itoa(t.degree) + " " + t.unit.String()
}

type SpeedUnit int

const (
	KmPerHour    SpeedUnit = 0
	MilesPerHour SpeedUnit = 1
)

// Add a String method to SpeedUnit
func (s SpeedUnit) String() string {
	switch s {
	case KmPerHour:
		return "km/h"
	case MilesPerHour:
		return "mph"
	default:
		panic("invalid speed unit")
	}
}

type Speed struct {
	magnitude int
	unit      SpeedUnit
}

// Add a String method to Speed
func (s Speed) String() string {
	return strconv.Itoa(s.magnitude) + " " + s.unit.String()

}

type MeteorologyData struct {
	location      string
	temperature   Temperature
	windDirection string
	windSpeed     Speed
	humidity      int
}

func (m MeteorologyData) String() string {
	return m.location + ": " + m.temperature.String() +
		", Wind " + m.windDirection + " at " + m.windSpeed.String() +
		", " + strconv.Itoa(m.humidity) + "% Humidity"
}

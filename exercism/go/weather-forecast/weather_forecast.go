// Package weather provides the current weather condition of various cities in Goblinocus.
package weather

// CurrentCondition holds the current weather condition of the last forecast.
var CurrentCondition string

// CurrentLocation holds the current weather location of the last forecast.
var CurrentLocation string

// Forecast returns a string with the current weather condition of the given city.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}

// Package weather from exercism.
package weather

// CurrentCondition in the location informed.
var CurrentCondition string

// CurrentLocation for check.
var CurrentLocation string

// Forecast funtion returns  the Weather condition in the city.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}

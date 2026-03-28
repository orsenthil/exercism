// Package weather provides details about the the weather conditions.
package weather

// CurrentCondition used for weather forecast calculation.
var CurrentCondition string

// CurrentLocation used for weather forecast calculation.
var CurrentLocation string

// Forecast returns the weather forecast given the city and condition.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
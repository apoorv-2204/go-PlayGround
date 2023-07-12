package exercism

// Package weather  forecast provides tools to forecast the current weather condition of various cities .

// CurrentCondition represents a certain weather.
var CurrentCondition string

// CurrentLocation represents a city.
var CurrentLocation string

// Forecast returns an string forcasting the ci.
func Forecast(city string, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}

// About comments

// Package weather provides tools to get a weather forecast.
package weather


var (
    // CurrentCondition is the latest weather condition set by Forecast.
	CurrentCondition string
    // CurrentLocation is the latest city set by Forecast.
	CurrentLocation  string
)

// Forecast sets the current location and condition, and returns a formatted summary string.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}

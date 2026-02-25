// Package weather provides tools for weather forecasting.
package weather

var (
    // CurrentCondition represents the state of the weather.
	CurrentCondition string
    // CurrentLocation represents the city or region.
	CurrentLocation  string
)

// Forecast returns a weather forecast for a location.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}

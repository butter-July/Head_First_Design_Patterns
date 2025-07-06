package main

func main() {
	weatherData := NewWeatherData()
	currentCoditionsDisplay := NewCurrentConditionsDisplay()
	staticDisplay := NewStatisticsDisplay()
	foreCastDisplay := NewForeCastDisplay()
	weatherData.RegisterObserver(currentCoditionsDisplay)
	weatherData.RegisterObserver(staticDisplay)
	weatherData.RegisterObserver(foreCastDisplay)
	weatherData.SetMeasurements(80, 65, 30.4)
	weatherData.SetMeasurements(82, 70, 29.2)
	weatherData.SetMeasurements(78, 90, 29.2)
}

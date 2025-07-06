package main

import "fmt"

type DisplayElement interface {
	display()
}
type CurrentConditionsDisplay struct {
	temperature, humidity float64
}

func NewCurrentConditionsDisplay() *CurrentConditionsDisplay {
	return &CurrentConditionsDisplay{}
}

func (ccd *CurrentConditionsDisplay) update(temp, humidity, pressure float64) {
	ccd.temperature = temp
	ccd.humidity = humidity
	ccd.Display()
}
func (ccd *CurrentConditionsDisplay) Display() {
	fmt.Printf("Current conditions: %.2f°F temperature and %.2f%% humidity\n", ccd.temperature, ccd.humidity)
}

type StatisticsDisplay struct {
	maxTemp, minTemp, avgTemp float64
	temperatureCount          int
	WeatherData
}

func NewStatisticsDisplay() *StatisticsDisplay {
	return &StatisticsDisplay{}
}
func (sd *StatisticsDisplay) update(temp, humidity, pressure float64) {
	if sd.temperatureCount == 0 {
		sd.minTemp = temp
		sd.maxTemp = temp
		sd.avgTemp = temp
	} else {
		if temp < sd.minTemp {
			sd.minTemp = temp
		}
		if temp > sd.maxTemp {
			sd.maxTemp = temp
		}
		sd.avgTemp = (sd.avgTemp*float64(sd.temperatureCount) + temp) / float64(sd.temperatureCount+1)
	}

	sd.temperatureCount++
	sd.Display()
}

func (sd *StatisticsDisplay) Display() {
	fmt.Printf("Temperature statistics: Min %.2f°F, Max %.2f°F, Average %.2f°F\n", sd.minTemp, sd.maxTemp, sd.avgTemp)
}
func NewForeCastDisplay() *ForeCastDisplay {
	return &ForeCastDisplay{}
}

type ForeCastDisplay struct {
	currentPressure, lastPressure float64
	WeatherData
}

func (fcd *ForeCastDisplay) update(temp, humidity, pressure float64) {
	fcd.lastPressure = fcd.currentPressure
	fcd.currentPressure = pressure
	fcd.Display()
}
func (fcd *ForeCastDisplay) Display() {
	fmt.Println("Forecast:")
	if fcd.currentPressure > fcd.lastPressure {
		fmt.Println("Improving weather on the way!")
	} else if fcd.lastPressure == fcd.currentPressure {
		fmt.Println("More of the same")
	} else {
		fmt.Println("Watch out for cooler, rainy weather")
	}
}

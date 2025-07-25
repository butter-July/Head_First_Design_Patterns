package main

type WeatherData struct {
	observers                       []Observer
	temperature, humidity, pressure float64
}

func NewWeatherData() *WeatherData {
	return &WeatherData{}
}

func (wd *WeatherData) RegisterObserver(observer Observer) {
	wd.observers = append(wd.observers, observer)
}

func (wd *WeatherData) RemoveObserver(observer Observer) {
	index := -1
	for i, obs := range wd.observers {
		if obs == observer {
			index = i
			break
		}
	}
	if index >= 0 {
		wd.observers = append(wd.observers[:index], wd.observers[index+1:]...)
	}
}

func (wd *WeatherData) NotifyObservers() {
	for _, observer := range wd.observers {
		observer.update(wd.temperature, wd.humidity, wd.pressure)
	}
}

func (wd *WeatherData) measurementsChanged() {
	wd.NotifyObservers()
}

func (wd *WeatherData) SetMeasurements(temperature, humidity, pressure float64) {
	wd.temperature = temperature
	wd.humidity = humidity
	wd.pressure = pressure
	wd.measurementsChanged()
}

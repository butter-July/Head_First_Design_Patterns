package main

import "fmt"

type TheaterLights struct {
}

func (light *TheaterLights) dim() {
	fmt.Println("Theater Ceiling lights dimming to 10%")
}
func (light *TheaterLights) on() {
	fmt.Println("Theater Ceiling lights on")
}

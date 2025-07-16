package main

import "fmt"

type Amplifier struct {
	amp string
}

func (a *Amplifier) setStreamingPlayer() {
	fmt.Println("Amplifier setting Steaming player to Streaming Player")
}
func (a *Amplifier) setSurroundingSound() {
	fmt.Println("Amplifier surround sound on (5 speakers ,1 subwoofer)")
}
func (a *Amplifier) setVolume(volume int) {
	switch volume {
	case 1:
		fmt.Println(" setting volume to 1")
	case 2:
		fmt.Println("setting volume to 2")
	case 3:
		fmt.Println("setting volume to 3")
	case 4:
		fmt.Println("setting volume to 4")
	case 5:
		fmt.Println("setting volume to 5")
	default:
		fmt.Println("wrong number!!")
	}
}
func (a *Amplifier) on() {
	fmt.Println("Amplifier on ")
}
func (a *Amplifier) off() {
	fmt.Println("Amplifier off!")
}

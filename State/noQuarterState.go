package main

import (
	"fmt"
)

type NoQuarter struct {
	machine *GumballMachine
}

func (noQuarter *NoQuarter) insertQuarter() {
	fmt.Println("you can insert quarter")
	noQuarter.machine.setState(noQuarter.machine.HasQuarterState)
}

func (noQuarter *NoQuarter) ejectQuarter() {
	fmt.Println("you haven't insert a quarter!!")
}

func (noQuarter *NoQuarter) turnCrank() {
	fmt.Println("you can turn but there's no quarter")
}
func (noQuarter *NoQuarter) dispense() {
	fmt.Println("you need to pay first")
}

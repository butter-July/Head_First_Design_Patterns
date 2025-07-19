package main

import "fmt"

type GumballMonitor struct {
	machine *GumballMachine
}

func (monitor *GumballMonitor) report() {
	fmt.Printf("Gumball Machine :%s", monitor.machine.getLocation())
	fmt.Printf("Gumball inventory:%d gumballs", monitor.machine.getcout())
}

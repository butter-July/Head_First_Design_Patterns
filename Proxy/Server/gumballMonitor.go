package main

import "fmt"

type GumballMonitor struct {
	machine *GumballMachine
}

func (monitor *GumballMonitor) report() {
	fmt.Println(monitor.machine.getLocation())
	fmt.Printf("Gumball inventory:%d gumballs", monitor.machine.getcout())
}

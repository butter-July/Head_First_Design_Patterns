package main

import (
	"fmt"
)

type SoldOut struct {
	machine *GumballMachine
}

func (SoldOut *SoldOut) insertQuarter() {
	fmt.Println("You can't insert quarter,the machine is  soldOut")
}
func (SoldOut *SoldOut) ejectQuarter() {
	fmt.Println("You can't eject,you haven't insert a quarter yet")
}
func (SoldOut *SoldOut) turnCrank() {
	fmt.Println("You can turn ,but there is  no candy")
}
func (SoldOut *SoldOut) dispense() {
	fmt.Println("No ...")
}

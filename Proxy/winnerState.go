package main

import "fmt"

type WinnerState struct {
	machine *GumballMachine
}

func (winner *WinnerState) insertQuarter() {
	fmt.Println("you can't insert quarter,the machine is sold out")
}

func (winner *WinnerState) ejectQuarter() {
	fmt.Println("Sorry ,you already turned the crank")
}

func (winner *WinnerState) turnCrank() {
	fmt.Println("Turning twice doesn't get you another candy")
}

func (winner *WinnerState) dispense() {
	winner.machine.release()
	if winner.machine.count == 0 {
		winner.machine.setState(&winner.machine.SoldOutState)
	} else {
		fmt.Println("You're A Winner! You got two gumball for you quarter!")
		if winner.machine.getcout() > 0 {
			winner.machine.setState(&winner.machine.NoQuarterState)
		} else {
			winner.machine.setState(&winner.machine.SoldOutState)
		}
	}

}

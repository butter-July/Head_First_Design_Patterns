package main

import "fmt"

type HasQuarterState struct {
	machine *GumballMachine
}

func (hasQuarter *HasQuarterState) insertQuarter() {
	fmt.Println("Can't insert again")
}

func (hasQuarter *HasQuarterState) ejectQuarter() {
	fmt.Println("Quarter returned")
	hasQuarter.machine.setState(&hasQuarter.machine.NoQuarterState)
}

func (hasQuarter *HasQuarterState) turnCrank() {
	fmt.Println("you can turn,candy is coming....")
	hasQuarter.machine.SoldState.dispense() ////////////important!!!!!!!!这里要换到SoldState
	hasQuarter.machine.setState(&hasQuarter.machine.NoQuarterState)

	//投币的情况下转动手柄,然后把状态设置为没有币了,然后进入出糖果的行为

}

func (hasQuarter *HasQuarterState) dispense() {
	fmt.Println("you need to turn the crank!")
}

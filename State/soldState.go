package main

import "fmt"

type Sold struct {
	machine *GumballMachine
}

func (sold *Sold) insertQuarter() {
	fmt.Println("you can't insert quarter,the machine is sold out")
}

func (sold *Sold) ejectQuarter() {
	fmt.Println("Sorry ,you already turned the crank")
}

func (sold *Sold) turnCrank() {
	fmt.Println("Turning twice doesn't get you another candy")
}

func (sold *Sold) dispense() {
	/*if sold.machine.count > 0 {
		fmt.Println("A candy comes rolling out the slot")
		//为什么这样不行		sold.machine.count--
		if sold.machine.count == 0 {
			sold.machine.setState(sold.machine.SoldOutState)
		} else {
			sold.machine.setState(sold.machine.NoQuarterState)
		}
	} else {
		fmt.Println("soldOut")
		sold.machine.setState(sold.machine.SoldOutState)
	}
	*/ // 出糖果
	if sold.machine.count > 0 {
		sold.machine.count = sold.machine.count - 1
		sold.machine.setState(sold.machine.NoQuarterState) // 还有糖果，回到无币状态
	} else {
		fmt.Println("Oops, out of gumballs!")
		sold.machine.setState(sold.machine.SoldOutState) // 无糖果，切换到售罄状态
	}

}

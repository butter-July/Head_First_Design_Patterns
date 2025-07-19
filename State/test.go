package main

import "fmt"

func main() {

	machine := NewGumballMachine(1)
	c := machine.getcout()
	fmt.Println(c)

	machine.insertQuarter()
	machine.turnCrank()
	machine.insertQuarter()
	machine.turnCrank()
	cc := machine.getcout()

	fmt.Println(cc)

}

package main

import "fmt"

type Coffee struct {
	CaffeineBeverage
}

func (c *Coffee) brew() {

	fmt.Println("Dripping Coffee through filter")

}
func (c *Coffee) addCondiments() {

	fmt.Println("concrete things")
	var thing string
	fmt.Scanln(&thing)
	fmt.Printf("Adding %s", thing)

}
func (c *Coffee) hook() bool {

	fmt.Println("Would you like something else?(y/n)")
	var answer string
	fmt.Scanln(&answer) //这里输入一次y/s
	return answer == "y"

}

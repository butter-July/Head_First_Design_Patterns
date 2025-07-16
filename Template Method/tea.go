package main

import "fmt"

type Tea struct {
	CaffeineBeverage
}

func (tea *Tea) hook() bool {
	fmt.Println("Would you like something else?(y/n)")
	var answer string
	fmt.Scanln(&answer)
	return answer == "y"
}
func (tea *Tea) brew() {
	fmt.Println("Dripping Coffee through filter")
}
func (tea *Tea) addCondiments() {
	fmt.Println("concrete things")
	var thing string
	fmt.Scanln(&thing)
	fmt.Printf("Adding %s", thing)

}

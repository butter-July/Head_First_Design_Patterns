package main

import "fmt"

type PopcornPopper struct {
}

func (popper *PopcornPopper) on() {
	fmt.Println("Popcorn Popper on")
}
func (popper *PopcornPopper) pop() {
	fmt.Println("Popcorn popper popping popcorn!")
}
func (popper *PopcornPopper) off() {
	fmt.Println("popper off")
}

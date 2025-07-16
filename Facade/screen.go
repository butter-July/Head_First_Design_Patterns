package main

import "fmt"

type Screen struct {
}

func (screen *Screen) down() {
	fmt.Println("Theater Screen going down")
}
func (screen *Screen) up() {
	fmt.Println("Theater Screen going up")
}

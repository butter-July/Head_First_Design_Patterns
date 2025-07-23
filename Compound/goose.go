package main

import "fmt"

type Goose struct {
}

func (g *Goose) honk() {
	fmt.Println("Honk")
}

type GooseAdapter struct {
	g *Goose
}

func (g *GooseAdapter) quack() {
	g.g.honk()
}

package main

import "fmt"

type MarllardDuck struct {
}

func (m *MarllardDuck) quack() {
	fmt.Println("Quack")

}

type RedheadDuck struct {
}

func (r *RedheadDuck) quack() {
	fmt.Println("Quack")
}

type DuckCall struct {
}

func (d *DuckCall) quack() {
	fmt.Println("Kwak")
}

type RubberDuck struct {
}

func (rr *RubberDuck) quack() {
	fmt.Println("Squeak")
}

package main

import "fmt"

type Quack struct {
}

func (q *Quack) Quack() {
	fmt.Println("Quack!")
}

type MuteQuack struct {
}

func (m *MuteQuack) Quack() {
	fmt.Println("<<Silence>>")
}

type Squeak struct {
}

func (s *Squeak) Quack() {
	fmt.Println("Squeak")
}

func (d *Duck) SetQuackBehaviour(qb QuackBehavior) {
	d.quackBehavior = qb
}

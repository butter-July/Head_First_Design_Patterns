package main

import "fmt"

type Duck interface {
	quack()
	fly()
}
type MallardDuck struct {
}

func (m *MallardDuck) quack() {
	fmt.Println("Quack!")
}
func (m *MallardDuck) fly() {
	fmt.Println("I'm flying!")
}

type Turkey struct {
}

func (t *Turkey) fly() {
	fmt.Println("I'flying a short distance!")
}
func (t *Turkey) gobble() {
	fmt.Println("Gobble gobble")
}

type TurkeyAdapter struct {
	TT *Turkey
}

func (t *TurkeyAdapter) quack() {
	t.TT.gobble()
}
func (t *TurkeyAdapter) fly() {
	t.TT.fly()
}
func main() {
	t := &Turkey{}
	adapter := &TurkeyAdapter{t}
	adapter.quack()
	adapter.fly()
}

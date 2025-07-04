package main

import "fmt"

type FlyBehavior interface {
	Fly()
}

type FlyWithWings struct {
}

func (f *FlyWithWings) Fly() {
	fmt.Println("I.m flying!!")
}

type FlyNoWay struct {
}

func (f *FlyNoWay) Fly() {
	fmt.Println("I.can't fly..")
}

type FlyRocketPowered struct {
}

func (d *Duck) SetFlyBehaviour(fb FlyBehavior) {
	d.flyBehavior = fb
}
func (f FlyRocketPowered) Fly() {
	fmt.Println("I'm flying with rocket!!")
}

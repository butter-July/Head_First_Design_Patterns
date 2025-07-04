package main

import "fmt"

type Duck struct {
	flyBehavior   FlyBehavior
	quackBehavior QuackBehavior
}

func (d *Duck) performanceFly() {
	d.flyBehavior.Fly()
}
func (d *Duck) performanceQuack() {
	d.quackBehavior.Quack()
}
func (d *Duck) swim() {
	fmt.Println("All ducks float,even decoys!")
}

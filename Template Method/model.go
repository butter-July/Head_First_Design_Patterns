package main

import "fmt"

type PeraRecipe interface {
	boilWater()
	pourInCup()
	brew()
	addCondiments()//这里不需要进来东西,他不判断
	hook() bool
}
type CaffeineBeverage struct {
}

func (beverage *CaffeineBeverage) boilWater() {
	fmt.Println("Boiling water")
}

func (Beverage *CaffeineBeverage) pourInCup() {
	fmt.Println("Pouring into cup")
}

func PrepareRecipe(c PeraRecipe) {
	c.boilWater()
	c.brew()
	c.pourInCup()
	if c.hook() {
		c.addCondiments()
	}
}

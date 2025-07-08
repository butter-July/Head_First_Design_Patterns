package main

import "fmt"

func main() {
	nyPizzaStore := newNYPizzaStore()
	pizza := nyPizzaStore.orderPizza("cheese")
	fmt.Printf("the pizza is:", pizza.getName())
}

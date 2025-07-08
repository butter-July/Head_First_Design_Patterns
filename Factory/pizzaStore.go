package main

import "fmt"

type PizzaStore interface {
	orderPizza(pizzaType string) PPizza
	createPizza(pizzaType string) (PPizza, error)
}

type PPizzaStore struct {
	createPizza func(pizzaType string) (PPizza, error)
}

func (a *PPizzaStore) orderPizza(pizzaType string) PPizza {
	if pizza, err := a.createPizza(pizzaType); err != nil {
		fmt.Println(err.Error())
		return nil
	} else {
		pizza.prepare()
		pizza.bake()
		pizza.cut()
		pizza.box()
		return pizza
	}
}

type nyPizzaStore struct {
	*PPizzaStore
}

func newNYPizzaStore() PizzaStore {
	basePizzaStore := &PPizzaStore{}

	nyPizzaStore := &nyPizzaStore{basePizzaStore}

	nyPizzaStore.PPizzaStore.createPizza = nyPizzaStore.createPizza

	return nyPizzaStore
}
func (n *nyPizzaStore) createPizza(PizzaType string) (PPizza, error) {
	switch PizzaType {
	case "cheese":
		return newNYStyleCheesePizza(), nil
	case "clam":
		return newNYStyleClamPizza(), nil
	case "pepperoni":
		return newNYStylePepperoniPizza(), nil
	case "Veggie":
		return newNYStyleVeggiePizza(), nil

	}
	return nil, fmt.Errorf("Invalid pizza type")
}

package main

type PPizza interface {
	prepare()
	bake()
	cut()
	box()
	getName() string
}

type Pizza struct {
	name     string
	dough    string
	sauce    string
	toppings []string
}

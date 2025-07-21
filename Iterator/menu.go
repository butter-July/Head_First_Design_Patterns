package main

type MenuItem struct {
	name        string
	description string
	vegetarian  bool
	price       float64
}
type Menu interface {
	createIter() Iterator
}

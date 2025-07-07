package main

type Beverage interface {
	getDescription() string
	Cost() float64
}

type Espresso struct {
}

func NewEspresso() Beverage {
	return &Espresso{}
}
func (e *Espresso) Cost() float64 {
	return 1.99
}
func (e *Espresso) getDescription() string {
	return "Espresso"
}

type DarkRoast struct {
}

func NewDarkRoast() Beverage {
	return &DarkRoast{}
}
func (d *DarkRoast) getDescription() string {
	return "DarkRost"
}
func (d *DarkRoast) Cost() float64 {
	return 0.99
}

type HouseBlend struct {
}

func NewHouseBlend() Beverage {
	return &HouseBlend{}
}
func (h *HouseBlend) getDescription() string {
	return "HouseBlend"
}
func (h *HouseBlend) Cost() float64 {
	return 0.89
}

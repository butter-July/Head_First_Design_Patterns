package main

type CondimentDecorator interface {
	Beverage
	addDescription() string
}
type Mocha struct {
	beverage Beverage
}

func NewMocha(beverage Beverage) Beverage {
	return &Mocha{beverage: beverage}
}
func (m *Mocha) Cost() float64 {
	return m.beverage.Cost() + 0.20
}
func (m *Mocha) addDescription() string {
	return ",+Mocha"
}
func (m *Mocha) getDescription() string {
	return m.beverage.getDescription() + m.addDescription()
}

type Soy struct {
	beverage Beverage
}

func NewSoy(beverage Beverage) Beverage {
	return &Soy{beverage: beverage}
}
func (m *Soy) Cost() float64 {
	return m.beverage.Cost() + 0.15
}
func (m *Soy) addDescription() string {
	return ",+Soy"
}
func (m *Soy) getDescription() string {
	return m.beverage.getDescription() + m.addDescription()
}

type Whip struct {
	beverage Beverage
}

func NewWhip(beverage Beverage) Beverage {
	return &Whip{beverage: beverage}
}
func (m *Whip) Cost() float64 {
	return m.beverage.Cost() + 0.10
}
func (m *Whip) addDescription() string {
	return ",+Whip"
}
func (m *Whip) getDescription() string {
	return m.beverage.getDescription() + m.addDescription()
}

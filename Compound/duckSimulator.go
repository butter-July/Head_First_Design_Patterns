package main

type DuckSimulator struct {
}

func NewDuckSimulator() *DuckSimulator {
	return &DuckSimulator{}
}
func (ds *DuckSimulator) Simulator(duck Quackable) {
	duck.quack()
}
func (si *DuckSimulator) simulate() {
	mallardDuck := &MarllardDuck{}
	redDuck := &RedheadDuck{}
	duckcall := &DuckCall{}
	rubberDuck := &RubberDuck{}
	goose := &Goose{}
	gooseAdapter := &GooseAdapter{g: goose}

	si.Simulator(mallardDuck)
	si.Simulator(redDuck)
	si.Simulator(duckcall)
	si.Simulator(rubberDuck)
	si.Simulator(gooseAdapter)
}

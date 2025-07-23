package main

import "fmt"

func main() {
	simulator := NewDuckSimulator()

	mallardDuck := &MarllardDuck{}
	redheadDuck := &RedheadDuck{}
	duckCall := &DuckCall{}
	rubberDuck := &RubberDuck{}
	goose := &Goose{}
	gooseAdapter := &GooseAdapter{g: goose}
	decoratedMallard := &QuackCounter{quackable: mallardDuck} //   为什么不能像之前的一样层层的
	decoratedRedhead := &QuackCounter{quackable: redheadDuck}
	decoratedDuckCall := &QuackCounter{quackable: duckCall}
	decoratedRubber := &QuackCounter{quackable: rubberDuck}
	decoratedGoose := &QuackCounter{quackable: gooseAdapter}
	simulator.Simulator(decoratedMallard)
	simulator.Simulator(decoratedMallard)
	simulator.Simulator(decoratedRedhead)
	simulator.Simulator(decoratedDuckCall)
	simulator.Simulator(decoratedRubber)
	simulator.Simulator(decoratedGoose)
	fmt.Println("MallardDuck quack count:", decoratedMallard.GetQuackCounter())
	fmt.Println("RedheadDuck quack count:", decoratedRedhead.GetQuackCounter())
	fmt.Println("Total quacks:",
		decoratedMallard.GetQuackCounter()+
			decoratedRedhead.GetQuackCounter()+
			decoratedDuckCall.GetQuackCounter()+
			decoratedRubber.GetQuackCounter()+
			decoratedGoose.GetQuackCounter())
}

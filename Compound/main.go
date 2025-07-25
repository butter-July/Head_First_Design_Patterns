package main

import "fmt"

func main() {
	si := DuckSimulator{}
	duckFactory := DuckFactory{}
	mallard := duckFactory.createMallardDuck()
	duckCall := duckFactory.createDuckCall()
	redheadDuck := duckFactory.createRedheadDuck()
	rubber := duckFactory.createRubberDuck()
	goose := Goose{}
	gooseAdapter := GooseAdapter{g: &goose}
	si.Simulator(&gooseAdapter)
	si.Simulator(mallard)
	si.Simulator(duckCall)
	si.Simulator(redheadDuck)
	si.Simulator(rubber)
	CounterFactory := CountingDuckFactory{factory: &duckFactory}
	CounterFactory.createMallardDuck()
	CounterFactory.createDuckCall()
	CounterFactory.createRedheadDuck()
	CounterFactory.createRubberDuck() //让里面的number++,但是他没getCount方法,所以要再写一个把之前的删掉
	i := CounterFactory.Counter()
	fmt.Println(i) //4 goose不统计
	/*ducks := []Quackable{
			countingFactory.createMallardDuck(),
			countingFactory.createRedheadDuck(),
			countingFactory.createDuckCall(),
			countingFactory.createRubberDuck(),
		}
	}
	totalQuacks := getTotalQuacks(ducks)
		fmt.Printf("Total quacks: %d\n", totalQuacks)
	}为什么不行捏
	*/
}

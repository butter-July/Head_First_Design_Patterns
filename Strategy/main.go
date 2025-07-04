package main

type QuackBehavior interface {
	Quack()
}

type MallradDuck struct {
	Duck
}
type ModelDuck struct {
	Duck
}

func main() {
	m := MallradDuck{Duck{flyBehavior: &FlyWithWings{}, quackBehavior: &Quack{}}}
	m.performanceFly()
	m.performanceQuack()
	mm := ModelDuck{Duck{flyBehavior: &FlyNoWay{}}}
	mm.performanceFly()
	mm2 := ModelDuck{Duck{flyBehavior: &FlyRocketPowered{}}}
	mm2.performanceFly()
}

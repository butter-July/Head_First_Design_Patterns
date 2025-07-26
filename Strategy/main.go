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
//定义一系列算法,并将每种算法分别放入独立的类中,使得算法的对象能够相互替换.

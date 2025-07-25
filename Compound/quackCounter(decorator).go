package main

type QuackCounter struct {
	quackable     Quackable
	NumberOfQuack int
}

func (qc *QuackCounter) quack() {
	if qc.quackable != nil {
		qc.quackable.quack()
		qc.NumberOfQuack++
	}
}
func (q *QuackCounter) getQuacks() int {
	return q.NumberOfQuack
}

// 变成工厂模式
type CountingDuckFactory struct {
	factory *DuckFactory
}

func (c *CountingDuckFactory) createMallardDuck() Quackable {
	duck := c.factory.createMallardDuck()
	return &QuackCounter{quackable: duck}
}
func (c *CountingDuckFactory) createRedheadDuck() Quackable {
	duck := c.factory.createRedheadDuck()
	return &QuackCounter{quackable: duck}
}
func (c *CountingDuckFactory) createDuckCall() Quackable {
	duck := c.factory.createDuckCall()
	return &QuackCounter{quackable: duck}
}
func (c *CountingDuckFactory) createRubberDuck() Quackable {
	duck := c.factory.createRubberDuck()
	return &QuackCounter{quackable: duck}
}

/*
	func getTotalQuacks(ducks []Quackable) int {
		total := 0
		for _, duck := range ducks {
			if counter, ok := duck.(*QuackCounter); ok {
				total += counter.getQuacks()
			}
		}
		return total
	}
*/ //不知道是不是但不对为什么
func (c *CountingDuckFactory) Counter() int {
	duck1 := c.factory.createMallardDuck()
	quackCounter := QuackCounter{quackable: duck1}
	quackCounter.NumberOfQuack++
	duck2 := c.factory.createRedheadDuck()
	quackCounter2 := QuackCounter{quackable: duck2}
	quackCounter2.NumberOfQuack++
	duck3 := c.factory.createDuckCall()
	quackCounter3 := QuackCounter{quackable: duck3}
	quackCounter3.NumberOfQuack++
	duck4 := c.factory.createRubberDuck()
	quackCounter4 := QuackCounter{quackable: duck4}
	quackCounter4.NumberOfQuack++
	return quackCounter.NumberOfQuack + quackCounter2.NumberOfQuack + quackCounter3.NumberOfQuack + quackCounter4.NumberOfQuack
}

//这一堆好像不是工厂模式,但是能用...

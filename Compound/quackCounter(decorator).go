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
func (quack *QuackCounter) GetQuackCounter() int {
	return quack.NumberOfQuack
}

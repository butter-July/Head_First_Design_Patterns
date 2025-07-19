package main

import "fmt"

type GumballMachine struct {
	HasQuarterState State
	NoQuarterState  State
	SoldState       State
	SoldOutState    State
	CurrentState    State //记录当前状态
	count           int
	WinnerState     State
	Location        string
}

// 他也要实现接口,才能换换换,在里面调用叭
// GumballMachine 要不要初始化,用的时候自己初始化叭,好像需要
func NewGumballMachine(count int) *GumballMachine {
	machine := &GumballMachine{count: count}
	machine.SoldOutState = &SoldOut{machine: machine}
	machine.SoldState = &Sold{machine: machine}
	machine.NoQuarterState = &NoQuarter{machine: machine}
	machine.HasQuarterState = &HasQuarterState{machine: machine}
	if count > 0 {
		machine.CurrentState = machine.NoQuarterState
	} else {
		machine.CurrentState = machine.SoldOutState
	}
	return machine
}

func (m *GumballMachine) getLocation() string {
	return m.Location
}
func (m *GumballMachine) setState(s *State) {
	m.CurrentState = *s
} //设置状态
func (m *GumballMachine) insertQuarter() {
	m.CurrentState.insertQuarter()
} //现在是已经投币状态
func (m *GumballMachine) ejectQuarter() {
	m.CurrentState.ejectQuarter()
} //现在是return Quarter状态
func (m *GumballMachine) turnCrank() {
	m.CurrentState.turnCrank()
} //现在是转动了手柄
func (m *GumballMachine) getcout() int {
	return m.count
}
func (m *GumballMachine) release() {
	m.count--
}
func (m *GumballMachine) getState() State {
	fmt.Println(m.CurrentState)
	return m.CurrentState

}

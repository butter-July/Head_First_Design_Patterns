package main

func main() {
	machine := NewGumballMachine(1)
	machine.insertQuarter() //这里为什么显示you can insert quarter--因为初始状态下是没有投币的,但是我要怎么投币呢
	machine.turnCrank()     //这里是you can turn,candy is coming....--这个说明已经又币了 you need to pay first
	//换了hasQuarterState中 turnCrank()的hasQuarter.machine.CurrentState.dispense()
	//	hasQuarter.machine.setState(hasQuarter.machine.NoQuarterState)位置又变成了ou can turn,candy is coming....
	//you need to turn the crank!---好奇怪
	//没有用到出糖果的
	/*
	c := machine.getcout()
	fmt.Println(c)
	machine.insertQuarter() //这里不应该没了嘛
	machine.turnCrank()
	cc := machine.getcout()
	*/
	//fmt.Println(cc) //为什么没变啊
	//单独用这个会变啊,哪错了.
	//明天再看叭..............................
}

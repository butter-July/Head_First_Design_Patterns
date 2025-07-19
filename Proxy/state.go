package main

type State interface {
	insertQuarter() //投币
	ejectQuarter()  //投过了,返回
	turnCrank()     //转动手柄
	dispense()      //出糖果
} //这个是机器会进行的四个动作的状态,然后机器会对应不同的自己的状态,有没有糖果,有没有投币巴拉巴拉的,然后机器会实现这个接口,根据自己的状态的不同,带来不同的结果,所以应该是机器自己不同的状态分别实现他

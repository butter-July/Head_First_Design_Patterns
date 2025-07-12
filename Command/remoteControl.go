package main

import "fmt"

func main() {

	remote := RemoteControl{}
	light := Light{}
	LigthOn := LightOnCommand{&light}
	Lightoff := LightOffCommand{&light}

	remote.SetCommand(1, &LigthOn, &Lightoff)
	remote.OnButtonWasPressed(1)
	remote.OffButtonWasPressed(1) //只有Off和on添加别的是不是还要添加不同的buttonpressed方法
	remote.UndoButtonWasPressed()
	fmt.Println(remote.getClassName(light))

	ceilingFan := CeilingFan{}
	ceilingfanon := CeilingFanOn{&ceilingFan}
	ceilingFanoff := CeilingFanOff{&ceilingFan}
	remote.SetCommand(2, &ceilingfanon, &ceilingFanoff)
	remote.OnButtonWasPressed(2)
	remote.OffButtonWasPressed(2)
	remote.UndoButtonWasPressed()
	ceilingFan.speedControl(3)
	ceilingFan.getSpeed()
	fmt.Println(remote.getClassName(ceilingFan))
	stere := Stere{}
	stereon := StereOnOnCommand{&stere}
	stereoff := StereOffCommand{&stere}
	remote.SetCommand(3, &stereon, &stereoff)
	remote.OnButtonWasPressed(3)
	remote.OffButtonWasPressed(3)
	fmt.Println(remote.getClassName(3))
	remote.UndoButtonWasPressed()

}

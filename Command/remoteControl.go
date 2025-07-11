package main

import "fmt"

func main() {
	remote := SimpleRemoteControl{}
	light := Light{}
	LigthOn := LightOnCommand{&light}
	Lightoff := LightOffCommand{&light}
	remote.SetCommand(1, &LigthOn, &Lightoff) //这里不对应该显示on和off
	remote.OnButtonWasPressed(1)
	fmt.Println(remote.getClassName(light))
	ceilingFan := CeilingFan{}
	ceilingfanon := CeilingFanOn{&ceilingFan}
	ceilingFanoff := CeilingFanOff{&ceilingFan}
	remote.SetCommand(2, &ceilingfanon, &ceilingFanoff)
	remote.OnButtonWasPressed(2)
	fmt.Println(remote.getClassName(ceilingFan))
	stere := Stere{}
	stereon := StereOnOnCommand{&stere}
	stereoff := StereOffCommand{&stere}    //其他的没进去 and  同上
	remote.SetCommand(3, &stereon, &stereoff)
	remote.OnButtonWasPressed(3)
	

}

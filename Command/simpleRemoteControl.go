package main

import (
	"fmt"
	"reflect"
)

type SimpleRemoteControl struct {
	//slot Command
	onCommand  [10]Command
	offCommand [10]Command
}

func (s *SimpleRemoteControl) Command(command [10]Command) {
	for i := 0; i <= 9; i++ {
		s.offCommand[i] = command[i]
		s.offCommand[i] = command[i]
	}
}
func (s *SimpleRemoteControl) SetCommand(slot int, oncommand Command, offCommand Command) {
	s.onCommand[slot] = oncommand
	s.offCommand[slot] = offCommand
}
func (s *SimpleRemoteControl) OnButtonWasPressed(slot int) {
	s.onCommand[slot].execute()
}

func (s *SimpleRemoteControl) OffButtonWasPressed(slot int) {
	s.offCommand[slot].execute()
}

func (s *SimpleRemoteControl) ToString() {
	r := fmt.Sprintf("\n-----Remote Control-----\n")
	for i := range s.onCommand {
		if s.onCommand[i] == nil {
			continue
		}
		onClass := s.getClassName(s.onCommand[i])
		offClass := s.getClassName(s.offCommand[i])
		r += fmt.Sprintf("[slot %d]%s   %s", i, onClass, offClass)
	}
}
func (s *SimpleRemoteControl) getClassName(myVar interface{}) string { //这段代码的核心是通过反射动态获取变量的类型名称，尤其适用于需要处理多种类型或指针类型的场景。
	if t := reflect.TypeOf(myVar); t.Kind() == reflect.Ptr {
		return t.Elem().Name() //是指针那么返回的是他指向了谁的类型名称(自己定义的名字)
	} else {
		return t.Name() //不是指针就直接返回了名字(自己定义的名字)
	}
}

/*
func (s *SimpleRemoteControl) SetCommand(command Command) {
	s.slot = command
}//传进来的是Command接口类型的变量,传的是LightOn--实现接口中方法的结构体
func (s *SimpleRemoteControl) buttonWasPressed() {
	s.slot.execute()
}
*/

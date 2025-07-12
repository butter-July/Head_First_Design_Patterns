package main

import (
	"fmt"
)

type Command interface {
	execute()
	undo()
}
type Light struct {
}
type LightOnCommand struct {
	light *Light
} //他的结构体要有light才能在execute里面调用on的命令,相当与"服务员"的作用,调用者
type LightOffCommand struct {
	light *Light
} //他的结构体要有light才能在execute里面调用on的命令,相当与"服务员"的作用,调用者
func (l *LightOffCommand) execute() {
	l.light.off()
}

func (l *LightOnCommand) execute() {
	l.light.on()
}
func (l *LightOffCommand) undo() {
	l.light.on()
}
func (l *LightOnCommand) undo() {
	l.light.off()
}
func (l *Light) on() {
	fmt.Println("Light is ON")
}

func (l *Light) off() {
	fmt.Println("Light is off!")
}

type Stere struct {
	volume int
}
type StereOnOnCommand struct {
	stere *Stere
}
type StereOffCommand struct {
	Stere *Stere
}
type StereSetCd struct {
	stere *Stere
}
type StereSetVolume struct {
	stere *Stere
}

func (s *Stere) on() {
	fmt.Println("on")
}
func (s *StereOnOnCommand) execute() {
	s.stere.on()
}
func (s *Stere) OFF() {
	fmt.Println("off")
}
func (s *StereOffCommand) execute() {
	s.Stere.OFF()
}
func (s *Stere) setCd() {
	fmt.Println("setCd")
}
func (s *StereSetCd) execute() {
	s.stere.setCd()
}
func (s *Stere) setVolume(volume int) {
	s.volume = volume
}
func (s *StereSetVolume) execute() {
	i, _ := fmt.Scanln()
	s.stere.setVolume(i)
}
func (s *StereOffCommand) undo() {
	s.Stere.on()
}
func (s *StereOnOnCommand) undo() {
	s.stere.OFF()
}

type GarageDoor struct {
}
type GarageDoorUp struct {
	garageDoor *GarageDoor
}
type GarageDoorDown struct {
	garageDoor *GarageDoor
}

func (g *GarageDoor) UP() {
	fmt.Println("GarageDoor UP")
}
func (g *GarageDoorUp) execute() {
	g.garageDoor.UP()
}
func (g *GarageDoor) Down() {
	fmt.Println("GarageDoor Down")
}
func (g *GarageDoorDown) execute() {
	g.garageDoor.Down()
}
func (g *GarageDoorUp) undo() {
	g.garageDoor.Down()
}
func (g *GarageDoorDown) undo() {
	g.garageDoor.UP()
}

type CeilingFan struct {
}
type CeilingFanOn struct {
	ceilinFan *CeilingFan
}
type CeilingFanOff struct {
	ceilingFan *CeilingFan
}

func (c *CeilingFan) on() {
	fmt.Println("ceilingfan on")
}
func (c *CeilingFanOn) execute() {
	c.ceilinFan.on()
}
func (c *CeilingFan) off() {
	fmt.Println("ceilingfan off ")
}
func (c *CeilingFanOff) execute() {
	c.ceilingFan.off()
}
func (c *CeilingFanOn) undo() {
	c.ceilinFan.off()
}
func (c *CeilingFanOff) undo() {
	c.ceilingFan.on()
}

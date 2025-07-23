package main

import (
	"fmt"
)

type GumballRemote struct {
	machine *GumballMachine
}

func (g *GumballRemote) InsertQuarter(args *struct{}, reply *string) error {
	g.machine.insertQuarter()
	*reply = "Inserted Quarter"
	return nil
}
func (g *GumballRemote) EjectQuarter(args *struct{}, reply *string) error {
	g.machine.ejectQuarter()
	*reply = "EjectQuarter"
	return nil
}
func (g *GumballRemote) TurnCrank(ags *struct{}, reply *string) error {
	g.machine.turnCrank()
	*reply = "Turned Crank"
	return nil
}
func (g *GumballRemote) GetState(args *struct{}, reply *string) error {
	state := fmt.Sprintf("%v", g.machine.getState())
	*reply = "Current State" + state
	return nil
}
func (g *GumballRemote) GetCount(args *struct{}, reply *int) error {
	*reply = g.machine.getcout()
	return nil
}

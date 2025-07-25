package main

import "fmt"

type MarllardDuck struct {
}

func (m *MarllardDuck) quack() {
	fmt.Println("Quack")
}

type RedheadDuck struct {
}

func (r *RedheadDuck) quack() {
	fmt.Println("Quack")
}

type DuckCall struct {
}

func (d *DuckCall) quack() {
	fmt.Println("Kwak")
}

type RubberDuck struct {
}

func (rr *RubberDuck) quack() {
	fmt.Println("Squeak")
}

type AbstractDuckFactory interface {
	createMallardDuck() Quackable
	createRedheadDuck() Quackable
	createDuckCall() Quackable
	createRubberDuck() Quackable
	Counter() int
} //一个工厂来制造鸭子,但是感觉乱七八糟 ,奇奇怪怪的一大堆
type DuckFactory struct {
}

func (d *DuckFactory) createMallardDuck() Quackable {
	return &MarllardDuck{}
}

func (d *DuckFactory) createRedheadDuck() Quackable {
	return &RedheadDuck{}
}

func (d *DuckFactory) createDuckCall() Quackable {
	return &DuckCall{}
}
func (d *DuckFactory) createRubberDuck() Quackable {
	return &RubberDuck{}
}
func (d *DuckFactory) Counter() int {
	return 0
}

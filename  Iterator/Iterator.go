package main

type Iterator interface {
	HasNext() bool
	next() *MenuItem
}

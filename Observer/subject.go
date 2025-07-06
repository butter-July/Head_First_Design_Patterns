package main

type Suject interface {
	RegisterObserver(observer Observer)
	RemoveObserver(observer Observer)
	notifyObserver()
}

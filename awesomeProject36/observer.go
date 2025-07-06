package main

type Observer interface {
	update(temp, humidity, pressure float64)
}

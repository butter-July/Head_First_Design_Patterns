package main

import "fmt"

func main() {
	hometheater := &HomeTheaterFacade{}
	fmt.Println("-----------begin-------------")
	hometheater.watchMovie()
	fmt.Println()
	fmt.Println("-----------------end-------------")
	hometheater.endMovie()
}

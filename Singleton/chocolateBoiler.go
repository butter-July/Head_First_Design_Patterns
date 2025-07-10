package main

import (
	"fmt"
	"sync"
)

var lock = &sync.Mutex{}

type ChocolateBoiler struct {
}

var singleInstance *ChocolateBoiler

func chocolateBoiler() *ChocolateBoiler {
	if singleInstance == nil {
		lock.Lock()
		defer lock.Unlock()
		if singleInstance == nil {
			fmt.Println("the ChocolateBoiler is empty")
			singleInstance = &ChocolateBoiler{}
		} else {
			fmt.Println("the ChocolateBoiler is not empty ")
		}
	} else {
		fmt.Println("the ChocolateBoiler is not empty.")
	}

	return singleInstance
}

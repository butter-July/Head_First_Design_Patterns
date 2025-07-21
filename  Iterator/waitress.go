package main

import "fmt"

type Waitress struct {
	pancakeHouseMenu Menu
	dinerMenu        Menu
}

func NewWaitress(pancakeHouseMenu Menu, dinerMenu Menu) *Waitress {
	return &Waitress{
		pancakeHouseMenu: pancakeHouseMenu,
		dinerMenu:        dinerMenu,
	}
}

func (w *Waitress) PrintMenu() {
	pancakeIterator := w.pancakeHouseMenu.createIter()
	dinerIterator := w.dinerMenu.createIter()

	fmt.Println("MENU-------pancake")
	w.printMenu(pancakeIterator)

	fmt.Println("Menu-----diner")
	w.printMenu(dinerIterator)
}

func (w *Waitress) printMenu(iterator Iterator) {
	for iterator.HasNext() {
		menuItem := iterator.next()
		fmt.Printf("%s, $%.2f -- %s\n", menuItem.name, menuItem.price, menuItem.description)
	}
}

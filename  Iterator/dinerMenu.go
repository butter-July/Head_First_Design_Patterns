package main

import (
	"errors"
)

type DinerMenu struct {
	NumberOfItems int
	MenuItems     [6]*MenuItem
}

func NewDinerMenu() *DinerMenu {
	dinerMenu := DinerMenu{}
	dinerMenu.AddItem("Vegetarian BLT", "(Fakin') Bacon with lettuce & tomato on whole wheat", true, 2.99)
	dinerMenu.AddItem("BLT", "Bacon with lettuce & tomato on whole wheat", false, 2.99)
	dinerMenu.AddItem("Soup of the day", "Soup of the day, with a side of potato salad", false, 3.29)
	dinerMenu.AddItem("Hotdog", "A hot dog, with sauerkraut, relish, onions, topped with cheese", false, 3.05)
	return &dinerMenu
}

func (dm *DinerMenu) AddItem(name string, description string, vegetarian bool, price float64) error {
	if dm.NumberOfItems >= 6 {
		return errors.New("Sorry, menu is full! Can't add item to menu")
	}
	menuItem := &MenuItem{name, description, vegetarian, price}
	dm.MenuItems[dm.NumberOfItems] = menuItem
	dm.NumberOfItems++
	return nil
}
func (diner *DinerMenu) GetMenuItems() [6]*MenuItem {
	return diner.MenuItems
}

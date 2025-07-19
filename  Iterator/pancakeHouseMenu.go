package main

type PancakeHouseMenu struct {
	menuItem []MenuItem
}

func NewPancakeMenu() *PancakeHouseMenu {
	menu := &PancakeHouseMenu{}
	menu.addItem("K&B's Pancake Breakfast", "Pancakes with scrambled eggs and toast", true, 2.99)
	menu.addItem("Regular Pancake Breakfast", "Pancakes with fried eggs, sausage", false, 2.99)
	menu.addItem("Blueberry Pancakes", "Pancakes made with fresh blueberries", true, 3.49)
	menu.addItem("Waffles", "Waffles with your choice of blueberries or strawberries", true, 3.59)
	return menu
}
func (pancake *PancakeHouseMenu) addItem(name string, description string, vegetarian bool, price float64) {
	menuItem := MenuItem{name: name, description: description, vegetarian: vegetarian, price: price}
	pancake.menuItem = append(pancake.menuItem, menuItem)
}
func (pancake *PancakeHouseMenu) getMenuItems() MenuItem {
	return MenuItem{}
}

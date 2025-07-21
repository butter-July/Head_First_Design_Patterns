package main

type PancakeHouseMenu struct {
	menuItem []MenuItem
}

func (pancake *PancakeHouseMenu) addItem(name string, description string, vegetarian bool, price float64) {
	menuItem := MenuItem{name: name, description: description, vegetarian: vegetarian, price: price}
	pancake.menuItem = append(pancake.menuItem, menuItem)
}
func (pancake *PancakeHouseMenu) createIter() Iterator {
	return &PancakeHouseIterator{menu: pancake.menuItem}
}
func (pancake *PancakeHouseMenu) getMenuItems() MenuItem {
	return MenuItem{}
}

package main

func main() {
	pancakeHouseMenu := &PancakeHouseMenu{}
	pancakeHouseMenu.addItem("K&B's Pancake Breakfast", "Pancakes with scrambled eggs and toast", true, 2.99)
	pancakeHouseMenu.addItem("Regular Pancake Breakfast", "Pancakes with fried eggs, sausage", false, 2.99)

	dinerMenu := &DinerMenu{}
	dinerMenu.AddItem("Vegetarian BLT", "(Fakin') Bacon with lettuce & tomato on whole wheat", true, 2.99)
	dinerMenu.AddItem("BLT", "Bacon with lettuce & tomato on whole wheat", false, 2.99)

	waitress := NewWaitress(pancakeHouseMenu, dinerMenu)
	waitress.PrintMenu()
}

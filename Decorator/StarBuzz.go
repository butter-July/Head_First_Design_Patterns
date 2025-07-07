package main

import "fmt"

func main() {
	beverage1 := &Espresso{}
	fmt.Println("des:", beverage1.getDescription())
	fmt.Println("cost:", beverage1.Cost())

	beverage2 := NewMocha(NewMocha(NewWhip(&DarkRoast{})))
	fmt.Println("des:", beverage2.getDescription())
	fmt.Println("cost:", beverage2.Cost())
	beverage3 := NewSoy(NewMocha(NewWhip(&HouseBlend{})))
	fmt.Println("des:", beverage3.getDescription())
	fmt.Println("cost:", beverage3.Cost())

}

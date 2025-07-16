package main

import "fmt"

func main() {
	fmt.Println("Making coffee:")
	coffee := &Coffee{}
	PrepareRecipe(coffee)
	fmt.Println("Making tea:")
	tea := &Tea{}
	PrepareRecipe(tea)
}

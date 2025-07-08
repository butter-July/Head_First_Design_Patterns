package main

import "fmt"

func (p *Pizza) prepare() {
	fmt.Printf("Preparing %s", p.name)
	fmt.Printf("Tossing %s dough\n", p.dough)
	fmt.Printf("Adding %s sauce\n", p.sauce)
	fmt.Print("Adding toppings: \n")
	for _, t := range p.toppings {
		fmt.Printf("\t%s, ", t)
	}
	fmt.Println()
}
func (p *Pizza) bake() {
	fmt.Println("Bake for 25 minutes at 350 ")
}
func (p *Pizza) cut() {
	fmt.Println("Cutting the Pizza into diagonal slices")
}
func (p *Pizza) box() {
	fmt.Println("Place Pizza in official PizzaStore box")
}
func (p *Pizza) getName() string {
	return p.name
}

type nyStyleCheesePizza struct {
	*Pizza
}

func newNYStyleCheesePizza() PPizza {
	p := &Pizza{name: "NY Style Sauce and Cheese Pizza",
		dough:    "Thin Crust Gough",
		sauce:    "Marinara Sauce",
		toppings: []string{"Grated Reggiano Cheese"},
	}
	return &nyStyleCheesePizza{Pizza: p}
}

func newNYStylePepperoniPizza() PPizza {
	p := &Pizza{name: "NY Style Sauce and Pepperoni Pizza",
		dough:    "Crust",
		sauce:    "Marinara sauce",
		toppings: []string{"Sliced Pepperoni,Sliced Onion,Grated parmesan cheese"},
	}
	return &nyStyleCheesePizza{Pizza: p}
}

func newNYStyleClamPizza() PPizza {
	p := &Pizza{name: "NY Style Sauce and Clam Pizza",
		dough:    "Thin crust",
		sauce:    "White garlic sauce",
		toppings: []string{"Grated parmesan cheese"},
	}
	return &nyStyleCheesePizza{Pizza: p}
}
func newNYStyleVeggiePizza() PPizza {
	p := &Pizza{name: "NY Style Sauce and Veggie Pizza",
		dough:    "TCrust",
		sauce:    "Marinara sauce",
		toppings: []string{"Shredded mozzarella,Grated parmesan,Diced onion,Sliced mushrooms,Sliced red pepper,Sliced black olives"},
	}
	return &nyStyleCheesePizza{Pizza: p}
}

/*
type chicagoStyleCheesePizza struct {
	*Pizza
}

func newChicagoStyleCheesePizza() PPizza {
	p := &Pizza{
		name:     "Chicago Style Deep Dish Cheese Pizza",
		dough:    "Extra Thick Crust Dough",
		sauce:    "Plum Tomato Sauce",
		toppings: []string{"Shredded Mozzarella Cheese"},
	}
	return &chicagoStyleCheesePizza{
		Pizza: p,
	}
}

func newChicagoStylePepperoniPizza() PPizza {
	p := &Pizza{name: "NY Style Sauce and Pepperoni Pizza",
		dough:    "Crust",
		sauce:    "Marinara sauce",
		toppings: []string{"Sliced Pepperoni,Sliced Onion,Grated parmesan cheese"},
	}
	return &nyStyleCheesePizza{Pizza: p}
}

func newChicagoStyleClamPizza() PPizza {
	p := &Pizza{name: "NY Style Sauce and Clam Pizza",
		dough:    "Thin crust",
		sauce:    "White garlic sauce",
		toppings: []string{"Grated parmesan cheese"},
	}
	return &nyStyleCheesePizza{Pizza: p}
}
func newChicagoStyleVeggiePizza() PPizza {
	p := &Pizza{name: "NY Style Sauce and Veggie Pizza",
		dough:    "TCrust",
		sauce:    "Marinara sauce",
		toppings: []string{"Shredded mozzarella,Grated parmesan,Diced onion,Sliced mushrooms,Sliced red pepper,Sliced black olives"},
	}
	return &nyStyleCheesePizza{Pizza: p}
}
*/

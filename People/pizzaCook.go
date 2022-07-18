package people

import (
	Pizza "com/github/spencerk/PizzariaParallel/Pizza"
	"fmt"
)

//pizzaPie alias exists coz it's defined in the package scope in customer.go

const OVEN_CAPACITY int = 3 //How many pizzas can fit into the oven at once

func PizzaCook(orders <-chan string, customer chan<- *pizzaPie) {
	ovenIn := make(chan *pizzaPie, OVEN_CAPACITY)
	ovenOut := make(chan *pizzaPie)

	go fireUpOven(ovenIn, ovenOut)

	//Put orders into the oven, and take them out when they're ready
	//Allows cook to multitask by making more pizzas and putting them into the oven while others cook
	go putPizzaInOven(orders, ovenIn)
	go takePizzaOutOfOven(ovenOut, customer)
}

func putPizzaInOven(orders <-chan string, oven chan<- *pizzaPie) {
	//Get order from customer. Put pizza into oven
	for order := range orders {
		pie := Pizza.New(order)
		fmt.Printf("Pizza cook: Putting %s pizza in the oven.\n", pie.TypeOfPizza)
		oven <- pie
	}
}

func takePizzaOutOfOven(oven <-chan *pizzaPie, customer chan<- *pizzaPie) {
	//Take ready pizza out of oven. Give it to customer
	for pie := range oven {
		fmt.Printf("Pizza cook: Pizza is ready. Handing %s pizza over to customer.\n", pie.TypeOfPizza)
		customer <- pie
	}
}

func fireUpOven(pieIn <-chan *pizzaPie, pieOut chan<- *pizzaPie) {
	//Cook puts the pizza in the oven and takes it out when ready
	for pie := range pieIn {
		fmt.Printf("Oven: The %s pizza cooking away.\n", pie.TypeOfPizza)
		pie.Cook()
		pieOut <- pie
	}
}

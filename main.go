package main

import (
	People "com/github/spencerk/PizzariaParallel/People"
	Pizza "com/github/spencerk/PizzariaParallel/Pizza"
)

type pizzaPie = Pizza.PizzaPie

func main() {
	orderChan := make(chan string, People.NUM_CUSTOMERS)    //Channel used to order pizza
	pizzaChan := make(chan *pizzaPie, People.NUM_CUSTOMERS) //Channel used to recieve pizza
	quitChan := make(chan int)                              //Channel used to tell main when customers leave the pizzaria (when to quit)

	//Make customers
	for i := 0; i < People.NUM_CUSTOMERS; i++ {
		go People.Customer(orderChan, pizzaChan, quitChan)
	}

	//Make pizza cook
	go People.PizzaCook(orderChan, pizzaChan)

	//Wait for customers to finish eating all the pizza
	for i := 0; i < People.NUM_CUSTOMERS; i++ {
		<-quitChan
	}
}

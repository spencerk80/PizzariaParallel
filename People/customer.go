package people

import (
	Pizza "com/github/spencerk/PizzariaParallel/Pizza"
	"fmt"
	"math/rand"
	"time"
)

type pizzaPie = Pizza.PizzaPie

const NUM_CUSTOMERS int = 2     //Sets how many customers are in the pizzaria at one time
const NUM_PIZZAS_WANTED int = 5 //Number of pizzas each customer will order

func randPizzaType() (typeOfPizza string) {
	switch rand.Intn(4) {
	case 0:
		typeOfPizza = "pepperoni"
	case 1:
		typeOfPizza = "cheese"
	case 2:
		typeOfPizza = "meat lovers"
	case 3:
		typeOfPizza = "supreme"
	}

	return
}

func Customer(orderChan chan<- string, pizzaria <-chan *pizzaPie, closeChan chan<- int) {
	//The customer wants 5 pizzas, but will order back-to-back
	//Remember, each customer is doing this in parallel
	for i := 1; i <= NUM_PIZZAS_WANTED; i++ {
		var pie *pizzaPie
		//Decide on the type of pizza
		wantedPieType := randPizzaType()

		//Place order
		fmt.Printf("Customer: Ordering a %s pizza.\n", wantedPieType)
		orderChan <- wantedPieType

		//Wait on pizza
		fmt.Println("Custmoer: Waiting for my pizza...")

		//Recieve and consume
		pie = <-pizzaria
		fmt.Printf("Custmoer: Got my %s pizza! Let's eat!\n", pie.TypeOfPizza)
		pie.Eat()

		//Stop ordering if last pie
		if i != NUM_PIZZAS_WANTED {
			//Be full for a moment, then get hungi and order again
			time.Sleep(3 * time.Second)
			fmt.Println("Customer: Hungry again... Gonna eat order another.")
		}
	}

	//Customer is finally full
	fmt.Println("Customer: I'm full now! Thanks for the pizzas :)")
	closeChan <- 0
}

package pizza

import "time"

type PizzaPie struct {
	TypeOfPizza string
}

func New(typeOfPizza string) *PizzaPie {
	return &PizzaPie{TypeOfPizza: typeOfPizza}
}

func (p *PizzaPie) Cook() {
	time.Sleep(4 * time.Second)
}

func (p *PizzaPie) Eat() {
	time.Sleep(8 * time.Second)
}

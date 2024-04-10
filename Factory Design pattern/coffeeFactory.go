package main

type CoffeeFactory struct{}

func (r CoffeeFactory) GetCoffee(coffeeType string) Coffee {
	if coffeeType == "cappuccino" {
		return &Cappuccino{}
	} else if coffeeType == "expresso" {
		return &Expresso{}
	} else {
		return &Robusta{}
	}
}

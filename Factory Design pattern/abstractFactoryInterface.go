package main

type AbstractCoffeeFactory interface {
	GetCoffee(coffeeType string) Coffee
}

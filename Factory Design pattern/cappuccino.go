package main

import "fmt"

type Cappuccino struct{}

// Implement the brew method for Robusta
func (r Cappuccino) Brew() {
	fmt.Println("Brewing Cappuccino coffee...")
}

// Implement the boil method for Robusta
func (r Cappuccino) Boil() {
	fmt.Println("Boiling Cappuccino coffee...")
}

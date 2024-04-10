package main

import "fmt"

type Expresso struct{}

// Implement the brew method for Robusta
func (r Expresso) Brew() {
	fmt.Println("Brewing Expresso coffee...")
}

// Implement the boil method for Robusta
func (r Expresso) Boil() {
	fmt.Println("Boiling Expresso coffee...")
}

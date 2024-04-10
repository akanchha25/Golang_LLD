package main

import "fmt"

type Robusta struct{}

// Implement the brew method for Robusta
func (r Robusta) Brew() {
	fmt.Println("Brewing Robusta coffee...")
}

// Implement the boil method for Robusta
func (r Robusta) Boil() {
	fmt.Println("Boiling Robusta coffee...")
}

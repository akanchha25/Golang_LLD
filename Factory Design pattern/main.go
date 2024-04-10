package main

import "fmt"

func main() {
	// Create a new coffee factory
	factory := CoffeeFactory{}

	// Create a new coffee server with the factory
	server := NewCoffeeServer(factory)

	// Serve different types of coffee
	cappuccino := server.Serve("cappuccino")
	fmt.Println("Served cappuccino:", cappuccino)

	expresso := server.Serve("expresso")
	fmt.Println("Served expresso:", expresso)

	robusta := server.Serve("robusta")
	fmt.Println("Served robusta:", robusta)
}

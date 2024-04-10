package main

type CoffeeServer struct {
	Factory AbstractCoffeeFactory
}

// NewCoffeeServer creates a new instance of CoffeeServer with the provided factory.
func NewCoffeeServer(factory AbstractCoffeeFactory) *CoffeeServer {
	return &CoffeeServer{
		Factory: factory,
	}
}

func (s *CoffeeServer) Serve(coffeeType string) Coffee {
	coffee := s.Factory.GetCoffee(coffeeType)
	coffee.Brew()
	coffee.Boil()
	return coffee
}

package data

type DataAccessor struct {}


func (da DataAccessor) GetFoodItemsWithName(name string) DataAccessResult {
	return DataAccessResult{}
}

func (da *DataAccessor) GetCart(user User) DataAccessResult {
	// Implement the logic to get the cart for the user
	return DataAccessResult{}
}

func (da *DataAccessor) AddItemToCart(user User, foodItem FoodItem) {
	// Implement the logic to add an item to the cart
}

func (da *DataAccessor) DeleteItemFromCart(user User, foodItem FoodItem) {
	// Implement the logic to delete an item from the cart
}

func (da *DataAccessor) CheckOutCart(user User) {
	// Implement the logic to check out the cart
}
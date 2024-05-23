package data

type CartItem struct{
	FoodItem FoodItem
	Quantity int
}

func NewCartItem(foodItem FoodItem, quantity int) *CartItem {
	return &CartItem{
		FoodItem: foodItem,
		Quantity: quantity,
	}

}

func (c *CartItem) GetFoodItem() FoodItem {
	return c.FoodItem
}

func (c *CartItem) GetQuantity() int {
	return c.Quantity
}
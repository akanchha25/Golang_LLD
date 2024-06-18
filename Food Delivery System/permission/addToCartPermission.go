package permission

import (
	"food_delivery_system/searchers"
	"food_delivery_system/data"
)


type AddToCartPermission struct{
	User *data.User
	FoodItem *data.FoodItem
}

func NewAddToCartPermission(user *data.User, foodItem *data.FoodItem) *AddToCartPermission {
	return &AddToCartPermission{
		User: user,
		FoodItem: foodItem,
	}
}

// IsPermitted method
func (p *AddToCartPermission) IsPermitted() bool {
	if !p.FoodItem.IsItemAvailable() {
		return false
	}
	restaurant, error := searchers.NewRestaurantSearcher().SearchById(p.FoodItem.RestaurantID)
	if error != nil {
		return false
	}
}
package filters

import (
	"food_delivery_system/data"
)

type FoodItemFilter interface {
     Filter(foodItem data.FoodItem) bool
}


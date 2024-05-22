package filters

import (
	"food_delivery_system/data"
)

type RestaurantFilter interface {
	RestaurantFilteration(restaurant data.Restaurant) bool
}
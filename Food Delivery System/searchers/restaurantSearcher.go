package searchers

import (
	"food_delivery_system/data"
	"food_delivery_system/filters"
)

type RestaurantSearcher struct{}

func NewRestaurantSearcher() *RestaurantSearcher {
	return &RestaurantSearcher{}
}

func (s *RestaurantSearcher) Search(restaurantName string, filterList []filters.RestaurantFilter) []data.Restaurant {
	// Mock database of restaurants
	restaurants := []data.Restaurant{
		// Populate with sample data
	}

	// Apply filters
	for _, filter := range filterList {
		var filtered []data.Restaurant
		for _, restaurant := range restaurants {
			if filter.RestaurantFilteration(restaurant) {
				filtered = append(filtered, restaurant)
			}
		}
		restaurants = filtered
	}

	// If restaurantName is provided, filter by name
	if restaurantName != "" {
		var filtered []data.Restaurant
		for _, restaurant := range restaurants {
			if restaurant.Name == restaurantName {
				filtered = append(filtered, restaurant)
			}
		}
		return filtered
	}

	return restaurants
}


func (s *RestaurantSearcher) SearchById(id int) (data.Restaurant, error) {
	var restaurant data.Restaurant

	return restaurant, nil
}

package apis

import (
	"food_delivery_system/data"
	"food_delivery_system/filters"
	"food_delivery_system/searchers"
)

type RestaurantSearcherAPI struct{}

func NewRestaurantSearcherAPI() *RestaurantSearcherAPI {
	return &RestaurantSearcherAPI{}
}

func (api *RestaurantSearcherAPI) SearchRestaurant(restaurantName string, mealType data.MealType, cuisineTypes []data.CuisineType, rating data.StarRating) []data.Restaurant {
	var filterList []filters.RestaurantFilter

	if mealType != data.MealTypeUnset {
		filterList = append(filterList, filters.NewMealTypeFilter(mealType))
	}
	if cuisineTypes != nil {
		filterList = append(filterList, filters.NewCuisineTypeFilter(cuisineTypes))
	}
	if rating != data.StarRatingUnset { // Correct check for unset value
		filterList = append(filterList, filters.NewStarRatingFilter(rating))
	}

	return searchers.NewRestaurantSearcher().Search(restaurantName, filterList)
}

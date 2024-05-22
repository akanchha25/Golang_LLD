package filters

import "food_delivery_system/data"


// CuisineTypeFilter is a filter that checks if a FoodItem has a specific CuisineType.
type CuisineTypeFilter struct{
	cuisineTypes []data.CuisineType

}

// NewCuisineTypeFilter creates a new CuisineTypeFilter.
func NewCuisineTypeFilter(cuisineTypes []data.CuisineType) *CuisineTypeFilter {
	return &CuisineTypeFilter{
		cuisineTypes: cuisineTypes,
	}
}

// Filter checks if the FoodItem has one of the specified CuisineTypes
func (f CuisineTypeFilter) Filter(foodItem data.FoodItem) bool {
	for _, cuisineType := range f.cuisineTypes {
		if foodItem.CuisineType == cuisineType {
			return true
		}
	}

	return false
}

func (f CuisineTypeFilter) RestaurantFilteration(restaurant data.Restaurant) bool {
	cuisines := restaurant.GetCuisineType()
	for _, cuisineType := range f.cuisineTypes {
		for _, cuisine := range cuisines {
			if(cuisine == cuisineType){
				return true
			} 
		}
	}
	return false
}

// Assert that CuisineTypeFilter implements FoodItemFilter interface
// var _ FoodItemFilter = (*CuisineTypeFilter)(nil)
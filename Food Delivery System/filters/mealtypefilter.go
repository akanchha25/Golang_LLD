package filters

import "food_delivery_system/data"


// MealTypeFilter is a filter that checks the MealType of a FoodItem.
type MealTypeFilter struct{
	mealType data.MealType
}

// NewMealTypeFilter creates a new MealTypeFilter.
func NewMealTypeFilter(mealType data.MealType) *MealTypeFilter {
	return &MealTypeFilter{
		mealType: mealType,
	}
}

// Filter checks if the FoodItem has the specified MealType
func (f MealTypeFilter) Filter(foodItem data.FoodItem) bool {
	return foodItem.MealType == f.mealType

}

// Assert that MealTypeFilter implements FoodItemFilter interface.
var _ FoodItemFilter = (*MealTypeFilter)(nil)
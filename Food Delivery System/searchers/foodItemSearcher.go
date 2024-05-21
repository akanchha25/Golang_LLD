package searchers

import (
	"errors"
	"food_delivery_system/data"
	"food_delivery_system/filters"
)

type FoodItemSearcher struct{}

func NewFoodItemSearcher() *FoodItemSearcher {
	return &FoodItemSearcher{}
}

func (searcher *FoodItemSearcher) Search(foodItemName string, filters []filters.FoodItemFilter) ([] data.FoodItem, error) {
	if foodItemName == "" || filters == nil {
		return nil, errors.New("missing params")
	}

	dataAccessor := data.DataAccessor{}

	dataAccessResult := dataAccessor.GetFoodItemsWithName(foodItemName)

	foodItems := data.ConvertToFoodItems(dataAccessResult)

	for _, filter := range filters {
		var filteredFoodItems []data.FoodItem
		for _, foodItem := range foodItems {
			if filter.Filter(foodItem) {
				filteredFoodItems = append(filteredFoodItems, foodItem)
			}
		}

		foodItems = filteredFoodItems
	}

	return foodItems, nil
}
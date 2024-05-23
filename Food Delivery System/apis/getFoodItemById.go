package apis

import (
	"food_delivery_system/data"
	"food_delivery_system/searchers"
	"log"
)

type GetFoodItemByIdAPI struct{}

func NewGetFoodItemByIdAPI() *GetFoodItemByIdAPI {
	return &GetFoodItemByIdAPI{}
}

func (api *GetFoodItemByIdAPI) GetFoodItemById(id int) (data.FoodItem, error) {
	foodItem, err := searchers.NewFoodItemSearcher().SearchById(id)
	if err != nil {
		log.Printf("Error retrieving food item by ID %d: %v", id, err)
		return data.FoodItem{}, err
	}
	return foodItem, nil
}

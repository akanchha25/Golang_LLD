package apis

import (
	"food_delivery_system/data"
	"food_delivery_system/searchers"
	"log"
)

type GetRestaurantByIdAPI struct{}

func NewGetRestaurantByIdAPI() *GetRestaurantByIdAPI {
	return &GetRestaurantByIdAPI{}
}

func (api *GetRestaurantByIdAPI) GetRestaurantById(id int) (data.Restaurant, error) {
	 restaurant, err := searchers.NewRestaurantSearcher().SearchById(id)
     // logic to get restaurantById from database

	 if err != nil {
		log.Printf("Error retrieving restaurant by ID %d: %v", id, err)
		return data.Restaurant{}, err

	 }
	 return restaurant, nil
}
package apis

import (
	"errors"
	"food_delivery_system/data"
	"food_delivery_system/manager"
	"food_delivery_system/searchers"
)

type AddToCartAPI struct {
	userManager     *manager.UserManager
	foodItemSearcher *searchers.FoodItemSearcher
	cartManager     *manager.CartManager
}

func NewAddToCartAPI() *AddToCartAPI {
	return &AddToCartAPI{
		userManager:     manager.NewUserManager(),
		foodItemSearcher: searchers.NewFoodItemSearcher(),
		cartManager:     manager.NewCartManager(),
	}
}

func (api *AddToCartAPI) AddToCart(userToken string, foodItemId int) error {
	if userToken == "" || foodItemId < 0 {
		return errors.New("invalid input parameters")
	}

	user, err := api.userManager.GetUserByToken(userToken)
	if err != nil {
		return err
	}

	foodItem, err := api.foodItemSearcher.SearchById(foodItemId)
	if err != nil {
		return err
	}

	err = api.cartManager.AddItemToCart(user, foodItem)
	if err != nil {
		return err
	}

	return nil
}

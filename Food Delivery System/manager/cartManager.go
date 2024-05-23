package manager

import (
	"food_delivery_system/data"
)

type CartManager struct{}

func NewCartManager() *CartManager {
	return &CartManager{}
}

func (cm *CartManager) GetUserCart(user data.User) []data.CartItem {
	dataAccessor := data.DataAccessor{}
	dataAccessResult := dataAccessor.GetCart(user)
	return data.ConvertToCartItems(dataAccessResult)
}

// add 1 unit at a time
func (cm *CartManager) AddItemToCart(user data.User, item data.Item) error {
	permission := factory.PermissionFactory{}.GetAddToCartPermission(user, item)
	if !permission.IsPermitted() {
		return errors.New("permission denied")
	}
	if !cm.IsItemFromSameRestaurant(user, item) {
		return errors.New("your cart contains items from different restaurants")
	}
	data.DataAccessor{}.AddItemToCart(user, item)
	return nil
}

// delete 1 unit
func (cm *CartManager) DeleteItemFromCart(user data.User, item data.Item) error {
	permission := factory.PermissionFactory{}.GetDeleteFromCartPermission(user, item)
	if !permission.IsPermitted() {
		return errors.New("permission denied")
	}
	if !cm.IsItemPresentInCart(user, item) {
		return errors.New("item not in cart")
	}
	data.DataAccessor{}.DeleteItemFromCart(user, item)
	return nil
}

func (cm *CartManager) CheckOutUserCart(user data.User) error {
	permission := factory.PermissionFactory{}.GetCheckoutCartPermission(user)
	if !permission.IsPermitted() {
		return errors.New("permission denied")
	}
	if cm.IsCartEmpty(user) {
		return errors.New("cart empty")
	}
	data.DataAccessor{}.CheckOutCart(user)
	return nil
}

func (cm *CartManager) IsItemFromSameRestaurant(user data.User, item data.Item) bool {
	cartItems := cm.GetUserCart(user)
	if len(cartItems) == 0 {
		return true
	}
	return cartItems[0].(*data.FoodItem).RestaurantID == item.(*data.FoodItem).RestaurantID
}


func (cm *CartManager) IsItemPresentInCart(user data.User, item data.Item) bool {
	cartItems := cm.GetUserCart(user)
	for _, cartItem := range cartItems {
		if cartItem.GetID() == item.GetID() {
			return true
		}
	}
	return false
}

func (cm *CartManager) IsCartEmpty(user data.User) bool {
	cartItems := cm.GetUserCart(user)
	return len(cartItems) == 0
}
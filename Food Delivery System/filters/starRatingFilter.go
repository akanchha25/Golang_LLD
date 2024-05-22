package filters

import "food_delivery_system/data"


type StarRatingFilter struct{
	starRate data.StarRating
}

func NewStarRatingFilter(starRate data.StarRating) *StarRatingFilter {
	return &StarRatingFilter{
		starRate: starRate,
	}
}

func (f StarRatingFilter) Filter(foodItem data.FoodItem) bool {
	return foodItem.GetStarRating().GetVal() >= int(f.starRate.GetVal())
}

func (f StarRatingFilter) RestaurantFilteration(restaurant data.Restaurant) bool {
	return restaurant.GetStarRating().GetVal() >= f.starRate.GetVal()         
}  

// var _ FoodItemFilter = (*StarRatingFilter)(nil)
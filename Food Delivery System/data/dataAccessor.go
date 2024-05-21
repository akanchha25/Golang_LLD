package data

type DataAccessor struct {}


func (da DataAccessor) GetFoodItemsWithName(name string) DataAccessResult {
	return DataAccessResult{
		FoodItems: []FoodItem{
			{ID: 1, Name: "Pasta", Description: "Delicious Italian pasta", PriceINR: 250.0, MealType: VEG, CuisineType: ITALIAN, StarRating: FIVE, RestaurantID: 101, IsAvailable: true},
            {ID: 2, Name: "Pasta", Description: "Fresh sushi", PriceINR: 300.0, MealType: NON_VEG, CuisineType: ASIAN, StarRating: FOUR, RestaurantID: 102, IsAvailable: true},
        },
	}
}
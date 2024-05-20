package apis

import (
    "food_delivery_system/data"
)

type FoodItemSearcherAPI struct{}

// NewFoodItemSearcherAPI creates a new instance of FoodItemSearcherAPI.
func NewFoodItemSearcherAPI() *FoodItemSearcherAPI {
    return &FoodItemSearcherAPI{}
}

// SearchFoodItems searches for food items based on the given criteria.
func (api *FoodItemSearcherAPI) SearchFoodItems(foodItemName string, mealType data.MealType, cuisines []data.CuisineType, rating data.StarRating) []data.FoodItem {
    // Mock data - in a real implementation, this would come from a database or another data source
    // allFoodItems := []data.FoodItem{
    //     {Name: "Pasta", MealType: data.VEG, Cuisine: data.ITALIAN, Rating: data.FIVE},
    //     {Name: "Sushi", MealType: data.NON_VEG, Cuisine: data.ASIAN, Rating: data.FOUR},
    //     {Name: "Bratwurst", MealType: data.NON_VEG, Cuisine: data.GERMAN, Rating: data.THREE},
    //     {Name: "Paella", MealType: data.NON_VEG, Cuisine: data.SPANISH, Rating: data.FOUR},
    //     {Name: "Margherita Pizza", MealType: data.VEG, Cuisine: data.ITALIAN, Rating: data.FIVE},
    // }

    var result []data.FoodItem
    // for _, item := range allFoodItems {
    //     if (foodItemName == "" || item.Name == foodItemName) &&
    //        (mealType == item.MealType) &&
    //        containsCuisine(cuisines, item.Cuisine) &&
    //        (rating == 0 || item.Rating == rating) {
    //         result = append(result, item)
    //     }
    // }

    return result
	
}

// Helper function to check if a CuisineType is in a slice of CuisineType
func containsCuisine(cuisines []data.CuisineType, cuisine data.CuisineType) bool {
    for _, c := range cuisines {
        if c == cuisine {
            return true
        }
    }
    return false
}

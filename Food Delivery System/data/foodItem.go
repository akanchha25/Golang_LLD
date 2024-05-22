 package data


// Builder Pattern
// FoodItem represents a food item.
type FoodItem struct {
    ID           int
    Name         string
    Description  string
    PriceINR     float64
    MealType     MealType
    CuisineType  CuisineType
    StarRating   StarRating
    RestaurantID int
    IsAvailable  bool
}

// NewFoodItem creates a new instance of FoodItem.
func NewFoodItem(id int, name, description string, priceINR float64, mealType MealType, cuisineType CuisineType, starRating StarRating, restaurantId int, isAvailable bool) *FoodItem {
    return &FoodItem{
        ID:           id,
        Name:         name,
        Description:  description,
        PriceINR:     priceINR,
        MealType:     mealType,
        CuisineType:  cuisineType,
        StarRating:   starRating,
        RestaurantID: restaurantId,
        IsAvailable:  isAvailable,
    }
}



// Getter methods for FoodItem fields

func (f *FoodItem) GetID() int {
    return f.ID
}

func (f *FoodItem) GetName() string {
    return f.Name
}

func (f *FoodItem) GetDescription() string {
    return f.Description
}

func (f *FoodItem) GetPriceINR() float64 {
    return f.PriceINR
}

func (f *FoodItem) GetMealType() MealType {
    return f.MealType
}

func (f *FoodItem) GetCuisineType() CuisineType {
    return f.CuisineType
}

func (f *FoodItem) GetStarRating() StarRating {
    return f.StarRating
}

func (f *FoodItem) GetRestaurantID() int {
    return f.RestaurantID
}

func (f *FoodItem) IsItemAvailable() bool {
    return f.IsAvailable
}
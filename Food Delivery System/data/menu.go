package data

// Menu represents a menu containing a list of FoodItems.
type Menu struct {
    FoodItems []FoodItem
}

// NewMenu creates a new instance of Menu.
func NewMenu(foodItems []FoodItem) *Menu {
    return &Menu{
        FoodItems: foodItems,
    }
}

// GetFoodItems returns the list of FoodItems in the Menu.
func (m *Menu) GetFoodItems() []FoodItem {
    return m.FoodItems
}

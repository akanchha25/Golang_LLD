package data

type Restaurant struct {
	ID             int
	Name           string
	Description    string
	BusinessHours  BusinessHours
	MealType       MealType   
	CuisineType    []CuisineType
	StarRating     StarRating
	Menu           Menu
	Address        Address
}

func NewRestaurant(id int, name, description string, businessHours BusinessHours, mealType MealType, cuisineType []CuisineType, starRating StarRating, menu Menu, address Address) *Restaurant {
	return &Restaurant{
		ID:             id,
		Name:           name,
		Description:    description,
		BusinessHours:  businessHours,
		MealType:       mealType,
		CuisineType:    cuisineType,
		StarRating:     starRating,
		Menu:           menu,
		Address:        address,
	}
}

// Getter methods for Restaurant struct
func (r *Restaurant) GetID() int {
	return r.ID
}

func (r *Restaurant) GetName() string {
	return r.Name
}

func (r *Restaurant) GetDescription() string {
	return r.Description
}

func (r *Restaurant) GetBusinessHours() BusinessHours {
	return r.BusinessHours
}

func (r *Restaurant) GetMealType() MealType {
	return r.MealType
}

func (r *Restaurant) GetCuisineType() []CuisineType {
	return r.CuisineType
}

func (r *Restaurant) GetStarRating() StarRating {
	return r.StarRating
}

func (r *Restaurant) GetMenu() Menu {
	return r.Menu
}

func (r *Restaurant) GetAddress() Address {
	return r.Address
}

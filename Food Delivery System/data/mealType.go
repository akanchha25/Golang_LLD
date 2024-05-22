package data

// MealType represents the type of meal (vegetarian or non-vegetarian)
type MealType int


// Constants representing the MealType enum
const (
	MealTypeUnset MealType = iota
	VEG 
	NON_VEG
)
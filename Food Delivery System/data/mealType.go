package data

// MealType represents the type of meal (vegetarian or non-vegetarian)
type MealType int


// Constants representing the MealType enum
const (
	VEG MealType = iota
	NON_VEG
)
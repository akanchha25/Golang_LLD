package manager

import (
	"errors"
	"food_delivery_system/data"
)

type UserManager struct{}

func NewUserManager() *UserManager {
	return &UserManager{}
}

// GetUserById retrieves a user by their ID.
func (um *UserManager) GetUserById(id int) (data.User, error) {
	// Implement the logic to get the user by ID
	// Return a dummy user for demonstration; replace with actual implementation
	return data.User{}, errors.New("user not found")
}

// GetUserByToken retrieves a user by their token.
func (um *UserManager) GetUserByToken(token string) (data.User, error) {
	// Implement the logic to get the user by token
	// Return a dummy user for demonstration; replace with actual implementation
	return data.User{}, errors.New("user not found")
}

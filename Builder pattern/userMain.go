package main

import (
	userBuilder "builder_pattern/userBuilder"
	"fmt"
)

func main() {
	// Create a user using the builder pattern
	user := userBuilder.NewUserBuilder(1, "John Doe").
		SetPhoneNumber("1234567890").
		SetAge(30).
		Build()

	// Print user details
	fmt.Println("User Details:")
	fmt.Println("ID:", user.Id)
	fmt.Println("Name:", user.Name)
	fmt.Println("Phone Number:", user.PhoneNumber)
	fmt.Println("Age:", user.Age)
}

package data

type User struct {
	ID          int
	Name        string
	Address     Address
	PhoneNumber string
	EmailID     string
}

// NewUser is a constructor function to create a new User instance
func NewUser(id int, name string, address Address, phoneNumber string, emailID string) *User {
	return &User{
		ID:          id,
		Name:        name,
		Address:     address,
		PhoneNumber: phoneNumber,
		EmailID:     emailID,
	}
}

// GetID returns the ID of the user
func (u *User) GetID() int {
	return u.ID
}

// GetName returns the name of the user
func (u *User) GetName() string {
	return u.Name
}

// GetAddress returns the address of the user
func (u *User) GetAddress() Address {
	return u.Address
}

// GetPhoneNumber returns the phone number of the user
func (u *User) GetPhoneNumber() string {
	return u.PhoneNumber
}

// GetEmailID returns the email ID of the user
func (u *User) GetEmailID() string {
	return u.EmailID
}

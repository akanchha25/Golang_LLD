package userBuilder

// User represents a user object
type User struct {
	Id          int64
	Name        string
	PhoneNumber string // optional
	Age         int32  // optional
}

// UserBuilder constructs User objects using the builder pattern
type UserBuilder struct {
	id          int64
	name        string
	phoneNumber string
	age         int32
}

// NewUserBuilder creates a new instance of UserBuilder
func NewUserBuilder(id int64, name string) *UserBuilder {
	return &UserBuilder{
		id:   id,
		name: name,
	}
}

// SetPhoneNumber sets the phoneNumber of the user
func (b *UserBuilder) SetPhoneNumber(phoneNumber string) *UserBuilder {
	b.phoneNumber = phoneNumber
	return b
}

// SetAge sets the age of the user
func (b *UserBuilder) SetAge(age int32) *UserBuilder {
	b.age = age
	return b
}

// Build constructs and returns the User object
func (b *UserBuilder) Build() *User {
	return &User{
		Id:          b.id,
		Name:        b.name,
		PhoneNumber: b.phoneNumber,
		Age:         b.age,
	}
}

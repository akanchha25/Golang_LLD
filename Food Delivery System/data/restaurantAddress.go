package data

type Address struct {
    AddressLine1 string
    AddressLine2 string
    AddressLine3 string
    City         string
    State        string
    Zip          string
    Country      string
}

// NewAddress creates a new instance of Address.
func NewAddress(addressLine1, addressLine2, addressLine3, city, state, zip, country string) *Address {
    return &Address{
        AddressLine1: addressLine1,
        AddressLine2: addressLine2,
        AddressLine3: addressLine3,
        City:         city,
        State:        state,
        Zip:          zip,
        Country:      country,
    }
}

func (a *Address) GetAddressLine1() string {
    return a.AddressLine1
}

func (a *Address) GetAddressLine2() string {
    return a.AddressLine2
}

func (a *Address) GetAddressLine3() string {
    return a.AddressLine3
}

func (a *Address) GetCity() string {
    return a.City
}

func (a *Address) GetState() string {
    return a.State
}

func (a *Address) GetZip() string {
    return a.Zip
}

func (a *Address) GetCountry() string {
    return a.Country
}

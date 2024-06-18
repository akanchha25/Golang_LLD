package data

import "time"

// VehicleType represents the type of the vehicle (you need to define this).
type VehicleType int

const (
    // Define your vehicle types here.
    Car VehicleType = iota
    Bike
    Truck
)

// Vehicle represents a vehicle with a name, type, number, and entry time.
type Vehicle struct {
    Name       string
    VehicleType VehicleType
    Number     string
    EntryTime  time.Time
}

// NewVehicle creates a new Vehicle with the provided name, type, number, and entry time.
func NewVehicle(name string, vehicleType VehicleType, number string, entryTime time.Time) *Vehicle {
    return &Vehicle{
        Name:       name,
        VehicleType: vehicleType,
        Number:     number,
        EntryTime:  entryTime,
    }
}

// GetName returns the name of the vehicle.
func (v *Vehicle) GetName() string {
    return v.Name
}

// GetVehicleType returns the type of the vehicle.
func (v *Vehicle) GetVehicleType() VehicleType {
    return v.VehicleType
}

// GetNumber returns the number of the vehicle.
func (v *Vehicle) GetNumber() string {
    return v.Number
}

// GetEntryTime returns the entry time of the vehicle.
func (v *Vehicle) GetEntryTime() time.Time {
    return v.EntryTime
}

package manager

import (
    "parking_slot_design/data"
    "fmt"
    "sync"
)

var (
    factory *VehicleTypeManagerFactory
    once    sync.Once
)



// The GetVehicleTypeManagerFactory function returns a singleton instance of VehicleTypeManagerFactory.
type VehicleTypeManagerFactory struct{}

func GetVehicleTypeManagerFactory() *VehicleTypeManagerFactory {
    once.Do(func() {
        factory = &VehicleTypeManagerFactory{}
    })
    return factory
}



// GetVehicleTypeManager returns the appropriate manager based on the vehicle type
func (v *VehicleTypeManagerFactory) GetVehicleTypeManager(vehicleType data.VehicleType) VehicleTypeManager {
	var vehicleTypeManager VehicleTypeManager

	switch vehicleType {
	case data.TWO_WHEELER:
		vehicleTypeManager = &TwoWheelerManager{}
	case data.FOUR_WHEELER:
		vehicleTypeManager = &FourWheelerManager{}
	case data.HEAVY:
		vehicleTypeManager = &HeavyWheelerManager{}
	default:
		fmt.Println("Unsupported vehicle type")
	}

	return vehicleTypeManager
}
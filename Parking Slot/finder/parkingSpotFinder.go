package finder

import (
	"parking_slot_design/data"
	 "parking_slot_design/manager"
	"parking_slot_design/selector"
)

// ParkingSpotFinder is a struct that holds references to VehicleTypeManager and ParkingSpotSelector.
type ParkingSpotFinder struct {
	vehicleTypeManager manager.VehicleTypeManager
	parkingSpotSelector selector.ParkingSpotSelector
}

// NewParkingSpotFinder creates a new instance of ParkingSpotFinder.
func NewParkingSpotFinder(vehicleTypeManager manager.VehicleTypeManager, parkingSpotSelector selector.ParkingSpotSelector) *ParkingSpotFinder {
	return &ParkingSpotFinder{
		vehicleTypeManager: vehicleTypeManager,
		parkingSpotSelector: parkingSpotSelector,
	}
}

// FindParkingSpot retrieves a list of parking spots from the VehicleTypeManager and selects a spot using the ParkingSpotSelector.
// here we are implementing strategy pattern, different strategy for management and selection
func (psf *ParkingSpotFinder) FindParkingSpot() *data.ParkingSpot {
	parkingSpots := psf.vehicleTypeManager.GetParkingSpots()
	return psf.parkingSpotSelector.SelectSpot(parkingSpots)
}
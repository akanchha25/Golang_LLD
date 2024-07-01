package apis

import (
	"parking_slot_design/data"
	"parking_slot_design/vacate"
)

type VacateParkingSpotAPI struct {}

func (v *VacateParkingSpotAPI) VacateParkingSpot(parkingSpot data.ParkingSpot) bool {
	// Create an instance of ParkingSpotVacator and call its method
	vacatorInstance := &vacator.ParkingSpotVacator{}
	isVacated := vacatorInstance.VacateParkingSpot(parkingSpot)
	return isVacated
}
package manager

import (
	"parking_slot_design/data"
)

type FourWheelerManager struct {}

func (f *FourWheelerManager) GetParkingSpots() [] data.ParkingSpot {
	return nil
}

func (f *FourWheelerManager) GetParkingFees(durationInHours float64) float64 {
	return 0
}
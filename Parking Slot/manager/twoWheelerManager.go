package manager

import (
	"parking_slot_design/data"
)

type TwoWheelerManager struct {}

func (t *TwoWheelerManager) GetParkingSpots() []data.ParkingSpot {
	return nil
}

func (t *TwoWheelerManager) GetParkingFees(durationInHours float64) float64 {
	return 0
}
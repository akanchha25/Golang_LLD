package manager

import (
	"parking_slot_design/data"
)

type HeavyWheelerManager struct {}

func (h *HeavyWheelerManager) GetParkingSpots() [] data.ParkingSpot {
	return nil
}

func (h *HeavyWheelerManager) GetParkingFees(durationInHours float64) float64 {
	return 0
}
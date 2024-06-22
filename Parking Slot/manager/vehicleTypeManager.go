package manager

import (
	"parking_slot_design/data"
)

type VehicleTypeManager interface {
	GetParkingSpots() []data.ParkingSpot
    GetParkingFees(durationInHours float64) float64
}
package selector

import (
	"parking_slot_design/data"
)

type ParkingSpotSelector interface {

	SelectSpot(parkingSpots []data.ParkingSpot) *data.ParkingSpot
}
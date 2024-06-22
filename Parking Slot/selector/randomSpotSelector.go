package selector

import (
	"parking_slot_design/data"
	"math/rand"
)

type RandomSelector struct{}

func (rs *RandomSelector) SelectSpot(parkingSpots []data.ParkingSpot) *data.ParkingSpot {
	if len(parkingSpots) == 0 {
		return nil
	}
	randomIndex := rand.Intn(len(parkingSpots))
	return &parkingSpots[randomIndex]
}


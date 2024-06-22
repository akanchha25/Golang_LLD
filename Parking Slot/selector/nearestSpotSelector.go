package selector

import (
	"parking_slot_design/data"
)

type NearestSelector struct {
	entryPoint data.EntryPoint
}

func NewNearestSelector(entryPoint data.EntryPoint) *NearestSelector {
	return &NearestSelector{
		entryPoint: entryPoint,
	}
}

func (ns *NearestSelector) SelectSpot(parkingSpots []data.ParkingSpot) *data.ParkingSpot {
	// Implementation of the selection logic will go here.
	return nil
}

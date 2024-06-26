package selector

import (
	"parking_slot_design/data"
)

type SpotSelectorFactory struct{}

func NewSpotSelectorFactory() *SpotSelectorFactory {
	return &SpotSelectorFactory{}
}

func (ssf SpotSelectorFactory) GetNearestParkingSpotSelector(entryPoint data.EntryPoint) ParkingSpotSelector {
	return &NearestSelector{entryPoint: entryPoint}
}

func (ssf SpotSelectorFactory) GetRandomSelector() ParkingSpotSelector {
	return &RandomSelector{}
}

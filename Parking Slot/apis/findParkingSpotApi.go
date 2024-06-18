package apis

import (
	"parking_slot_design/data"
)

type FindParkingSlotAPI struct{}


func NewFindParkingSlotAPI() *FindParkingSlotAPI {
	return &FindParkingSlotAPI{}
}

func (api *FindParkingSlotAPI) FindParkingSlot(entryPoint data.EntryPoint, vehicleType data.Vehicle, spotSelection data.SpotSelection) {

}
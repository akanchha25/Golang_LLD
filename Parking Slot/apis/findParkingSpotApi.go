package apis

import (
	"fmt"
	"parking_slot_design/data"
	 "parking_slot_design/manager"
	// "parking_slot_design/selector"
)

type FindParkingSlotAPI struct{}


func NewFindParkingSlotAPI() *FindParkingSlotAPI {
	return &FindParkingSlotAPI{}
}

func (api *FindParkingSlotAPI) FindParkingSlot(entryPoint data.EntryPoint, vehicleType data.VehicleType, spotSelection data.SpotSelection) {
      
	
	vehicleFactory := manager.GetVehicleTypeManagerFactory()
    var vehicleTypeManager manager.VehicleTypeManager =  vehicleFactory.GetVehicleTypeManager(vehicleType)

	fmt.Println(vehicleTypeManager)

    

	

}
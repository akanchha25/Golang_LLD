package apis

import (
	"fmt"
	"parking_slot_design/data"
	 "parking_slot_design/manager"
	"parking_slot_design/selector"
	"parking_slot_design/finder"
)

type FindParkingSlotAPI struct{}


func NewFindParkingSlotAPI() *FindParkingSlotAPI {
	return &FindParkingSlotAPI{}
}

func (api *FindParkingSlotAPI) FindParkingSlot(entryPoint data.EntryPoint, vehicleType data.VehicleType, spotSelection data.SpotSelection) *data.ParkingSpot {
      
	
	vehicleFactory := manager.GetVehicleTypeManagerFactory()
    var vehicleTypeManager manager.VehicleTypeManager =  vehicleFactory.GetVehicleTypeManager(vehicleType)

	fmt.Println(vehicleTypeManager)
	var parkingSpotSelector selector.ParkingSpotSelector

	// Check spotSelection and handle accordingly
	switch spotSelection {
	case data.RANDOM:
		
		parkingSpotSelector =  selector.SpotSelectorFactory{}.GetRandomSelector()
		// Implement logic for random spot selection
		fmt.Println(parkingSpotSelector)
	case data.NEAREST:
		parkingSpotSelector = selector.SpotSelectorFactory{}.GetNearestParkingSpotSelector(entryPoint)
		// Implement logic for nearest spot selection
		fmt.Println(parkingSpotSelector)
	default:
		fmt.Println("Unknown spot selection strategy")
	}

	psf := finder.NewParkingSpotFinder(vehicleTypeManager, parkingSpotSelector)

	return psf.FindParkingSpot()
}
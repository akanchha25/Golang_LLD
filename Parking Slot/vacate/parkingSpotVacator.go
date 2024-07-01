package vacator

import "parking_slot_design/data"

type ParkingSpotVacator struct{}

func (psv *ParkingSpotVacator) VacateParkingSpot(parkingSpot data.ParkingSpot) bool {
  //logic to free the spot in DATABASE
  return true

}
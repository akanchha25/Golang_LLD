package manager

type TwoWheelerManager struct {}

func (t *TwoWheelerManager) GetParkingSpots() [] ParkingSpot {
	return nil
}

func (t *TwoWheelerManager) GetParkingFees(durationInHours float64) float64 {
	return 0
}
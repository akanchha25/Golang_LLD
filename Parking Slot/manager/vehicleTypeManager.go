package manager

type VehicleTypeManager interface {
	GetParkingSpots() []ParkingSpot
    GetParkingFees(durationInHours float64) float64
}
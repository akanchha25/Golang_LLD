package data



type ParkingSpot struct {
	floorNum   string
	vehicleType VehicleType
	name       string
	isFree     bool
}

func NewParkingSpot(floorNum string, vehicleType VehicleType, name string, isFree bool) *ParkingSpot {
	return &ParkingSpot{
		floorNum:   floorNum,
		vehicleType: vehicleType,
		name:       name,
		isFree:     isFree,
	}
}

func (ps *ParkingSpot) GetFloorNum() string {
	return ps.floorNum
}

func (ps *ParkingSpot) GetVehicleType() VehicleType {
	return ps.vehicleType
}

func (ps *ParkingSpot) GetName() string {
	return ps.name
}

func (ps *ParkingSpot) IsFree() bool {
	return ps.isFree
}

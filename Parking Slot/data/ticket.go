package data

// Ticket represents a parking ticket.
type Ticket struct {
    refNum      string
    vehicle     *Vehicle
    parkingSpot ParkingSpot
}

// NewTicket creates a new Ticket with the given reference number, vehicle, and parking spot.
func NewTicket(refNum string, vehicle *Vehicle, parkingSpot ParkingSpot) *Ticket {
    return &Ticket{
        refNum:      refNum,
        vehicle:     vehicle,
        parkingSpot: parkingSpot,
    }
}

// GetRefNum returns the reference number of the ticket.
func (t *Ticket) GetRefNum() string {
    return t.refNum
}

// GetVehicle returns the vehicle associated with the ticket.
func (t *Ticket) GetVehicle() *Vehicle {
    return t.vehicle
}

// GetParkingSpot returns the parking spot associated with the ticket.
func (t *Ticket) GetParkingSpot() ParkingSpot {
    return t.parkingSpot
}

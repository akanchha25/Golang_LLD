package payment

import (
	"errors"
	"parking_slot_design/data"
	"parking_slot_design/manager"
	"time"
)

// ParkingFeeProcessor is responsible for processing parking fees.
type ParkingFeeProcessor struct{}

// GetParkingFees calculates the parking fees for the given ticket.
func (p *ParkingFeeProcessor) GetParkingFees(ticket *data.Ticket) float64 {
	
	duration := calculateDuration(ticket)
	// logic to figure out the duration can be added here
	vehicleFactory := manager.GetVehicleTypeManagerFactory()
	var vehicleTypeManager manager.VehicleTypeManager = vehicleFactory.GetVehicleTypeManager(ticket.GetVehicle().GetVehicleType())
	return vehicleTypeManager.GetParkingFees(duration)
}

// ProcessParkingFees processes the parking fees for the given ticket and payment processor.
func (p *ParkingFeeProcessor) ProcessParkingFees(ticket *data.Ticket, paymentProcessor PaymentProcessor) (bool, error) {
	fees := p.GetParkingFees(ticket)
	if fees != paymentProcessor.GetAmount() {
		return false, errors.New("payment amount does not match parking fees")
	}
	return paymentProcessor.ExecutePayment(), nil
}

func calculateDuration(ticket *data.Ticket) float64 {
    // Current time
    currentTime := time.Now()
    // Calculate duration in seconds
    duration := currentTime.Sub(ticket.GetVehicle().GetEntryTime()).Seconds()
    return duration
}


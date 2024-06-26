package ticket

import (
	"parking_slot_design/data"
	"fmt"
	"time"
)

// TicketGenerator is responsible for generating parking tickets.
type TicketGenerator struct{}

// GenerateTicket generates a new parking ticket for the given vehicle and parking spot.
func (tg *TicketGenerator) GenerateTicket(vehicle *data.Vehicle, parkingSpot data.ParkingSpot) *data.Ticket {
	ticketNum := tg.getUniqueTicketNum()
	// logic to check if the spot is free, park the vehicle, and persist in DB can be added here
	return data.NewTicket(ticketNum, vehicle, parkingSpot)
}

// getUniqueTicketNum generates a unique ticket number.
func (tg *TicketGenerator) getUniqueTicketNum() string {
	// Generate a unique ticket number. This is a simple example using timestamp.
	return fmt.Sprintf("TICKET-%d", time.Now().UnixNano())
}
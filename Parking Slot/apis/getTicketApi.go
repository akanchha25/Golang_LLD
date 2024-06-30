package apis

import "parking_slot_design/data"

type GetTicketAPI struct{}

func(gta *GetTicketAPI) GetTicket(vehicle data.Vehicle, parkingSpot data.ParkingSpot) data.Ticket {
	var ticket data.Ticket

	return ticket

}
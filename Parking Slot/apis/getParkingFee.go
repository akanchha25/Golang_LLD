package apis

import (
	"parking_slot_design/data"
)

type GetParkingFeeAPI struct {}


func(gpf GetParkingFeeAPI)  GetParkingFee(ticket data.Ticket) float64 {
	var duration float64 = 0

	return duration
}
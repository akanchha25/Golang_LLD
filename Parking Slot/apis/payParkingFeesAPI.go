package apis

import (
	"parking_slot_design/data"
	"parking_slot_design/payment"
	"strconv"
)

type PayParkingFeesAPI struct{}

func (ppf *PayParkingFeesAPI) PayParkingFees(ticket *data.Ticket, paymentMode payment.PaymentMode, paymentDetails map[string]string) (bool, error) {
	//var paymentProcessor payment.PaymentProcessor
	amount, err := strconv.ParseFloat(paymentDetails["AMOUNT"], 64)
	if err != nil {
		return false, err
	}

	ps := payment.NewPaymentSettle(paymentMode, paymentDetails, amount)

	paymentProcessor, err := ps.GetPaymentSettle(paymentMode, paymentDetails, amount)
	if err != nil {
		return false, err
	}

	parkingFeeProcessor := &payment.ParkingFeeProcessor{}
	return parkingFeeProcessor.ProcessParkingFees(ticket, paymentProcessor)
}




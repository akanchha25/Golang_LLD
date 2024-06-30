package apis

import (
	"errors"
	"parking_slot_design/data"
	"parking_slot_design/payment"
	"strconv"
)

type PayParkingFeesAPI struct{}

func (ppf *PayParkingFeesAPI) PayParkingFees(ticket data.Ticket, paymentMode payment.PaymentMode, paymentDetails map[string]string) (bool, error) {
	var paymentProcessor payment.PaymentProcessor
	amount, err := strconv.ParseFloat(paymentDetails["AMOUNT"], 64)
	if err != nil {
		return false, err
	}

	switch paymentMode {
	case payment.CARD:
		cardDetails := payment.CardDetails{
			NameOnCard: paymentDetails["NAME_ON_CARD"],
			Pin:        parsePin(paymentDetails["PIN"]),
			CardNumber: paymentDetails["CARD_NUMBER"],
		}
		paymentProcessor = payment.NewCardPaymentProcessor(amount, cardDetails)
	case payment.CASH:
		paymentProcessor = payment.NewCashPaymentProcessor(amount)
	default:
		return false, errors.New("invalid payment mode")
	}

	return ProcessParkingFees(ticket, paymentProcessor)
}

func parsePin(pinStr string) int {
	pin, _ := strconv.Atoi(pinStr)
	return pin
}



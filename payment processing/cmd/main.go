package main

import (
	"payment-processing/internal/domain"
	"payment-processing/internal/infrastructure/database/postgres"
	"payment-processing/pkg/notification"
	"payment-processing/pkg/payment"
)

func main() {
	// Initialize database connection
	db, err := postgres.Connect()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// Initialize concrete instances of payment and notification modules
	upiPayment := &payment.UPIPayment{}
	smsNotification := &notification.SMSNotification{}

	// Example usage: Process payment and send notification
	amount := 100.0
	message := "Payment processed successfully"
	err = domain.ProcessPaymentAndNotify(upiPayment, smsNotification, db, amount, message)
	if err != nil {
		panic(err)
	}
}

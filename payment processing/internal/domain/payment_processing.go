package domain

import (
	"database/sql"
	"payment-processing/pkg/notification"
	"payment-processing/pkg/payment"
)

func ProcessPaymentAndNotify(p payment.Payment, n notification.Notification, db *sql.DB, amount float64, message string) error {
	// Process payment
	if err := p.ProcessPayment(amount); err != nil {
		return err
	}

	// Send notification
	if err := n.SendNotification(message); err != nil {
		return err
	}

	// Save payment data to database (optional)
	// Implement logic to save payment details to PostgreSQL database
	// For simplicity, we're just printing the data here
	// fmt.Printf("Saving payment data to database: Amount=%.2f, Message=%s\n", amount, message)

	return nil
}

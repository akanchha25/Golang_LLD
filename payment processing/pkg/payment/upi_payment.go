package payment

import "fmt"

type UPIPayment struct{}

func (u *UPIPayment) ProcessPayment(amount float64) error {
	fmt.Printf("Processing UPI payment for amount %.2f\n", amount)
	return nil
}

package notification

import "fmt"

type SMSNotification struct{}

func (s *SMSNotification) SendNotification(message string) error {
	fmt.Println("Sending SMS notification:", message)
	return nil
}

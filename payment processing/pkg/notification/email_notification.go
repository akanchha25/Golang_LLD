package notification

import "fmt"

type EmailNotification struct{}

func (e *EmailNotification) SendNotification(message string) error {
	fmt.Println("Sending SMS notification:", message)
	return nil
}

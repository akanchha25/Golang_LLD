package notification

type Notification interface {
	SendNotification(message string) error
}

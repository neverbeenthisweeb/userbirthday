package notification

import "context"

const (
	NotificationTypeWA    = "WhatsApp"
	NotificationTypeEmail = "Email"

	DefaultNotificationSubject = "Happy Birthday!"
	DefaultNotificationBody    = "Hi {{.username}}, Here is a promo {{.promocode}} for you"
)

type Notification interface {
	Send(ctx context.Context, req NotificationRequest) error
}

type NotificationRequest struct {
	NotificationType string
	Subject          string
	Body             string
	Target           string
	TemplateData     map[string]string
}

func (nr NotificationRequest) Message() string {
	return nr.Subject + " --- " + nr.Body
}

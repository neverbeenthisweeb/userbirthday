package notification

import "context"

const (
	NotificationTypeWA    = "WhatsApp"
	NotificationTypeEmail = "Email"
)

// FIXME: Notification store can be in memory. Remove from ERD
type Notification interface {
	Send(ctx context.Context, req NotificationRequest) error
}

// FIXME: Add validation
type NotificationRequest struct {
	NotificationType string
	Target           string
	TemplateID       string
	TemplateData     map[string]string
}

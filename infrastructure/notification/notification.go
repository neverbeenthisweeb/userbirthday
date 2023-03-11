package notification

import "context"

const (
	NotificationTypeWA    = "WhatsApp"
	NotificationTypeEmail = "Email"

	// NotifTemplateWA    = "(from:sayakaya.wa) HBD %s! Here is a voucher for you %s"
	// NotifTemplateEmail = "(from:sayakaya.email) HBD %s! Here is a voucher for you %s"
)

// FIXME: Notification store can be in memory. Remove from ERD
type Notification interface {
	Send(ctx context.Context, req NotificationRequest) error
}

type NotificationRequest struct {
	NotificationType string
	Target           string
	TemplateID       string
	TemplateData     map[string]string
}

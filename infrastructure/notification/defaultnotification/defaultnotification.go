package defaultnotification

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"text/template"
	"userbirthday/common"
	"userbirthday/infrastructure/notification"
)

var (
	ErrInvalidNotificationType = errors.New("invalid notification type")
)

type DefaultNotification struct{}

func NewDefaulNotification() *DefaultNotification {
	return &DefaultNotification{}
}

func (dn *DefaultNotification) Send(ctx context.Context, req notification.NotificationRequest) error {
	tmpl, err := template.New("notification").Parse(req.Message())
	if err != nil {
		common.LogErr(ctx, "Failed to parse message", err)
		return err
	}

	bTmpl := new(strings.Builder)

	err = tmpl.Execute(bTmpl, req.TemplateData)
	if err != nil {
		common.LogErr(ctx, "Failed to execute", err)
		return err
	}

	switch req.NotificationType {
	case notification.NotificationTypeEmail:
		fmt.Printf("(EMAIL) Message sent: %s\n", bTmpl)
	case notification.NotificationTypeWA:
		fmt.Printf("(WA) Message is sent: %s\n", bTmpl)
	default:
		err := ErrInvalidNotificationType
		common.LogErr(ctx, fmt.Sprintf("Notification type %s is invalid", req.NotificationType), err)
		return err
	}

	return nil
}

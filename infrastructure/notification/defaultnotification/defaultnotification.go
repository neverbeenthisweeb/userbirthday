package defaultnotification

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"text/template"
	"userbirthday/common"
	"userbirthday/infrastructure"
	"userbirthday/infrastructure/notification"
	"userbirthday/infrastructure/repository"
)

var (
	ErrInvalidNotificationType = errors.New("invalid notification type")
)

type DefaultNotification struct {
	repoNotificationTemplate repository.NotificationTemplateRepository
}

func NewDefaulNotification(infra *infrastructure.Infrastructure) *DefaultNotification {
	return &DefaultNotification{
		repoNotificationTemplate: infra.RepoNotificationTemplate(),
	}
}

func (dn *DefaultNotification) Send(ctx context.Context, req notification.NotificationRequest) error {
	text, err := dn.repoNotificationTemplate.GetNotificationTemplate(ctx, req.TemplateID)
	if err != nil {
		common.LogErr(ctx, "Failed to get notification template", err)
		return err
	}

	tmpl, err := template.New("notification").Parse(text)
	if err != nil {
		common.LogErr(ctx, "Failed to parse text", err)
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
		fmt.Printf("(EMAIL) Message is sent: %s\n", bTmpl)
	case notification.NotificationTypeWA:
		fmt.Printf("(WA) Message is sent: %s\n", bTmpl)
	default:
		err := ErrInvalidNotificationType
		common.LogErr(ctx, fmt.Sprintf("Notification type %s is invalid", req.NotificationType), err)
		return err
	}

	return nil
}

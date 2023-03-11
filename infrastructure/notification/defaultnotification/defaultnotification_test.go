package defaultnotification_test

import (
	"context"
	"testing"
	"userbirthday/common"
	"userbirthday/infrastructure/notification"
	"userbirthday/infrastructure/notification/defaultnotification"

	"github.com/stretchr/testify/assert"
)

func TestDefaultNotification_Send(t *testing.T) {
	ntf := defaultnotification.NewDefaulNotification()

	cases := []struct {
		name string
		ctx  context.Context
		req  notification.NotificationRequest
		err  error
	}{
		{
			name: `When notification type is invalid
			Then return error`,
			ctx: common.ContextWithRequestID(),
			req: notification.NotificationRequest{
				NotificationType: "INVALID_NOTIFICATION_TYPE",
			},
			err: defaultnotification.ErrInvalidNotificationType,
		},
		{
			name: `When email notification is sent
			Then return no error`,
			ctx: common.ContextWithRequestID(),
			req: notification.NotificationRequest{
				NotificationType: notification.NotificationTypeEmail,
				Subject:          notification.DefaultNotificationSubject,
				Body:             notification.DefaultNotificationBody,
				TemplateData: map[string]string{
					"username":  "User Name",
					"promocode": "HBDUSERNAME2023",
				},
			},
			err: nil,
		},
		{
			name: `When WA notification is sent
			Then return no error`,
			ctx: common.ContextWithRequestID(),
			req: notification.NotificationRequest{
				NotificationType: notification.NotificationTypeWA,
				Subject:          notification.DefaultNotificationSubject,
				Body:             notification.DefaultNotificationBody,
				TemplateData: map[string]string{
					"username":  "User Name",
					"promocode": "HBDUSERNAME2023",
				},
			},
			err: nil,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			err := ntf.Send(tc.ctx, tc.req)
			assertError(t, tc.err, err)
		})
	}
}

func assertError(t *testing.T, wantErr, actualErr error) {
	if wantErr == nil {
		assert.Nil(t, actualErr)
		return
	}

	assert.EqualError(t, actualErr, wantErr.Error())
}

package defaultnotification_test

import (
	"context"
	"errors"
	"testing"
	"userbirthday/common"
	"userbirthday/infrastructure"
	"userbirthday/infrastructure/notification"
	"userbirthday/infrastructure/notification/defaultnotification"
	mocksRepository "userbirthday/mocks/infrastructure/repository"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestDefaultNotification_Send(t *testing.T) {
	notifTemplateRepo := &mocksRepository.NotificationTemplateRepository{}

	infra := infrastructure.NewInfrastructure()
	infra.SetRepoNotificationTemplate(notifTemplateRepo)

	ntf := defaultnotification.NewDefaulNotification(infra)

	cases := []struct {
		name   string
		ctx    context.Context
		req    notification.NotificationRequest
		err    error
		mockFn func()
	}{
		{
			name: `When failed to get template
			Then return error`,
			ctx: common.ContextWithRequestID(),
			req: notification.NotificationRequest{
				TemplateID: "1001",
			},
			err: errors.New("something went wrong"),
			mockFn: func() {
				notifTemplateRepo.On("GetNotificationTemplate", mock.Anything, "1001").
					Return("", errors.New("something went wrong")).
					Once()
			},
		},
		{
			name: `When notification type is invalid
			Then return error`,
			ctx: common.ContextWithRequestID(),
			req: notification.NotificationRequest{
				NotificationType: "INVALID_NOTIFICATION_TYPE",
				TemplateID:       "1001",
				TemplateData: map[string]string{
					"text": "HELLO WORLD",
				},
			},
			err: defaultnotification.ErrInvalidNotificationType,
			mockFn: func() {
				notifTemplateRepo.On("GetNotificationTemplate", mock.Anything, "1001").
					Return("{{.text}}", nil).
					Once()
			},
		},
		{
			name: `When email notification is sent
			Then return no error`,
			ctx: common.ContextWithRequestID(),
			req: notification.NotificationRequest{
				NotificationType: notification.NotificationTypeEmail,
				TemplateID:       "1001",
				TemplateData: map[string]string{
					"text": "HELLO WORLD",
				},
			},
			err: nil,
			mockFn: func() {
				notifTemplateRepo.On("GetNotificationTemplate", mock.Anything, "1001").
					Return("{{.text}}", nil).
					Once()
			},
		},
		{
			name: `When WA notification is sent
			Then return no error`,
			ctx: common.ContextWithRequestID(),
			req: notification.NotificationRequest{
				NotificationType: notification.NotificationTypeWA,
				TemplateID:       "1001",
				TemplateData: map[string]string{
					"text": "HELLO WORLD",
				},
			},
			err: nil,
			mockFn: func() {
				notifTemplateRepo.On("GetNotificationTemplate", mock.Anything, "1001").
					Return("{{.text}}", nil).
					Once()
			},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			tc.mockFn()
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

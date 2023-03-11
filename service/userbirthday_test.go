package service_test

import (
	"context"
	"errors"
	"testing"
	"userbirthday/common"
	"userbirthday/infrastructure"
	"userbirthday/infrastructure/notification"
	mocksNotification "userbirthday/mocks/infrastructure/notification"
	mocksRepository "userbirthday/mocks/infrastructure/repository"
	"userbirthday/model"
	"userbirthday/service"

	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestUserBirthday_GiveBirthdayPromo(t *testing.T) {
	userRepo := &mocksRepository.UserRepository{}
	promoRepo := &mocksRepository.PromoRepository{}
	notif := &mocksNotification.Notification{}

	infra := infrastructure.NewInfrastructure()
	infra.SetRepoUser(userRepo)
	infra.SetRepoPromo(promoRepo)
	infra.SetNotification(notif)

	svc := service.NewUserBirthday(infra)

	cases := []struct {
		name   string
		ctx    context.Context
		err    error
		mockFn func()
	}{
		{
			name: `When failed to get verified birthday users
			Then return error`,
			ctx: contextWithRequestID(),
			err: errors.New("something went wrong"),
			mockFn: func() {
				userRepo.On("GetVerifiedBirthdayUsers", mock.Anything).
					Return([]model.User{}, errors.New("something went wrong")).
					Once()
			},
		},
		{
			name: `When user already has birthday promo
			Then don't give birthday promo`,
			ctx: contextWithRequestID(),
			err: nil,
			mockFn: func() {
				userRepo.On("GetVerifiedBirthdayUsers", mock.Anything).
					Return([]model.User{
						{
							ID: "1001",
							Promos: []model.Promo{
								{
									ID:   "2001",
									Type: model.PromoTypeBirthday,
								},
							},
						},
					}, nil).
					Once()
			},
		},
		{
			name: `When failed to create birthday promo
			Then return error`,
			ctx: contextWithRequestID(),
			err: errors.New("something went wrong"),
			mockFn: func() {
				userRepo.On("GetVerifiedBirthdayUsers", mock.Anything).
					Return([]model.User{
						{
							ID:     "1001",
							Name:   "User Name",
							Promos: []model.Promo{},
						},
					}, nil).
					Once()
				promoRepo.On("Create", mock.Anything, fakePromo(nil)).
					Return(model.Promo{}, errors.New("something went wrong")).
					Once()
			},
		},
		{
			name: `When failed to set birthday promo to user
			Then return error`,
			ctx: contextWithRequestID(),
			err: errors.New("something went wrong"),
			mockFn: func() {
				userRepo.On("GetVerifiedBirthdayUsers", mock.Anything).
					Return([]model.User{
						{
							ID:     "1001",
							Name:   "User Name",
							Promos: []model.Promo{},
						},
					}, nil).
					Once()
				promoRepo.On("Create", mock.Anything, fakePromo(nil)).
					Return(fakePromo(func(m model.Promo) model.Promo {
						m.ID = "2001"
						return m
					}), nil).
					Once()
				userRepo.On("SetPromo", mock.Anything, "1001", "2001").
					Return(errors.New("something went wrong")).
					Once()
			},
		},
		{
			name: `When failed to send notification
			Then return no error`,
			ctx: contextWithRequestID(),
			err: nil,
			mockFn: func() {
				userRepo.On("GetVerifiedBirthdayUsers", mock.Anything).
					Return([]model.User{
						{
							ID:     "1001",
							Name:   "User Name",
							Email:  "user.name@email.com",
							Phone:  "+6201",
							Promos: []model.Promo{},
						},
					}, nil).
					Once()
				promoRepo.On("Create", mock.Anything, fakePromo(nil)).
					Return(fakePromo(func(m model.Promo) model.Promo {
						m.ID = "2001"
						return m
					}), nil).
					Once()
				userRepo.On("SetPromo", mock.Anything, "1001", "2001").
					Return(nil).
					Once()
				notif.On("Send", mock.Anything, notification.NotificationRequest{
					NotificationType: notification.NotificationTypeEmail,
					Target:           "user.name@email.com",
					TemplateID:       "email.birthday",
					TemplateData: map[string]string{
						"username":  "User Name",
						"promocode": fakePromo(nil).Code,
					},
				}).
					Return(errors.New("something went wrong")).
					Once()
				notif.On("Send", mock.Anything, notification.NotificationRequest{
					NotificationType: notification.NotificationTypeWA,
					Target:           "+6201",
					TemplateID:       "wa.birthday",
					TemplateData: map[string]string{
						"username":  "User Name",
						"promocode": fakePromo(nil).Code,
					},
				}).
					Return(errors.New("something went wrong")).
					Once()
			},
		},
		{
			name: `When 2 users need to be given birthday promo BUT 1st user already has birthday promo 
			Then continue to give the 2nd user AND return no error`,
			ctx: contextWithRequestID(),
			err: nil,
			mockFn: func() {
				userRepo.On("GetVerifiedBirthdayUsers", mock.Anything).
					Return([]model.User{
						{
							ID:    "1001",
							Name:  "User Name 1",
							Email: "user.name1@email.com",
							Phone: "+6201",
							Promos: []model.Promo{
								{
									ID:   "2001",
									Type: model.PromoTypeBirthday,
								},
							},
						},
						{
							ID:     "1002",
							Name:   "User Name 2",
							Email:  "user.name2@email.com",
							Phone:  "+6202",
							Promos: []model.Promo{},
						},
					}, nil).
					Once()
				promoRepo.On("Create", mock.Anything, fakePromo(func(m model.Promo) model.Promo {
					return model.NewBirthdayPromo("User Name 2")
				})).
					Return(fakePromo(func(m model.Promo) model.Promo {
						m.ID = "2002"
						return m
					}), nil).
					Once()
				userRepo.On("SetPromo", mock.Anything, "1002", "2002").
					Return(nil).
					Once()
				notif.On("Send", mock.Anything, notification.NotificationRequest{
					NotificationType: notification.NotificationTypeEmail,
					Target:           "user.name2@email.com",
					TemplateID:       "email.birthday",
					TemplateData: map[string]string{
						"username":  "User Name 2",
						"promocode": fakePromo(nil).Code,
					},
				}).
					Return(nil).
					Once()
				notif.On("Send", mock.Anything, notification.NotificationRequest{
					NotificationType: notification.NotificationTypeWA,
					Target:           "+6202",
					TemplateID:       "wa.birthday",
					TemplateData: map[string]string{
						"username":  "User Name 2",
						"promocode": fakePromo(nil).Code,
					},
				}).
					Return(nil).
					Once()
			},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			tc.mockFn()
			err := svc.GiveBirthdayPromo(tc.ctx)
			assertError(t, tc.err, err)
		})
	}
}

func contextWithRequestID() context.Context {
	return context.WithValue(context.Background(), common.CtxKeyRequestID, uuid.NewV4().String())
}

func assertError(t *testing.T, wantErr, actualErr error) {
	if wantErr == nil {
		assert.Nil(t, actualErr)
		return
	}

	assert.EqualError(t, actualErr, wantErr.Error())
}

func fakePromo(cb func(m model.Promo) model.Promo) model.Promo {
	ret := model.NewBirthdayPromo("User Name")

	if cb != nil {
		return cb(ret)
	}

	return ret
}

package service_test

import (
	"context"
	"errors"
	"testing"
	"userbirthday/common"
	"userbirthday/infrastructure"
	mocksRepository "userbirthday/mocks/infrastructure/repository"
	"userbirthday/model"
	"userbirthday/service"

	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestUserBirthday_GiveBirthdayPromo(t *testing.T) {
	userRepo := &mocksRepository.UserRepository{}

	infra := infrastructure.NewInfrastructure()
	infra.SetRepoUser(userRepo)

	svc := service.NewUserBirthday(infra)

	cases := []struct {
		name   string
		ctx    context.Context
		err    error
		mockFn func()
	}{
		{
			name: "Failed to get verified birthday users",
			ctx:  contextWithRequestID(),
			err:  errors.New("something went wrong"),
			mockFn: func() {
				userRepo.On("GetVerifiedBirthdayUsers", mock.Anything).
					Return([]model.User{}, errors.New("something went wrong")).
					Once()
			},
		},
		{
			name: "Skip user that already have birthday promo",
			ctx:  contextWithRequestID(),
			err:  nil,
			mockFn: func() {
				userRepo.On("GetVerifiedBirthdayUsers", mock.Anything).
					Return([]model.User{
						{
							ID: "1",
							Promos: []model.Promo{
								{
									ID:   "1",
									Type: model.PromoTypeBirthday,
								},
							},
						},
					}, nil).
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

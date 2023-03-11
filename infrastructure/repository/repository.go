package repository

import (
	"context"
	"userbirthday/model"
)

type UserRepository interface {
	GetVerifiedBirthdayUsers(ctx context.Context) ([]model.User, error)
	SetPromo(ctx context.CancelFunc, userID, promoCode string) (model.User, error)
}

type PromoRepository interface {
	Create(ctx context.Context, m model.Promo) (model.Promo, error)
}

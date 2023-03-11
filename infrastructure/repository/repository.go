package repository

import (
	"context"
	"userbirthday/model"
)

type UserRepository interface {
	GetVerifiedBirthdayUsers(ctx context.Context) ([]model.User, error)
	SetPromo(ctx context.Context, userID, promoCode string) error
}

type PromoRepository interface {
	Create(ctx context.Context, m model.Promo) (model.Promo, error)
}

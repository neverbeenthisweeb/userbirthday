package repository

import (
	"context"
	"userbirthday/model"
)

type UserRepository interface {
	GetVerifiedBirthdayUsers(ctx context.Context) ([]model.User, error)
	SetUserPromo(ctx context.Context, userID, promoCode string) error
}

type PromoRepository interface {
	CreatePromo(ctx context.Context, m model.Promo) (model.Promo, error)
}

package repository

import (
	"context"
	"userbirthday/model"
)

type UserRepository interface {
	GetVerifiedBirthdayUsers(ctx context.Context) ([]model.User, error)
	UpdateUserPromo(ctx context.Context, userID, promoID string) error
}

type PromoRepository interface {
	CreatePromo(ctx context.Context, m model.Promo) (model.Promo, error)
}

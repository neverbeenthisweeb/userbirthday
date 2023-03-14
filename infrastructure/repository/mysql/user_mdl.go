package mysql

import (
	"database/sql"
	"time"
	"userbirthday/model"
)

type User struct {
	ID         string    `db:"id"`
	Name       string    `db:"name"`
	Email      string    `db:"email"`
	Phone      string    `db:"phone"`
	IsVerified bool      `db:"is_verified"`
	Birthdate  time.Time `db:"birthdate"`
	CreatedAt  time.Time `db:"created_at"`
	UpdatedAt  time.Time `db:"updated_at"`
	Promos     []Promotion
}

func (u User) ToModel() model.User {
	ret := model.User{
		ID:         u.ID,
		Name:       u.Name,
		Email:      u.Email,
		Phone:      u.Phone,
		IsVerified: u.IsVerified,
		Birthdate:  u.Birthdate,
	}

	ps := []model.Promo{}
	for _, p := range u.Promos {
		ps = append(ps, p.ToModel())
	}

	ret.Promos = ps

	return ret
}

type UserPromo struct {
	ID            string    `db:"id"`
	UserID        string    `db:"user_id"`
	PromoID       string    `db:"promotion_id"`
	PromoUseCount int       `db:"promotion_use_count"`
	CreatedAt     time.Time `db:"created_at"`
	UpdatedAt     time.Time `db:"updated_at"`
}

type dtoGetVerifiedBirthdayUsers struct {
	UserID     string         `db:"user_id"`
	Name       string         `db:"name"`
	Email      string         `db:"email"`
	Phone      string         `db:"phone"`
	IsVerified bool           `db:"is_verified"`
	Birthdate  time.Time      `db:"birthdate"`
	PromoID    sql.NullString `db:"promotion_id"`
	Code       sql.NullString `db:"code"`
	Type       sql.NullString `db:"type"`
	UseCount   sql.NullInt64  `db:"use_count"`
	UseLimit   sql.NullInt64  `db:"use_limit"`
	ValidFrom  sql.NullTime   `db:"valid_from"`
	ValidTo    sql.NullTime   `db:"valid_to"`
}

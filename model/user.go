package model

import "time"

type User struct {
	Promos     []Promo
	ID         string
	Name       string
	Email      string
	Phone      string
	IsVerified bool
	Birthdate  time.Time
}

func (u User) HasBirthdayPromo() bool {
	for _, prm := range u.Promos {
		if prm.Type == PromoTypeBirthday {
			return true
		}
	}

	return false
}

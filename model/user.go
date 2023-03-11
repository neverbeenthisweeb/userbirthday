package model

import "time"

type User struct {
	ID         string
	Name       string
	Email      string
	Phone      string
	IsVerified bool
	Birthdate  time.Time
	Promos     []Promo
}

func (u User) HasBirthdayPromo() bool {
	for _, prm := range u.Promos {
		if prm.Type == PromoTypeBirthday {
			return true
		}
	}

	return false
}

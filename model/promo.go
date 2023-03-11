package model

import "time"

const (
	PromoTypeBirthday = "birthday"
)

type Promo struct {
	ID        string
	Code      string
	Name      string
	Type      string
	UseLimit  int
	ValidFrom time.Time
	ValidTo   time.Time
}

package model

import (
	"fmt"
	"strconv"
	"strings"
	"time"
	"userbirthday/common"
)

const (
	PromoTypeBirthday = "birthday"

	PromoUseLimitBirthday = 1
)

// FIXME: Remove "name" from docs
type Promo struct {
	ID        string
	Code      string
	Type      string
	UseLimit  int
	ValidFrom time.Time
	ValidTo   time.Time
}

func NewBirthdayPromo(userName string) Promo {
	now := time.Now()

	return Promo{
		Code: fmt.Sprintf("HBD%s%s",
			strings.TrimSpace(strings.ToUpper(userName)),
			strconv.Itoa(now.Year()),
		),
		Type:      PromoTypeBirthday,
		UseLimit:  PromoUseLimitBirthday,
		ValidFrom: common.GetBeginningOfToday(now),
		ValidTo:   common.GetBeginningOfTomorrow(now),
	}
}

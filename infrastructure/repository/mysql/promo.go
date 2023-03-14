package mysql

import (
	"context"
	"time"
	"userbirthday/common"
	"userbirthday/model"

	"github.com/jmoiron/sqlx"
	uuid "github.com/satori/go.uuid"
)

type PromoRepository struct {
	db *sqlx.DB
}

func NewPromoRepository(db *sqlx.DB) *PromoRepository {
	return &PromoRepository{db}
}

func (pr *PromoRepository) CreatePromo(ctx context.Context, m model.Promo) (model.Promo, error) {
	m.ID = uuid.NewV4().String()

	_, err := pr.db.NamedExec(`INSERT INTO promotions (id, code, type, use_limit, valid_from, valid_to) 
	VALUES (:id, :code, :type, :use_limit, :valid_from, :valid_to)`, Promotion{}.FromModel(m))
	if err != nil {
		common.LogErr(ctx, "Failed scan struct", err)
		return model.Promo{}, err
	}

	tempPr := Promotion{}
	err = pr.db.Get(&tempPr, `SELECT * FROM promotions WHERE id=?`, m.ID)
	if err != nil {
		common.LogErr(ctx, "Failed to get promotion", err)
		return model.Promo{}, err
	}

	return tempPr.ToModel(), nil
}

type Promotion struct {
	ID        string    `db:"id"`
	Code      string    `db:"code"`
	Type      string    `db:"type"`
	Amount    int       `db:"amount"`
	UseLimit  int       `db:"use_limit"`
	ValidFrom time.Time `db:"valid_from"`
	ValidTo   time.Time `db:"valid_to"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

func (p Promotion) ToModel() model.Promo {
	return model.Promo{
		ID:        p.ID,
		Code:      p.Code,
		Type:      p.Type,
		Amount:    p.Amount,
		UseCount:  p.UseLimit,
		ValidFrom: p.ValidFrom,
		ValidTo:   p.ValidTo,
	}
}

func (p Promotion) FromModel(m model.Promo) Promotion {
	return Promotion{
		ID:        m.ID,
		Code:      m.Code,
		Type:      m.Type,
		Amount:    m.Amount,
		UseLimit:  m.UseLimit,
		ValidFrom: m.ValidFrom,
		ValidTo:   m.ValidTo,
	}
}

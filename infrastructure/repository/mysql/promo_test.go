package mysql_test

import (
	"context"
	"testing"
	"time"
	"userbirthday/infrastructure/repository/mysql"
	"userbirthday/model"

	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/require"
)

func TestPromoRepository_CreatePromo(t *testing.T) {
	repo := mysql.NewPromoRepository(testDb)

	cases := []struct {
		name      string
		ctx       context.Context
		toCreate  model.Promo
		result    model.Promo
		err       error
		prepareFn func(t *testing.T) // Prepare func for something like seed, cleaning, etc
		assertFn  func(t *testing.T, p model.Promo)
	}{
		{
			name: "Create promo",
			ctx:  context.Background(),
			toCreate: model.Promo{
				Code:      "code",
				Type:      model.PromoTypeBirthday,
				UseLimit:  model.PromoUseLimitBirthday,
				ValidFrom: time.Now(),
				ValidTo:   time.Now(),
			},
			result: model.Promo{},
			err:    nil,
			prepareFn: func(t *testing.T) {
				err := truncateAll(testDb)
				require.NoError(t, err)
			},
			assertFn: func(t *testing.T, p model.Promo) {
				require.NotNil(t, p.ID)
			},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			tc.prepareFn(t)
			pr, err := repo.CreatePromo(tc.ctx, tc.toCreate)
			require.Equal(t, tc.err, err)
			tc.assertFn(t, pr)
		})
	}
}

func insertManyPromotions(db *sqlx.DB, mm []mysql.Promotion) error {
	_, err := db.NamedExec(`INSERT INTO promotions (id, code, type, use_limit, valid_from, valid_to) 
		VALUES (:id, :code, :type, :use_limit, :valid_from, :valid_to)`, mm)
	return err
}

func getPromotion(db *sqlx.DB, promoID string) (mysql.Promotion, error) {
	pr := mysql.Promotion{}
	err := db.Get(&pr, `SELECT * FROM promotions WHERE promotion_id=?`, promoID)
	return pr, err
}

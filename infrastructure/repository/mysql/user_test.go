package mysql_test

import (
	"context"
	"testing"
	"time"
	"userbirthday/common"
	"userbirthday/infrastructure/repository/mysql"
	"userbirthday/model"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/require"
)

func TestUserRepository_GetVerifiedBirthdayUsers(t *testing.T) {
	repo := mysql.NewUserRepository(testDb)

	cases := []struct {
		name      string
		ctx       context.Context
		users     []model.User
		err       error
		prepareFn func(t *testing.T) // Prepare func for something like seed, cleaning, etc
		assertFn  func(t *testing.T, uu []model.User)
	}{
		{
			name:  `Return verified birthday users`,
			ctx:   common.ContextWithRequestID(),
			users: []model.User{},
			err:   nil,
			prepareFn: func(t *testing.T) {
				err := truncateAll(testDb)
				require.NoError(t, err)

				err = insertManyUsers(testDb, []mysql.User{
					{
						ID:         "101",
						Name:       "name 1",
						Email:      "email 1",
						Phone:      "phone 1",
						IsVerified: true,
						Birthdate:  time.Now(),
					},
				})
				require.NoError(t, err)

				err = insertManyPromotions(testDb, []mysql.Promotion{
					{
						ID:        "201",
						Code:      "code 1",
						Type:      model.PromoTypeBirthday,
						UseLimit:  model.PromoUseLimitBirthday,
						ValidFrom: time.Now(),
						ValidTo:   time.Now(),
					},
				})
				require.NoError(t, err)

				err = insertManyUserPromotions(testDb, []mysql.UserPromo{
					{
						ID:            "301",
						UserID:        "101",
						PromoID:       "201",
						PromoUseCount: 0,
					},
				})
				require.NoError(t, err)
			},
			assertFn: func(t *testing.T, uu []model.User) {
				require.Len(t, uu, 1)
			},
		},
		{
			name:  `Return verified birthday users in different years`,
			ctx:   common.ContextWithRequestID(),
			users: []model.User{},
			err:   nil,
			prepareFn: func(t *testing.T) {
				err := truncateAll(testDb)
				require.NoError(t, err)

				err = insertManyUsers(testDb, []mysql.User{
					{
						ID:         "101",
						Name:       "name 1",
						Email:      "email 1",
						Phone:      "phone 1",
						IsVerified: true,
						Birthdate:  time.Now().AddDate(-1, 0, 0),
					},
					{
						ID:         "102",
						Name:       "name 2",
						Email:      "email 2",
						Phone:      "phone 2",
						IsVerified: true,
						Birthdate:  time.Now().AddDate(-2, 0, 0),
					},
				})
				require.NoError(t, err)

				err = insertManyPromotions(testDb, []mysql.Promotion{
					{
						ID:        "201",
						Code:      "code 1",
						Type:      model.PromoTypeBirthday,
						UseLimit:  model.PromoUseLimitBirthday,
						ValidFrom: time.Now(),
						ValidTo:   time.Now(),
					},
					{
						ID:        "202",
						Code:      "code 2",
						Type:      model.PromoTypeBirthday,
						UseLimit:  model.PromoUseLimitBirthday,
						ValidFrom: time.Now(),
						ValidTo:   time.Now(),
					},
				})
				require.NoError(t, err)

				err = insertManyUserPromotions(testDb, []mysql.UserPromo{
					{
						ID:            "301",
						UserID:        "101",
						PromoID:       "201",
						PromoUseCount: 0,
					},
					{
						ID:            "302",
						UserID:        "102",
						PromoID:       "202",
						PromoUseCount: 0,
					},
				})
				require.NoError(t, err)
			},
			assertFn: func(t *testing.T, uu []model.User) {
				require.Len(t, uu, 2)
			},
		},
		{
			name:  `Don't return unverified birthday users`,
			ctx:   common.ContextWithRequestID(),
			users: []model.User{},
			err:   nil,
			prepareFn: func(t *testing.T) {
				err := truncateAll(testDb)
				require.NoError(t, err)

				err = insertManyUsers(testDb, []mysql.User{
					{
						ID:         "101",
						Name:       "name 1",
						Email:      "email 1",
						Phone:      "phone 1",
						IsVerified: false,
						Birthdate:  time.Now(),
					},
				})
				require.NoError(t, err)

				err = insertManyPromotions(testDb, []mysql.Promotion{
					{
						ID:        "201",
						Code:      "code 1",
						Type:      model.PromoTypeBirthday,
						UseLimit:  model.PromoUseLimitBirthday,
						ValidFrom: time.Now(),
						ValidTo:   time.Now(),
					},
				})
				require.NoError(t, err)

				err = insertManyUserPromotions(testDb, []mysql.UserPromo{
					{
						ID:            "301",
						UserID:        "101",
						PromoID:       "201",
						PromoUseCount: 0,
					},
				})
				require.NoError(t, err)
			},
			assertFn: func(t *testing.T, uu []model.User) {
				require.Empty(t, uu)
			},
		},
		{
			name:  `Don't return verified not-having-birthday users`,
			ctx:   common.ContextWithRequestID(),
			users: []model.User{},
			err:   nil,
			prepareFn: func(t *testing.T) {
				err := truncateAll(testDb)
				require.NoError(t, err)

				err = insertManyUsers(testDb, []mysql.User{
					{
						ID:         "101",
						Name:       "name 1",
						Email:      "email 1",
						Phone:      "phone 1",
						IsVerified: true,
						Birthdate:  time.Now().AddDate(0, 0, 1),
					},
				})
				require.NoError(t, err)

				err = insertManyPromotions(testDb, []mysql.Promotion{
					{
						ID:        "201",
						Code:      "code 1",
						Type:      model.PromoTypeBirthday,
						UseLimit:  model.PromoUseLimitBirthday,
						ValidFrom: time.Now(),
						ValidTo:   time.Now(),
					},
				})
				require.NoError(t, err)

				err = insertManyUserPromotions(testDb, []mysql.UserPromo{
					{
						ID:            "301",
						UserID:        "101",
						PromoID:       "201",
						PromoUseCount: 0,
					},
				})
				require.NoError(t, err)
			},
			assertFn: func(t *testing.T, uu []model.User) {
				require.Empty(t, uu)
			},
		},
		{
			name:  `Return all user's promotions`,
			ctx:   common.ContextWithRequestID(),
			users: []model.User{},
			err:   nil,
			prepareFn: func(t *testing.T) {
				err := truncateAll(testDb)
				require.NoError(t, err)

				err = insertManyUsers(testDb, []mysql.User{
					{
						ID:         "101",
						Name:       "name 1",
						Email:      "email 1",
						Phone:      "phone 1",
						IsVerified: true,
						Birthdate:  time.Now(),
					},
				})
				require.NoError(t, err)

				err = insertManyPromotions(testDb, []mysql.Promotion{
					{
						ID:        "201",
						Code:      "code 1",
						Type:      model.PromoTypeBirthday,
						UseLimit:  model.PromoUseLimitBirthday,
						ValidFrom: time.Now(),
						ValidTo:   time.Now(),
					},
					{
						ID:        "202",
						Code:      "code 2",
						Type:      model.PromoTypeBirthday,
						UseLimit:  model.PromoUseLimitBirthday,
						ValidFrom: time.Now(),
						ValidTo:   time.Now(),
					},
					{
						ID:        "203",
						Code:      "code 3",
						Type:      model.PromoTypeBirthday,
						UseLimit:  model.PromoUseLimitBirthday,
						ValidFrom: time.Now(),
						ValidTo:   time.Now(),
					},
				})
				require.NoError(t, err)

				err = insertManyUserPromotions(testDb, []mysql.UserPromo{
					{
						ID:            "301",
						UserID:        "101",
						PromoID:       "201",
						PromoUseCount: 0,
					},
					{
						ID:            "302",
						UserID:        "101",
						PromoID:       "202",
						PromoUseCount: 0,
					},
					{
						ID:            "303",
						UserID:        "101",
						PromoID:       "203",
						PromoUseCount: 0,
					},
				})
				require.NoError(t, err)
			},
			assertFn: func(t *testing.T, uu []model.User) {
				require.Len(t, uu, 1)
				require.Len(t, uu[0].Promos, 3)
			},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			tc.prepareFn(t)
			users, err := repo.GetVerifiedBirthdayUsers(tc.ctx)
			require.Equal(t, tc.err, err)
			tc.assertFn(t, users)
		})
	}
}

func TestUserRepository_UpdateUserPromo(t *testing.T) {
	repo := mysql.NewUserRepository(testDb)

	cases := []struct {
		name      string
		ctx       context.Context
		userID    string
		promoID   string
		err       error
		prepareFn func(t *testing.T) // Prepare func for something like seed, cleaning, etc
		assertFn  func(t *testing.T, userID, promoID string)
	}{
		{
			name:    "Create user promo",
			ctx:     context.Background(),
			userID:  "101",
			promoID: "201",
			err:     nil,
			prepareFn: func(t *testing.T) {
				err := truncateAll(testDb)
				require.NoError(t, err)
			},
			assertFn: func(t *testing.T, userID, promoID string) {
				up, err := getUserPromo(testDb, "101", "201")
				require.NoError(t, err)
				require.Equal(t, "101", up.UserID)
				require.Equal(t, "201", up.PromoID)
			},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			tc.prepareFn(t)
			err := repo.UpdateUserPromo(tc.ctx, tc.userID, tc.promoID)
			require.Equal(t, tc.err, err)
			tc.assertFn(t, tc.userID, tc.promoID)
		})
	}
}

func insertManyUsers(db *sqlx.DB, mm []mysql.User) error {
	_, err := db.NamedExec(`INSERT INTO users (id, name, email, phone, is_verified, birthdate) 
		VALUES (:id, :name, :email, :phone, :is_verified, :birthdate)`, mm)
	return err
}

func insertManyUserPromotions(db *sqlx.DB, mm []mysql.UserPromo) error {
	_, err := db.NamedExec(`INSERT INTO users_promotions (id, user_id, promotion_id, promotion_use_count) 
		VALUES (:id, :user_id, :promotion_id, :promotion_use_count)`, mm)
	return err
}

func getUserPromo(db *sqlx.DB, userID, promoID string) (mysql.UserPromo, error) {
	up := mysql.UserPromo{}
	err := db.Get(&up, `SELECT * FROM users_promotions WHERE user_id=? AND promotion_id=?`, userID, promoID)
	return up, err
}

func truncateAll(db *sqlx.DB) error {
	_, err := db.Exec(`TRUNCATE TABLE users_promotions; `)
	if err != nil {
		return err
	}

	_, err = db.Exec(`TRUNCATE TABLE users;`)
	if err != nil {
		return err
	}

	_, err = db.Exec(`TRUNCATE TABLE promotions; `)
	if err != nil {
		return err
	}

	return err
}

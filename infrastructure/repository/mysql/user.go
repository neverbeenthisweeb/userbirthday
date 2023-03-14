package mysql

import (
	"context"
	"userbirthday/common"
	"userbirthday/model"

	"github.com/jmoiron/sqlx"
	uuid "github.com/satori/go.uuid"
)

const (
	queryGetVerifiedBirthdayUsers = `
SELECT
	u.id as user_id,
	u.name ,
	u.email ,
	u.phone ,
	u.is_verified ,
	u.birthdate ,
	p.id as promotion_id,
	p.code,
	p.type ,
	up.promotion_use_count as use_count,
	p.use_limit ,
	p.valid_from ,
	p.valid_to
FROM
	users u
LEFT JOIN users_promotions up ON
	u.id = up.user_id
LEFT JOIN promotions p ON
	p.id = up.promotion_id
WHERE u.is_verified = TRUE
	AND MONTH(u.birthdate) = MONTH(NOW())
	AND DAY(u.birthdate) = DAY(NOW())`
)

type UserRepository struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) *UserRepository {
	return &UserRepository{db}
}

func (ur *UserRepository) GetVerifiedBirthdayUsers(ctx context.Context) ([]model.User, error) {
	rows, err := ur.db.Queryx(queryGetVerifiedBirthdayUsers)
	if err != nil {
		common.LogErr(ctx, "Failed to query", err)
		return nil, err
	}

	var users []model.User
	usersMap := map[string]model.User{}

	for rows.Next() {
		var tmpUserPromo dtoGetVerifiedBirthdayUsers
		err := rows.StructScan(&tmpUserPromo)
		if err != nil {
			common.LogErr(ctx, "Failed scan struct", err)
			return []model.User{}, err
		}

		if v, ok := usersMap[tmpUserPromo.UserID]; !ok {
			usersMap[tmpUserPromo.UserID] = model.User{
				ID:         tmpUserPromo.UserID,
				Name:       tmpUserPromo.Name,
				Email:      tmpUserPromo.Email,
				Phone:      tmpUserPromo.Phone,
				IsVerified: tmpUserPromo.IsVerified,
				Birthdate:  tmpUserPromo.Birthdate,
				Promos: []model.Promo{
					{
						ID:        tmpUserPromo.PromoID.String,
						Code:      tmpUserPromo.Code.String,
						Type:      tmpUserPromo.Type.String,
						UseCount:  int(tmpUserPromo.UseCount.Int64),
						UseLimit:  int(tmpUserPromo.UseLimit.Int64),
						ValidFrom: tmpUserPromo.ValidFrom.Time,
						ValidTo:   tmpUserPromo.ValidTo.Time,
					},
				},
			}
		} else {
			v.Promos = append(v.Promos, model.Promo{
				ID:        tmpUserPromo.PromoID.String,
				Code:      tmpUserPromo.Code.String,
				Type:      tmpUserPromo.Type.String,
				UseCount:  int(tmpUserPromo.UseCount.Int64),
				UseLimit:  int(tmpUserPromo.UseLimit.Int64),
				ValidFrom: tmpUserPromo.ValidFrom.Time,
				ValidTo:   tmpUserPromo.ValidTo.Time,
			})
			usersMap[tmpUserPromo.UserID] = v
		}
	}

	for _, v := range usersMap {
		users = append(users, v)
	}

	return users, nil
}

func (ur *UserRepository) UpdateUserPromo(ctx context.Context, userID, promoID string) error {
	id := uuid.NewV4().String()

	_, err := ur.db.Exec(`INSERT INTO users_promotions (id, user_id, promotion_id) 
	VALUES (?, ?, ?)`, id, userID, promoID)
	if err != nil {
		common.LogErr(ctx, "Failed exec", err)
		return err
	}

	return nil
}

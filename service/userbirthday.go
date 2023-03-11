package service

import (
	"context"
	"userbirthday/common"
	"userbirthday/infrastructure"
	"userbirthday/infrastructure/repository"
)

type UserBirthday struct {
	repoUser repository.UserRepository
}

func NewUserBirthday(infra *infrastructure.Infrastructure) UserBirthday {
	return UserBirthday{
		repoUser: infra.RepoUser(),
	}
}

// GiveBirthdayPromo gives birthday promo to verified birthday users
func (ub UserBirthday) GiveBirthdayPromo(ctx context.Context) error {
	// Get verified birthday users
	users, err := ub.repoUser.GetVerifiedBirthdayUsers(ctx)
	if err != nil {
		common.LogErr(ctx, err)
		return err
	}

	for _, usr := range users {
		// Validate if user don't have birthday promo yet
		if usr.HasBirthdayPromo() {
			common.LogInfo(ctx, "Skip user "+usr.ID+". This user already has birthday promo")
			continue
		}

		// Create birthday promo

		// Set birthday promo to user

		// Send notification
	}

	return nil
}

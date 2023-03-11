package service

import (
	"context"
	"fmt"
	"userbirthday/common"
	"userbirthday/infrastructure"
	"userbirthday/infrastructure/notification"
	"userbirthday/infrastructure/repository"
	"userbirthday/model"
)

type UserBirthday struct {
	repoUser  repository.UserRepository
	repoPromo repository.PromoRepository
	notif     notification.Notification
}

func NewUserBirthday(infra *infrastructure.Infrastructure) UserBirthday {
	return UserBirthday{
		repoUser:  infra.RepoUser(),
		repoPromo: infra.RepoPromo(),
		notif:     infra.Notification(),
	}
}

// GiveBirthdayPromo gives birthday promo to verified birthday users
func (ub UserBirthday) GiveBirthdayPromo(ctx context.Context) error {
	// Get verified birthday users
	users, err := ub.repoUser.GetVerifiedBirthdayUsers(ctx)
	if err != nil {
		common.LogErr(ctx, "Failed to get verified birthday users", err)
		return err
	}

	for _, user := range users {
		// Validate if user don't have birthday promo yet
		if user.HasBirthdayPromo() {
			common.LogInfo(ctx, "Skip user ID="+user.ID+". This user already has birthday promo")
			continue
		}

		// Create birthday promo
		// TODO: Can implement DB transaction between create promo and set promo
		// to reduce stale data in DB
		promo := model.NewBirthdayPromo(user.Name)
		createdPrm, err := ub.repoPromo.Create(ctx, promo)
		if err != nil {
			common.LogErr(ctx, "Failed to create birthday promo", err)
			return err
		}

		// Set birthday promo to user
		err = ub.repoUser.SetPromo(ctx, user.ID, createdPrm.ID)
		if err != nil {
			common.LogErr(ctx, fmt.Sprintf("Failed to set user ID=%s with promo ID=%s", user.ID, createdPrm.ID), err)
			return err
		}

		// Send notification
		// TODO: As an improvement, can send notification in async
		err = ub.sendNotification(ctx, user, createdPrm)
		if err != nil {
			common.LogErr(ctx, "Failed to send notification", err)
		}
	}

	return nil
}

func (ub UserBirthday) sendNotification(ctx context.Context, usr model.User, prm model.Promo) error {
	if usr.Email != "" {
		err := ub.notif.Send(ctx, notification.NotificationRequest{
			NotificationType: notification.NotificationTypeEmail,
			Target:           usr.Email,
			TemplateID:       "email.birthday",
			TemplateData: map[string]string{
				"username":  usr.Name,
				"promocode": prm.Code,
			},
		})
		if err != nil {
			common.LogWarn(ctx, "Failed to send email notification", err)
		}
	}

	if usr.Phone != "" {
		err := ub.notif.Send(ctx, notification.NotificationRequest{
			NotificationType: notification.NotificationTypeWA,
			Target:           usr.Phone,
			TemplateID:       "wa.birthday",
			TemplateData: map[string]string{
				"username":  usr.Name,
				"promocode": prm.Code,
			},
		})
		if err != nil {
			common.LogWarn(ctx, "Failed to send WA notification", err)
		}
	}

	return nil
}

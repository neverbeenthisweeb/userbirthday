package infrastructure

import (
	"userbirthday/infrastructure/notification"
	"userbirthday/infrastructure/repository"
)

type Infrastructure struct {
	notification notification.Notification
	repoUser     repository.UserRepository
	repoPromo    repository.PromoRepository
}

func NewInfrastructure() *Infrastructure {
	return &Infrastructure{}
}

func (i *Infrastructure) Notification() notification.Notification {
	return i.notification
}

func (i *Infrastructure) SetNotification(notif notification.Notification) {
	i.notification = notif
}

func (i *Infrastructure) RepoUser() repository.UserRepository {
	return i.repoUser
}

func (i *Infrastructure) SetRepoUser(repo repository.UserRepository) {
	i.repoUser = repo
}

func (i *Infrastructure) RepoPromo() repository.PromoRepository {
	return i.repoPromo
}

func (i *Infrastructure) SetRepoPromo(repo repository.PromoRepository) {
	i.repoPromo = repo
}

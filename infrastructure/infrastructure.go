package infrastructure

import (
	"fmt"
	"os"
	"userbirthday/infrastructure/notification"
	"userbirthday/infrastructure/notification/defaultnotification"
	"userbirthday/infrastructure/repository"
	"userbirthday/infrastructure/repository/mysql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type Infrastructure struct {
	notification notification.Notification
	repoUser     repository.UserRepository
	repoPromo    repository.PromoRepository
}

func NewInfrastructure() *Infrastructure {
	db := sqlx.MustOpen(
		"mysql",
		// "root:password@tcp(127.0.0.1:3332)/userbirthday?parseTime=true",
		fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
			os.Getenv("DB_USER"),
			os.Getenv("DB_PASSWORD"),
			os.Getenv("DB_URL"),
			os.Getenv("DB_PORT"),
			os.Getenv("DB_DATABASE"),
		),
	)

	if err := db.Ping(); err != nil {
		panic("Failed to ping DB. Error: " + err.Error())
	}

	return &Infrastructure{
		notification: defaultnotification.NewDefaulNotification(),
		repoUser:     mysql.NewUserRepository(db),
		repoPromo:    mysql.NewPromoRepository(db),
	}
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

package main

import (
	"os"
	"os/signal"
	"syscall"
	"userbirthday/common"
	"userbirthday/infrastructure"
	"userbirthday/service"

	"github.com/joho/godotenv"
	"github.com/robfig/cron/v3"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		panic("Failed to load env. Error: " + err.Error())
	}

	infra := infrastructure.NewInfrastructure()
	svc := service.NewUserBirthday(infra)

	scheduler := cron.New()
	defer scheduler.Stop()
	scheduler.AddFunc(os.Getenv("USER_BIRTHDAY_CRON_EVENT"), func() {
		ctx := common.ContextWithRequestID()
		err := svc.GiveBirthdayPromo(ctx)
		if err != nil {
			common.LogErr(ctx, "Failed to give birthday promo with request_id="+ctx.Value(common.CtxKeyRequestID).(string), err)
		}
	})
	scheduler.Start()

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	<-sig
}

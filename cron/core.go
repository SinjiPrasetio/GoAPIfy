package cron

import (
	"GoAPIfy/service/appService"
)

type Cron interface {
	Start()
}

type cron struct {
	appService appService.AppService
}

func NewCron(appService appService.AppService) *cron {
	return &cron{appService}
}

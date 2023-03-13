package cron

import (
	"GoAPIfy/core/helper"
	"fmt"

	cronJob "github.com/robfig/cron/v3"
)

func (c *cron) Start() {
	job := cronJob.New()
	job.AddFunc("@daily", func() {
		// Do Something
	})

	fmt.Println(helper.ColorizeCmd(helper.Magenta, "Cron jobs started."))
	job.Start()

}

package cron

import (
	"GoAPIfy/core/helper"
	"fmt"

	cronJob "github.com/robfig/cron/v3"
)

/*
*

	Starts a new cron job and schedules a function to be executed once a day.

	@param {cron} c - The cron instance on which the function will be scheduled

	@returns {void}
*/
func (c *cron) Start() {
	// Create a new cron job
	job := cronJob.New()

	// Core function. IMPORTANT: DO NOT CHANGE THIS
	// Start -----
	job.AddFunc("*/5 * * * *", DeleteExpiredTemporaryFiles)
	// End -----

	// Schedule a function to be executed once a day
	job.AddFunc("@daily", func() {
		// Do something
	})

	// Start the cron job
	fmt.Println(helper.ColorizeCmd(helper.Magenta, "Cron jobs started."))
	job.Start()
}

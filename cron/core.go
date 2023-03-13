package cron

import (
	"GoAPIfy/service/appService"
)

/*
*
Defines a new interface for managing cron jobs.
@interface {Cron} - The interface for the cron job service
*/
type Cron interface {
	Start()
}

/*
*

	Defines a new type that implements the Cron interface.
	@typedef {cron} - The implementation of the Cron interface
	@property {appService} appService - The application service instance to be used in the cron job implementation
*/
type cron struct {
	appService appService.AppService
}

/*
*

	Creates a new instance of the cron job service.
	@param {appService} appService - The application service instance to be used in the cron job implementation
	@returns {cron} - A new instance of the cron job service
*/
func NewCron(appService appService.AppService) *cron {
	return &cron{appService}
}

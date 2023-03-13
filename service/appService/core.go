/**

This module contains the services required for the GoAPIfy application.
The AppService struct defines the type that encapsulates the application services,
with a Model property that is of type model.Model, and a MeiliSearch property that is of type *meilisearch.Client.
These properties can be accessed and used by the various services of the application.
*/

package appService

import (
	"GoAPIfy/model"

	"github.com/meilisearch/meilisearch-go"
)

/*
*

	Defines a new type for application services.
	@typedef {AppService} - The type that encapsulates the application services
	@property {model.Model} Model - The model instance to be used in the application service
	@property {meilisearch.Client} MeiliSearch - The MeiliSearch client instance to be used in the application service
*/
type AppService struct {
	Model       model.Model
	MeiliSearch *meilisearch.Client
}

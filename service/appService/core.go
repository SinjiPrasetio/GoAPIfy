package appService

import (
	"GoAPIfy/model"

	"github.com/meilisearch/meilisearch-go"
)

type AppService struct {
	Model       model.Model
	MeiliSearch *meilisearch.Client
}

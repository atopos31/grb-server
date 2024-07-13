package core

import (
	"gvb/config"

	"github.com/meilisearch/meilisearch-go"
)

func InitMeiliSearch(conf config.Meilisearch) *meilisearch.Client {
	client := meilisearch.NewClient(meilisearch.ClientConfig{
		Host:   conf.Dsn(),
		APIKey: conf.ApiKey,
	})

	return client
}

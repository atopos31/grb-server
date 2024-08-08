package core

import (
	"gvb/config"

	"github.com/meilisearch/meilisearch-go"
)

func NewMeiliSearchClient(conf config.Meilisearch) *meilisearch.Client {
	client := meilisearch.NewClient(meilisearch.ClientConfig{
		Host:   conf.Dsn(),
		APIKey: conf.ApiKey,
	})

	_, err := client.Health()
	if err != nil {
		panic(err)
	}

	return client
}

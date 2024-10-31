package core

import "github.com/google/wire"

var ProvideSet = wire.NewSet(
	NewGormDB,
	NewRedisCache,
	NewMeiliSearchClient,
)

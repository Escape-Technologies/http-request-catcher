package database

import (
	"github.com/Escape-Technologies/http-request-catcher/api"
	"github.com/Escape-Technologies/http-request-catcher/internal"
	"github.com/go-redis/redis"
)

var (
	Provider api.DatabaseProvider
	Db       *redis.Client
)

// Connect to the database.
// Use provider strategy.
func Connect(provider api.DatabaseProvider) {
	Provider = provider

	switch provider {
	case api.REDIS:
		Db = redisConnect(
			internal.REDIS_HOST,
			internal.REDIS_PORT,
			internal.REDIS_PWD,
			internal.REDIS_DB,
		)
	default:
		panic("Unknown database type")
	}
}

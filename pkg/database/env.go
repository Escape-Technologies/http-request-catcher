package database

import (
	"github.com/Escape-Technologies/http-request-catcher/api"
	"github.com/Escape-Technologies/http-request-catcher/internal"
	"github.com/go-redis/redis"
	"github.com/nitishm/go-rejson"
)

var (
	Provider  api.DatabaseProvider
	Db        *redis.Client
	DbHandler *rejson.Handler
)

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

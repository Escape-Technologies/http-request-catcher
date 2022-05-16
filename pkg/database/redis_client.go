package database

import (
	"fmt"

	"github.com/go-redis/redis"
	"github.com/nitishm/go-rejson"
)

func redisJsonSetup(client *redis.Client) {
	rh := rejson.NewReJSONHandler()

	rh.SetGoRedisClient(client)

	DbHandler = rh
}

func redisConnect(host string, port string, password string, db int) *redis.Client {
	fmt.Println("Connecting to redis server: " + host + ":" + port)
	fmt.Printf("Connecting to redis database: %d\n", db)

	rdb := redis.NewClient(&redis.Options{
		Addr:     host + ":" + port,
		Password: password,
		DB:       db,
	})

	if _, err := rdb.Ping().Result(); err != nil {
		panic(err)
	}

	fmt.Println("Connected to redis server: " + host + ":" + port)

	redisJsonSetup(rdb)

	return rdb
}

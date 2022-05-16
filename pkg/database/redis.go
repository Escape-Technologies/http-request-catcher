package database

import (
	"encoding/json"
	"fmt"

	"github.com/Escape-Technologies/http-request-catcher/api"
	"github.com/Escape-Technologies/http-request-catcher/internal"
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
		fmt.Println("Error connecting to redis server: " + err.Error())
		panic(err)
	}

	fmt.Println("Connected to redis server: " + host + ":" + port)

	redisJsonSetup(rdb)

	return rdb
}

func storeRequestRedis(r *api.BucketEntry) {

	packedEntry, _ := json.Marshal(r)

	_ = Db.HSet(r.BucketId, r.Id, packedEntry)
	_ = Db.Expire(r.BucketId, internal.REQUEST_EXPIRATION_TIME)
}

func getBucketRedis(id string) api.Bucket {
	data := Db.HGetAll(id)

	entries := make(api.Bucket, 0)
	for _, entry := range data.Val() {
		jsonEntry := api.BucketEntry{}

		_ = json.Unmarshal([]byte(entry), &jsonEntry)
		entries = append(entries, jsonEntry)
	}

	return entries
}

func deleteBucketRedis(id string) int {
	count := 0
	fields := Db.HGetAll(id)
	for field := range fields.Val() {
		_ = Db.HDel(id, field)
		count += 1
	}

	return count
}

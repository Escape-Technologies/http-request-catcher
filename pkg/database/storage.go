package database

import (
	"encoding/json"

	"github.com/Escape-Technologies/http-request-catcher/api"
	"github.com/Escape-Technologies/http-request-catcher/internal"
)

func storeRequestRedis(r *api.BucketEntry) {

	packedEntry, _ := json.Marshal(r)

	_ = Db.HSet(r.BucketId, r.Id, packedEntry)
	_ = Db.Expire(r.BucketId, internal.REQUEST_EXPIRATION_TIME)
}

func StoreRequest(r *api.BucketEntry) {
	switch Provider {
	case api.REDIS:
		storeRequestRedis(r)
	default:
		panic("Unknown database type")
	}
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

func GetBucket(id string) api.Bucket {
	switch Provider {
	case api.REDIS:
		return getBucketRedis(id)
	default:
		panic("Unknown database type")
	}
}

func DeleteBucket(id string) int {
	switch Provider {
	case api.REDIS:

		count := 0
		fields := Db.HGetAll(id)
		for field := range fields.Val() {
			_ = Db.HDel(id, field)
			count += 1
		}

		return count
	default:
		panic("Unknown database type")
	}
}

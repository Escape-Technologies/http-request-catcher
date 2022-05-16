package database

import (
	"github.com/Escape-Technologies/http-request-catcher/api"
)

func StoreRequest(r *api.BucketEntry) {
	switch Provider {
	case api.REDIS:
		storeRequestRedis(r)
	default:
		panic("Unknown database type")
	}
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
		return deleteBucketRedis(id)
	default:
		panic("Unknown database type")
	}
}

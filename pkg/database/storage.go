package database

import (
	"github.com/Escape-Technologies/http-request-catcher/api"
)

// Stores a request in a bucket.
func StoreRequest(r *api.BucketEntry) {
	switch Provider {
	case api.REDIS:
		storeRequestRedis(r)
	default:
		panic("Unknown database type")
	}
}

// Returns a bucket.
func GetBucket(id string) api.Bucket {
	switch Provider {
	case api.REDIS:
		return getBucketRedis(id)
	default:
		panic("Unknown database type")
	}
}

// Deletes a bucket and all its entries.
func DeleteBucket(id string) int {
	switch Provider {
	case api.REDIS:
		return deleteBucketRedis(id)
	default:
		panic("Unknown database type")
	}
}

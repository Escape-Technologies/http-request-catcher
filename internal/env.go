package internal

import (
	"os"
	"strconv"
	"time"
)

func GetEnv(key string, defaultValue string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return defaultValue
}

func _atoi(value string) int {
	if i, err := strconv.Atoi(value); err == nil {
		return i
	}
	return 0
}

var (
	CATCHER_PORT            = GetEnv("CATCHER_PORT", "8080")
	REDIS_HOST              = GetEnv("REDIS_HOST", "localhost")
	REDIS_PORT              = GetEnv("REDIS_PORT", "6379")
	REDIS_PWD               = GetEnv("REDIS_PWD", "")
	REDIS_DB                = _atoi(GetEnv("REDIS_DB", "0"))
	REQUEST_EXPIRATION_TIME = time.Second * time.Duration(_atoi(GetEnv("REQUEST_EXPIRATION_TIME", "120")))
)

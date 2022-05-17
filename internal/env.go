package internal

import (
	"os"
	"strconv"
	"time"
)

func _getEnv(key string, defaultValue string) string {
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
	CATCHER_PORT            = _getEnv("CATCHER_PORT", "8080")
	REDIS_HOST              = _getEnv("REDIS_HOST", "localhost")
	REDIS_PORT              = _getEnv("REDIS_PORT", "6379")
	REDIS_PWD               = _getEnv("REDIS_PWD", "")
	REDIS_DB                = _atoi(_getEnv("REDIS_DB", "0"))
	REQUEST_EXPIRATION_TIME = time.Second * time.Duration(_atoi(_getEnv("REQUEST_EXPIRATION_TIME", "120")))
)

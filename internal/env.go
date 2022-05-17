package internal

import (
	"os"
	"strconv"
	"time"
)

// Get the environment variable value.
// If the variable is not set, return the default value.
func _getEnv(key string, defaultValue string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return defaultValue
}

// Convert a string to an int.
// If the string is not convertable, return 0.
func _atoi(value string) int {
	if i, err := strconv.Atoi(value); err == nil {
		return i
	}
	return 0
}

// Storage for all environment variables.
// Provides a default value if the environment variable is not set.
var (
	CATCHER_PORT            = _atoi(_getEnv("CATCHER_PORT", "8080"))
	REDIS_HOST              = _getEnv("REDIS_HOST", "localhost")
	REDIS_PORT              = _getEnv("REDIS_PORT", "6379")
	REDIS_PWD               = _getEnv("REDIS_PASSWORD", "")
	REDIS_DB                = _atoi(_getEnv("REDIS_DB", "0"))
	REQUEST_EXPIRATION_TIME = time.Second * time.Duration(_atoi(_getEnv("REQUEST_EXPIRATION_TIME", "120")))
)

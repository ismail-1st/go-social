package env

import (
	"os"
	"strconv"
	"time"
)

func GetString(key string, fallback string) string {
	val, ok := os.LookupEnv(key)
	if !ok {
		return fallback
	}

	return val
}

func GetDuration(envKey string, fallback string) time.Duration {
	val := os.Getenv(envKey)
	if val == "" {
		// environment variable not set -> use fallback
		duration, _ := time.ParseDuration(fallback)
		return duration
	}

	parsedDuration, err := time.ParseDuration(val)
	if err != nil {
		// invalid format in env var -> use fallback
		duration, _ := time.ParseDuration(fallback)
		return duration
	}

	return parsedDuration
}

func GetInt(key string, fallback int) int {
	val, ok := os.LookupEnv(key)
	if !ok {
		return fallback
	}

	intVal, err := strconv.Atoi(val)

	if err != nil {
		return fallback
	}
	return intVal
}

func GetBool(key string, fallback bool) bool {
	if val, ok := os.LookupEnv(key); ok {
		if b, err := strconv.ParseBool(val); err == nil {
			return b
		}
	}
	return fallback
}

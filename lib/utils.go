package lib

import (
	"os"
	"strings"
)

// GetEnv an environment variable. If no such envar is found, `fallback` is used.
func GetEnv(key string, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func IsLocalRequest(addr string) bool {
	return strings.HasPrefix(addr, "[::1]") || strings.HasPrefix(addr, "127.0.0.1")
}

package env

import "os"

const (
	DefaultPort = "50001"

	EnvPort = "PORT"
)

func GetEnvFallback(key, fallback string) string {
	if _, isSet := os.LookupEnv(key); !isSet {
		return fallback
	}

	return os.Getenv(key)
}

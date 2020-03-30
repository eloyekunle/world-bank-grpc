package env

import "os"

const (
	EnvPort = "PORT"
	EnvHost = "HOST"
)

func GetEnvFallback(key, fallback string) string {
	if _, isSet := os.LookupEnv(key); !isSet {
		return fallback
	}

	return os.Getenv(key)
}

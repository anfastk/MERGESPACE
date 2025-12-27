package config

import "os"

type Config struct {
	HTTPPort string
	AuthGRPC string
}

func Load() *Config {
	return &Config{
		HTTPPort: getEnv("HTTP_PORT", "8080"),
		AuthGRPC: getEnv("AUTH_GRPC_ADDR", "localhost:50051"),
	}
}

func getEnv(key, def string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return def
}

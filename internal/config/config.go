package config

import (
	"fmt"
	"os"
)

type Config struct {
	DBURL      string
	JWTSecret  string
	JWTIssuer  string
	JWTExpireH int
}

func Load() *Config {
	return &Config{
		DBURL:      env("DATABASE_URL", "postgres://postgres:postgres@localhost:5432/school?sslmode=disable"),
		JWTSecret:  env("JWT_SECRET", "dev-secret-change-me"),
		JWTIssuer:  env("JWT_ISSUER", "school-platform"),
		JWTExpireH: envInt("JWT_EXPIRE_HOURS", 24),
	}
}

func env(key, def string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return def
}
func envInt(key string, def int) int {
	if v := os.Getenv(key); v != "" {
		var out int
		_, _ = fmt.Sscanf(v, "%d", &out)
		if out != 0 {
			return out
		}
	}
	return def
}

package config

import "os"

type Config struct {
	AppName string
}

func Load() Config {
	return Config{AppName: envOr("APP_NAME", "MultiStock")}
}

func envOr(key, def string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return def
}

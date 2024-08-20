package config

import (
	"github.com/spf13/cast"
	"os"
)

type Config struct {
	HTTP_PORT   string
	DB_HOST     string
	DB_PORT     string
	DB_USER     string
	DB_PASSWORD string
	DB_NAME     string
}

func Load() *Config {
	config := &Config{}

	config.HTTP_PORT = cast.ToString(coalesce("HTTP_PORT", "productserver:8082"))
	config.DB_HOST = cast.ToString(coalesce("DB_HOST", "postgresdb1"))
	config.DB_PORT = cast.ToString(coalesce("DB_PORT", 5432))
	config.DB_USER = cast.ToString(coalesce("DB_USER", "postgres"))
	config.DB_PASSWORD = cast.ToString(coalesce("DB_PASSWORD", "dodi"))
	config.DB_NAME = cast.ToString(coalesce("DB_NAME", "product"))

	return config
}

func coalesce(key string, value interface{}) interface{} {
	val, exists := os.LookupEnv(key)
	if exists {
		return val
	}
	return value
}

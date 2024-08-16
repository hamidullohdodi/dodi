package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/cast"
)

type Config struct {
	HTTP_PORT           string
	USER_SERVICE        string
	USER_ROUTER         string
	DB_HOST             string
	DB_PORT             string
	DB_USER             string
	DB_PASSWORD         string
	DB_NAME             string
	SIGNING_KEY         string
	REFRESH_SIGNING_KEY string
	REDIS_ADDRESS     string
	REDIS_PASSWORD    string
	REDIS_DB          int

	ACCESS_TOKEN_KEY  string
	REFRESH_TOKEN_KEY string
	
}

func Load() Config {
	if err := godotenv.Load(".env"); err != nil {
		log.Print("No .env file found?")
	}

	config := Config{}
	config.HTTP_PORT = cast.ToString(Coalesce("HTTP_PORT", "auth:8081"))
	config.DB_HOST = cast.ToString(Coalesce("DB_HOST", "postgres"))
	config.DB_PORT = cast.ToString(Coalesce("DB_PORT", 5432))
	config.DB_USER = cast.ToString(Coalesce("DB_USER", "postgres"))
	config.DB_PASSWORD = cast.ToString(Coalesce("DB_PASSWORD", "0101"))
	config.DB_NAME = cast.ToString(Coalesce("DB_NAME", "auth_exam"))
	config.USER_SERVICE = cast.ToString(Coalesce("USER_SERVICE", "auth:8081"))
	config.SIGNING_KEY = cast.ToString(Coalesce("SIGNING_KEY", "secret"))
	config.REFRESH_SIGNING_KEY = cast.ToString(Coalesce("REFRESH_SIGNING_KEY", "secret1"))

	config.REDIS_ADDRESS = cast.ToString(Coalesce("REDIS_ADDRESS", "redis:6379"))
	config.REDIS_PASSWORD = cast.ToString(Coalesce("REDIS_PASSWORD", ""))
	config.REDIS_DB = cast.ToInt(Coalesce("REDIS_DB", "0"))

	config.ACCESS_TOKEN_KEY = cast.ToString(Coalesce("ACCESS_TOKEN_KEY", "key"))
	config.REFRESH_TOKEN_KEY = cast.ToString(Coalesce("REFRESH_TOKEN_KEY", "key"))
	return config
}

func Coalesce(key string, defaultValue interface{}) interface{} {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return defaultValue
}

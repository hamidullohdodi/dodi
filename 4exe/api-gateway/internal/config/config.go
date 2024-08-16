package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/cast"
)

type Config struct {
	BOOKING      string
	USER         string
	SIGNING_KEY2 string
	SIGNING_KEY  string
	API_GATEWAY  string
	USER_ROUTER  string
	DB_PASSWORD  string
	MONGODB_URI  string
	GMAIL_CODE   string
	DB_HOST      string
	DB_PORT      string
	DB_USER      string
	DB_NAME      string
	MONGODB_URL  string
}

func Load() *Config {

	if err := godotenv.Load("/.env"); err != nil {
		log.Print("No .env file found")
	}

	config := Config{}
	config.DB_HOST = cast.ToString(Coalesce("DB_HOST", "postgres"))
	config.DB_PORT = cast.ToString(Coalesce("DB_PORT", "5432"))
	config.DB_USER = cast.ToString(Coalesce("DB_USER", "postgres"))
	config.DB_PASSWORD = cast.ToString(Coalesce("DB_PASSWORD", "0101"))
	config.DB_NAME = cast.ToString(Coalesce("DB_NAME", "auth_exam"))
	config.USER = cast.ToString(Coalesce("AUTH_SERVICE", ":8081"))
	config.SIGNING_KEY = cast.ToString(Coalesce("SIGNING_KEY", "secret"))


	config.API_GATEWAY = cast.ToString(Coalesce("API_GATEWAY", ":9090"))
	config.MONGODB_URL = cast.ToString(Coalesce("MONGODB_URL", "mongodb://mongodb:27017"))
	config.BOOKING = cast.ToString(Coalesce("BOOKING_SERVICE", ":50053"))

	return &config
}

func Coalesce(key string, defaultValue interface{}) interface{} {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return defaultValue
}

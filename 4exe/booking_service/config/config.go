package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/cast"
)

type Config struct {
	BookingService string
	RedisAddr      string
	RedisPassword  string
	RedisDB        int
	MongoDb        MongoDbConfig
}
type MongoDbConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
}

func Load() *Config {
	if err := godotenv.Load(".env"); err != nil {
		log.Print("No .env file found")
	}

	config := Config{}
	config.BookingService = cast.ToString(Coalesce("BOOKING_SERVICE", "booking:50053"))
	config.RedisPassword = cast.ToString(Coalesce("REDIS_PASSWORD", ""))
	config.RedisDB = cast.ToInt(Coalesce("REDIS_DB", 0))
	config.MongoDb.Host = cast.ToString(Coalesce("MONGO_HOST", "mongodb"))
	config.MongoDb.Port = cast.ToString(Coalesce("MONGO_PORT", "27017"))
	config.MongoDb.User = cast.ToString(Coalesce("MONGO_USER", ""))
	config.MongoDb.Password = cast.ToString(Coalesce("MONGO_PASSWORD", ""))
	config.MongoDb.DBName = cast.ToString(Coalesce("MONGO_DB_NAME", "localhos"))
	

	return &config
}

func Coalesce(key string, defaultValue interface{}) interface{} {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return defaultValue
}

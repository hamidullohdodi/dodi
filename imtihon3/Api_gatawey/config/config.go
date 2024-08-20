package config

import (
	"github.com/joho/godotenv"
	"github.com/spf13/cast"
	"log"
	"os"
)

type Config struct {
	HTTP_PORT            string
	API_GATEWAY_PORT     string
	AUTH_SERVICE_PORT    string
	PRODUCT_SERVICE_PORT string
	SIGNING_KEY          string
}

func Load() *Config {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Cannot load .env ", err)
	}

	cfg := &Config{}

	cfg.API_GATEWAY_PORT = cast.ToString(coalesce("HTTP_PORT", "apigateway:50051"))
	cfg.AUTH_SERVICE_PORT = cast.ToString(coalesce("AUTH_SERVICE_PORT", "authservice:8083"))
	cfg.PRODUCT_SERVICE_PORT = cast.ToString(coalesce("PRODUCTS_SERVICE_PORT", "productserver:8082"))
	cfg.SIGNING_KEY = cast.ToString(coalesce("SIGNING_KEY", "GARD"))

	return cfg
}

func coalesce(key string, value interface{}) interface{} {
	val, exists := os.LookupEnv(key)
	if exists {
		return val
	}
	return value
}

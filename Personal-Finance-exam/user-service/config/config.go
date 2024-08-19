package config

import (
	"fmt"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

type DatabaseConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Name     string
}

type Config struct {
	AuthPort string
	AuthHost string

	UserPort string
	UserHost string

	Database DatabaseConfig
}

func Load(path string) (*Config, error) {
	err := godotenv.Load(path + "/.env")
	if err != nil {
		return nil, fmt.Errorf("failed to load .env file: %w", err)
	}

	conf := viper.New()
	conf.AutomaticEnv()

	cfg := Config{
		AuthPort: conf.GetString("AUTH_PORT"),
		AuthHost: conf.GetString("AUTH_HOST"),
		UserPort: conf.GetString("USER_PORT"),
		UserHost: conf.GetString("USER_HOST"),
		Database: DatabaseConfig{
			Host:     conf.GetString("DB_HOST"),
			Port:     conf.GetString("DB_PORT"),
			User:     conf.GetString("DB_USER"),
			Password: conf.GetString("DB_PASSWORD"),
			Name:     conf.GetString("DB_NAME"),
		},
	}
	return &cfg, nil
}

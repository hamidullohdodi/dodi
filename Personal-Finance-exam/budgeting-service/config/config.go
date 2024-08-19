package config

import (
	"fmt"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

type Config struct {
	BudgetingPort string
	BudgetingHost string
}

func Load(path string) (*Config, error) {
	err := godotenv.Load(path + "/.env")
	if err != nil {
		return nil, fmt.Errorf("failed to load .env file: %w", err)
	}

	conf := viper.New()
	conf.AutomaticEnv()

	cfg := Config{
		BudgetingPort: conf.GetString("BUDGETING_SERVICE_PORT"),
		BudgetingHost: conf.GetString("BUDGETING_SERVICE_HOST"),
	}

	return &cfg, nil
}

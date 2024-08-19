package config

import (
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

type Config struct {
	UserHost      string
	UserPort      string
	BudgetingHost string
	BudgetingPort string
	ApiHost       string
	ApiPort       string
	SigningKey    string
}

func Load(path string) (*Config, error) {

	err := godotenv.Load(path + "/.env")
	if err != nil {
		return nil, err
	}

	conf := viper.New()
	conf.AutomaticEnv()

	cfg := Config{
		UserHost:      conf.GetString("USER_SERVICE_HOST"),
		UserPort:      conf.GetString("USER_SERVICE_PORT"),
		BudgetingHost: conf.GetString("BUDGETING_SERVICE_HOST"),
		BudgetingPort: conf.GetString("BUDGETING_SERVICE_PORT"),
		ApiHost:       conf.GetString("API_GATEWAY_HOST"),
		ApiPort:       conf.GetString("API_GATEWAY_PORT"),
		SigningKey:    conf.GetString("SIGNING_KEY"),
	}

	return &cfg, nil
}

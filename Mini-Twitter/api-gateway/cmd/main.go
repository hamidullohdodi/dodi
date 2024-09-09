package main

import (
	"apigateway/api"
	config2 "apigateway/pkg/config"
	"apigateway/pkg/logger"
	"github.com/casbin/casbin/v2"
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
	"os"
)

func main() {
	appLogger := logger.NewLogger()

	config := config2.Config{}

	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	path, err := os.Getwd()
	if err != nil {
		appLogger.Error("Failed to get current working directory")
		return
	}

	casbinEnforcer, err := casbin.NewEnforcer(path+"/pkg/config/model.conf", path+"/pkg/config/policy.csv")
	if err != nil {
		panic(err)
	}

	controller := api.NewRouter(&config, ch, appLogger, casbinEnforcer)
	controller.Run(":8087")
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

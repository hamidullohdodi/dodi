package main

import (
	"api/api"
	cfg "api/config"
	logger "api/log"
	srvc "api/service"
	"fmt"
	"log"
	"os"

	"github.com/casbin/casbin/v2"
	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
	log := logger.InitLogger()

	config, err := cfg.Load(".")
	if err != nil {
		fmt.Println(err)
	}
	service, err := srvc.NewServiceManager(*config)
	if err != nil {
		log.Error("Failed to initialize service manager")
		return
	}
	conn, err := amqp.Dial("amqp://guest:guest@rabbitmq:5672/")

	failOnError(err, "Failed to connect to RabbitMQ with api")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	path, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	}

	casbinEnforcer, err := casbin.NewEnforcer(path+"/config/model.conf", path+"/config/policy.csv")
	if err != nil {
		panic(err)
	}

	controller := api.NewRouter(casbinEnforcer, service, log, ch)
	target := fmt.Sprintf("%s:%s", config.ApiHost, config.ApiPort)
	controller.Run(target)
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}

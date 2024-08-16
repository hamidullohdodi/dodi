package main

import (
	"api-geteway/api"
	"api-geteway/internal/config"
	"api-geteway/service"
	"log"
	"time"

	"github.com/casbin/casbin"
	rmq "github.com/rabbitmq/amqp091-go"
)

func main() {
	cfg := config.Load()
	casbinEnforcer := casbin.NewEnforcer("./internal/config/model.conf", "./internal/config/policy.csv")
	time.Sleep(15*time.Second)
	conn, err := rmq.Dial("amqp://guest:guest@rabbitmq:5672/")
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %v", err)
	}
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Faniled to open a channel")
	defer ch.Close()
	servis, err := service.NewService(cfg)
	if err != nil {
		log.Panicf("Failed to create service: %v", err)
	}
	r := api.NewRouter(servis, ch, casbinEnforcer)
	r.Run("gateway:9090")

}

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}

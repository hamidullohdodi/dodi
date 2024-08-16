package main

import (
	"booking_service/config"
	pb "booking_service/genproto/booking"
	messages "booking_service/massBrok"
	"booking_service/service"
	Mongo "booking_service/storage/mongodb"
	"booking_service/storage/repo"
	"context"
	"fmt"
	"log"
	"net"
	"sync"
	"time"

	rmq "github.com/rabbitmq/amqp091-go"
	"google.golang.org/grpc"
)

func main() {
	listener, err := net.Listen("tcp", config.Load().BookingService)
	if err != nil {
		log.Println(err)
	}
	time.Sleep(15 * time.Second)

	conn, err := rmq.Dial("amqp://guest:guest@rabbitmq:5672/")
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %v", err)
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		log.Printf("Failed to open a channel: %v", err)
	}
	defer ch.Close()

	IStorage := repo.NewStorage()
	servis := grpc.NewServer()

	bookingService := service.NewBookingService(IStorage)
	pb.RegisterBookingServer(servis, bookingService)

	crt, err := get(ch, "create_booking")
	if err != nil {
		log.Println(err)
	}

	del, err := get(ch, "booking_cancelled")
	if err != nil {
		log.Println(err)
	}
	crtMsgs, err := getMessage(ch, crt)
	if err != nil {
		log.Println(err)
	}
	delMsgs, err := getMessage(ch, del)
	if err != nil {
		log.Println(err)
	}

	db, err := Mongo.NewMongoClient(config.Load())
	if err != nil {
		log.Println(err)
	}

	res := messages.New(bookingService, ch, crtMsgs, delMsgs, &sync.WaitGroup{}, 6, db)
	go res.StartToConsume(context.Background())

	fmt.Printf("Server is listening on port %s\n", config.Load().BookingService)
	if err = servis.Serve(listener); err != nil {
		log.Println(err)
	}

}

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}

func get(ch *rmq.Channel, queueName string) (rmq.Queue, error) {
	return ch.QueueDeclare(
		queueName, // name
		false,     // durable
		false,     // delete when unused
		false,     // exclusive
		false,     // no-wait
		nil,       // arguments
	)
}

func getMessage(ch *rmq.Channel, q rmq.Queue) (<-chan rmq.Delivery, error) {
	return ch.Consume(
		q.Name, // queue
		"",     // consumer
		false,  // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
}

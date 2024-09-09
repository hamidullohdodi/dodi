package main

import (
	"context"
	"fmt"
	amqp "github.com/rabbitmq/amqp091-go"
	"google.golang.org/grpc"
	"log"
	"log/slog"
	"net"
	"sync"
	"twitt-service/genproto/tweet"
	"twitt-service/pkg/config"
	"twitt-service/pkg/logger"
	messagebroker "twitt-service/pkg/rebbitmq"
	"twitt-service/service"
	"twitt-service/storage/postgres"
)

func main() {
	logger := logger.InitLogger()
	cfg := config.Load()

	db, err := postgres.ConnectPostgres(cfg)
	if err != nil {
		logger.Error("Error connecting to database")
		log.Fatal(err)
	}

	conn, ch, err := initRabbitMQ(logger)
	if err != nil {
		log.Fatalf("Failed to initialize RabbitMQ: %v", err)
	}
	defer conn.Close()
	defer ch.Close()

	queues, err := initQueues(ch, logger)
	if err != nil {
		log.Fatalf("Failed to initialize queues: %v", err)
	}

	tweetSt := postgres.NewTweetRepo(db)
	tweetSt1 := postgres.NewCommentRepo(db)
	tweetSt2 := postgres.NewLikeRepo(db)
	tweetSr := service.NewTweetService(tweetSt, tweetSt2, tweetSt1, logger)

	listen, err := net.Listen("tcp", cfg.TWITT_SERVICE)
	fmt.Println("Listening on " + cfg.TWITT_SERVICE)
	if err != nil {
		logger.Error("Error listening on port "+cfg.TWITT_SERVICE, "error", err)
		log.Fatal(err)
	}

	res := messagebroker.New(tweetSr, ch, queues["PostComment"], queues["UpdateComment"], queues["PostTweet"], queues["AddLike"], queues["UpdateTweet"], &sync.WaitGroup{}, 5, db)
	go res.StartToConsume(context.Background())

	server := grpc.NewServer()
	tweet.RegisterTweetServiceServer(
		server,
		tweetSr,
	)

	if err := server.Serve(listen); err != nil {
		logger.Error("Error starting server on port "+cfg.TWITT_SERVICE, "error", err)
		log.Fatal(err)
	}
}

func initRabbitMQ(logger *slog.Logger) (*amqp.Connection, *amqp.Channel, error) {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		logger.Error("Failed to connect to RabbitMQ", "error", err)
		return nil, nil, err
	}

	ch, err := conn.Channel()
	if err != nil {
		logger.Error("Failed to open a channel", "error", err)
		return nil, nil, err
	}

	return conn, ch, nil
}

func initQueues(ch *amqp.Channel, logger *slog.Logger) (map[string]<-chan amqp.Delivery, error) {
	queueNames := []string{"PostComment", "UpdateComment", "PostTweet", "UpdateTweet", "AddLike"}
	queues := make(map[string]<-chan amqp.Delivery)

	for _, name := range queueNames {
		q, err := getQueue(ch, name)
		if err != nil {
			logger.Error("Failed to declare queue: "+name, "error", err)
			return nil, err
		}

		rc, err := getMessageQueue(ch, q)
		if err != nil {
			logger.Error("Failed to consume messages from queue: "+name, "error", err)
			return nil, err
		}

		queues[name] = rc
	}

	return queues, nil
}

func getQueue(ch *amqp.Channel, queueName string) (amqp.Queue, error) {
	return ch.QueueDeclare(
		queueName,
		false,
		false,
		false,
		false,
		nil,
	)
}

func getMessageQueue(ch *amqp.Channel, q amqp.Queue) (<-chan amqp.Delivery, error) {
	return ch.Consume(
		q.Name,
		"",
		false,
		false,
		false,
		false,
		nil,
	)
}

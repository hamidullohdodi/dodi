package main

import (
	"budgeting/config"
	pb "budgeting/genproto/account"
	pbb "budgeting/genproto/budget"
	pbp "budgeting/genproto/category"
	pbbp "budgeting/genproto/goal"
	pn "budgeting/genproto/notification"
	ptp "budgeting/genproto/transaction"
	logger "budgeting/log"
	messagebroker "budgeting/rabbitmq"
	"budgeting/service"
	"budgeting/storage/mongodb"
	"budgeting/storage/mongodb/account"
	"budgeting/storage/mongodb/budget"
	"budgeting/storage/mongodb/category"
	"budgeting/storage/mongodb/goal"
	"budgeting/storage/mongodb/notification"
	transac "budgeting/storage/mongodb/transaction"
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"sync"
	"syscall"

	amqp "github.com/rabbitmq/amqp091-go"
	"google.golang.org/grpc"
)

func main() {
	db, err := mongodb.Connect(context.Background())
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}

	cfg, err := config.Load(".")
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}
	log.Printf("Configuration: %v", cfg)

	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", cfg.BudgetingPort))

	if err != nil {
		log.Fatalf("Failed to listen on port %s: %v", cfg.BudgetingPort, err)
	}
	defer lis.Close()

	conn, err := amqp.Dial("amqp://guest:guest@rabbitmq:5672/")
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %v", err)
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("Failed to open a channel: %v", err)
	}
	defer ch.Close()

	queues := []string{"CreateTransaction", "UpdateBudget", "UpdateGoal"}
	messageQueues := make(map[string]<-chan amqp.Delivery)
	for _, queueName := range queues {
		q, err := getQueue(ch, queueName)
		if err != nil {
			log.Fatalf("Failed to declare %s queue: %v", queueName, err)
		}

		msgQueue, err := getMessageQueue(ch, q)
		if err != nil {
			log.Fatalf("Failed to consume %s messages: %v", queueName, err)
		}

		messageQueues[queueName] = msgQueue
	}

	logg := logger.InitLogger()

	storageA := account.NewAccountRepo(db, logg)
	serversA := service.NewAccountIService(storageA, logg)

	storageB := budget.NewBudgetRepo(db, logg)
	serversB := service.NewBudgetService(storageB, logg)

	storageC := category.NewCategoryRepo(db, logg)
	serversC := service.NewCategoryIService(storageC, logg)

	storageG := goal.NewGoalRepo(db, logg)
	serversG := service.NewGoalService(storageG, logg)

	storageT := transac.NewTransactionRepo(db, logg)
	serversT := service.NewTransactionIService(storageT, logg)

	storageDBR := notification.NewAccountRepoI(db, logg)
	serversDBR := service.NewNotificationServiceRepo(storageDBR, logg)

	res := messagebroker.New(serversT, serversG, serversB, ch, messageQueues["CreateTransaction"], messageQueues["UpdateBudget"], messageQueues["UpdateGoal"], &sync.WaitGroup{}, 3, db)
	go res.StartToConsume(context.Background())

	server := grpc.NewServer()
	pn.RegisterNotificationServiceServer(server, serversDBR)
	pb.RegisterAccountServiceServer(server, serversA)
	pbb.RegisterBudgetServiceServer(server, serversB)
	pbp.RegisterCategoryServiceServer(server, serversC)
	pbbp.RegisterGoalServiceServer(server, serversG)
	ptp.RegisterTransactionServiceServer(server, serversT)

	// Graceful shutdown
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-stop
		log.Println("Shutting down gracefully...")
		server.GracefulStop()
		// Add any additional cleanup code here if necessary
	}()

	log.Println("Starting server on", "localhost"+cfg.BudgetingPort)
	if err := server.Serve(lis); err != nil && err != grpc.ErrServerStopped {
		log.Fatalf("Failed to serve: %v", err)
	}
}

func getQueue(ch *amqp.Channel, queueName string) (amqp.Queue, error) {
	return ch.QueueDeclare(
		queueName, // name
		false,     // durable
		false,     // delete when unused
		false,     // exclusive
		false,     // no-wait
		nil,       // arguments
	)
}

func getMessageQueue(ch *amqp.Channel, q amqp.Queue) (<-chan amqp.Delivery, error) {
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

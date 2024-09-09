package rebbitmq

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/jmoiron/sqlx"
	amqp "github.com/rabbitmq/amqp091-go"
	"log/slog"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"twitt-service/genproto/tweet"
	"twitt-service/pkg/logger"
	"twitt-service/service"
)

type MsgBroker struct {
	service          *service.TweetService
	channel          *amqp.Channel
	PostComment      <-chan amqp.Delivery
	UpdateComment    <-chan amqp.Delivery
	AddLike          <-chan amqp.Delivery
	PostTweet        <-chan amqp.Delivery
	UpdateTweet      <-chan amqp.Delivery
	logger           *slog.Logger
	wg               *sync.WaitGroup
	numberOfServices int
	Db               *sqlx.DB
}

func New(
	service *service.TweetService,
	channel *amqp.Channel,
	PostComment <-chan amqp.Delivery,
	UpdateComment <-chan amqp.Delivery,
	AddLike <-chan amqp.Delivery,
	PostTweet <-chan amqp.Delivery,
	UpdateTweet <-chan amqp.Delivery,
	wg *sync.WaitGroup,
	numberOfServices int,
	Db *sqlx.DB) *MsgBroker {
	return &MsgBroker{
		service:          service,
		channel:          channel,
		PostComment:      PostComment,
		UpdateComment:    UpdateComment,
		AddLike:          AddLike,
		PostTweet:        PostTweet,
		UpdateTweet:      UpdateTweet,
		logger:           logger.InitLogger(),
		wg:               wg,
		numberOfServices: numberOfServices,
		Db:               Db,
	}
}

func (m *MsgBroker) StartToConsume(ctx context.Context) {
	m.wg.Add(m.numberOfServices)
	consumerCtx, cancel := context.WithCancel(ctx)
	defer cancel()

	go m.consumeMessages(consumerCtx, m.PostComment, "PostComment")
	go m.consumeMessages(consumerCtx, m.UpdateComment, "UpdateComment")
	go m.consumeMessages(consumerCtx, m.AddLike, "AddLike")
	go m.consumeMessages(consumerCtx, m.PostTweet, "PostTweet")
	go m.consumeMessages(consumerCtx, m.UpdateTweet, "UpdateTweet")

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	<-c

	m.logger.Info("Shutting down, waiting for consumers to finish")
	cancel()
	m.wg.Wait()
	m.logger.Info("All consumers have stopped")
}

func (m *MsgBroker) consumeMessages(ctx context.Context, messages <-chan amqp.Delivery, logPrefix string) {
	defer m.wg.Done()
	for {
		select {
		case val, ok := <-messages:
			if !ok {
				m.logger.Info("Message channel closed", "consumer", logPrefix)
				return
			}

			var err error
			switch logPrefix {
			case "PostComment":
				err = m.handlePostComment(ctx, val)
			case "UpdateComment":
				err = m.handleUpdateComment(ctx, val)
			case "PostTweet":
				err = m.handlePostTweet(ctx, val)
			case "UpdateTweet":
				err = m.handleUpdateTweet(ctx, val)
			case "AddLike":
				err = m.handleAddLike(ctx, val)
			}

			if err != nil {
				m.logger.Error(fmt.Sprintf("Failed in %s: %v", logPrefix, err))
				val.Nack(false, false)
				continue
			}

			val.Ack(false)

		case <-ctx.Done():
			m.logger.Info("Context done, stopping consumer", "consumer", logPrefix)
			return
		}
	}
}

func (m *MsgBroker) handlePostComment(ctx context.Context, msg amqp.Delivery) error {
	var req tweet.Comment
	if err := json.Unmarshal(msg.Body, &req); err != nil {
		return fmt.Errorf("error while unmarshaling PostComment data: %v", err)
	}
	fmt.Printf("%+v\n", req)
	fmt.Println("---------")
	_, err := m.service.PostComment(ctx, &req)
	return err
}

func (m *MsgBroker) handleUpdateComment(ctx context.Context, msg amqp.Delivery) error {
	var req tweet.UpdateAComment
	if err := json.Unmarshal(msg.Body, &req); err != nil {
		return fmt.Errorf("error while unmarshaling UpdateComment data: %v", err)
	}
	fmt.Println(msg.Body, req)
	fmt.Println("++++++++")
	_, err := m.service.UpdateComment(ctx, &req)
	return err
}

func (m *MsgBroker) handlePostTweet(ctx context.Context, msg amqp.Delivery) error {
	var req tweet.LikeReq
	if err := json.Unmarshal(msg.Body, &req); err != nil {
		return fmt.Errorf("error while unmarshaling AddLike data: %v", err)
	}

	fmt.Println(msg.Body, &req)
	fmt.Println("=========")
	fmt.Println(string(msg.Body), req)
	r, err := m.service.AddLike(ctx, &req)
	fmt.Println(r, err)
	return err
}

func (m *MsgBroker) handleUpdateTweet(ctx context.Context, msg amqp.Delivery) error {
	var req tweet.UpdateATweet
	if err := json.Unmarshal(msg.Body, &req); err != nil {
		return fmt.Errorf("error while unmarshaling UpdateTweet data: %v", err)
	}
	fmt.Println(msg.Body, req)
	fmt.Println("aaaaaaaaa")
	_, err := m.service.UpdateTweet(ctx, &req)
	return err
}

func (m *MsgBroker) handleAddLike(ctx context.Context, msg amqp.Delivery) error {
	var req tweet.Tweet
	if err := json.Unmarshal(msg.Body, &req); err != nil {
		return fmt.Errorf("error while unmarshaling PostTweet data: %v", err)
	}
	fmt.Println(msg.Body, &req)
	fmt.Println("ddddddd")
	r, err := m.service.PostTweet(ctx, &req)
	fmt.Println(r, err)
	return err
}

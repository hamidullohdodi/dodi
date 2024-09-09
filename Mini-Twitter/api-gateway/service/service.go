package service

import (
	pbb "apigateway/genproto/tweet"
	pb "apigateway/genproto/user"
	"apigateway/pkg/config"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Service interface {
	UserService() pb.UserServiceClient
	TweetService() pbb.TweetServiceClient
}

type service struct {
	userClient  pb.UserServiceClient
	tweetClient pbb.TweetServiceClient
}

func (s *service) UserService() pb.UserServiceClient {
	return s.userClient
}

func (s *service) TweetService() pbb.TweetServiceClient {
	return s.tweetClient
}

func NewService(cfg *config.Config) (Service, error) {
	userConn, err := grpc.NewClient("localhost:50050", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	twitter, err := grpc.NewClient("localhost:8088", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	return &service{
		userClient:  pb.NewUserServiceClient(userConn),
		tweetClient: pbb.NewTweetServiceClient(twitter),
	}, nil

}

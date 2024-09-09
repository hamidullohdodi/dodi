package service

import (
	"log/slog"
	pb "twitt-service/genproto/tweet"
	"twitt-service/storage"
)

type TweetService struct {
	tweet    storage.TweetStorage
	like     storage.LikesStorage
	comments storage.CommentsStorage
	logger   *slog.Logger
	pb.UnimplementedTweetServiceServer
}

func NewTweetService(st storage.TweetStorage, sl storage.LikesStorage, l storage.CommentsStorage, logger *slog.Logger) *TweetService {
	return &TweetService{
		tweet:    st,
		like:     sl,
		comments: l,
		logger:   logger,
	}
}

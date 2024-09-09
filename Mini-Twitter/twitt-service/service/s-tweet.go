package service

import (
	"context"
	"log"
	pb "twitt-service/genproto/tweet"
)

func (s *TweetService) PostTweet(ctx context.Context, in *pb.Tweet) (*pb.TweetResponse, error) {
	res, err := s.tweet.PostTweet(in)
	if err != nil {
		s.logger.Error("failed to post tweet", err)
		return nil, err
	}
	return res, nil
}

func (s *TweetService) UpdateTweet(ctx context.Context, in *pb.UpdateATweet) (*pb.TweetResponse, error) {
	res, err := s.tweet.UpdateTweet(in)
	if err != nil {
		s.logger.Error("failed to update tweet", err)
		return nil, err
	}
	return res, nil
}

func (s *TweetService) AddImageToTweet(ctx context.Context, in *pb.Url) (*pb.Message, error) {
	res, err := s.tweet.AddImageToTweet(in)
	if err != nil {
		s.logger.Error("failed to add image to tweet", err)
		return nil, err
	}
	return res, nil
}

func (s *TweetService) UserTweets(ctx context.Context, in *pb.UserId) (*pb.Tweets, error) {
	res, err := s.tweet.UserTweets(in)
	if err != nil {
		s.logger.Error("failed to user tweets", err)
		return nil, err
	}
	return res, nil
}

func (s *TweetService) GetTweet(ctx context.Context, in *pb.TweetId) (*pb.TweetResponse, error) {
	res, err := s.tweet.GetTweet(in)
	if err != nil {
		s.logger.Error("failed to get tweet", err)
		return nil, err
	}
	return res, nil
}

func (s *TweetService) GetAllTweets(ctx context.Context, in *pb.TweetFilter) (*pb.Tweets, error) {
	res, err := s.tweet.GetAllTweets(in)
	if err != nil {
		s.logger.Error("failed to get tweets", err)
		return nil, err
	}
	return res, nil
}

func (s *TweetService) RecommendTweets(ctx context.Context, in *pb.UserId) (*pb.Tweets, error) {
	res, err := s.tweet.RecommendTweets(in)
	if err != nil {
		s.logger.Error("failed to recommend tweets", err)
		return nil, err
	}
	return res, nil
}

func (s *TweetService) GetNewTweets(ctx context.Context, in *pb.UserId) (*pb.Tweets, error) {
	res, err := s.tweet.GetNewTweets(in)
	if err != nil {
		s.logger.Error("failed to get tweets", err)
		return nil, err
	}
	return res, nil
}
func (s *TweetService) ReTweet(ctx context.Context, in *pb.ReTweetReq) (*pb.TweetResponse, error) {
	res, err := s.tweet.AddReTweet(in)
	if err != nil {
		s.logger.Error("failed to get tweets", err)
		return nil, err
	}
	log.Printf("%+v", res)
	return res, nil
}

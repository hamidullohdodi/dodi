package service

import (
	"context"
	pb "twitt-service/genproto/tweet"
)

func (s *TweetService) AddLike(ctx context.Context, in *pb.LikeReq) (*pb.LikeRes, error) {
	res, err := s.like.AddLike(in)
	if err != nil {
		s.logger.Error("failed to create liked tweet", err)
		return nil, err
	}
	return res, nil
}

func (s *TweetService) DeleteLike(ctx context.Context, in *pb.LikeReq) (*pb.DLikeRes, error) {
	res, err := s.like.DeleteLIke(in)
	if err != nil {
		s.logger.Error("failed to delete liked tweet", err)
		return nil, err
	}
	return res, nil
}

func (s *TweetService) GetUserLikes(ctx context.Context, in *pb.UserId) (*pb.TweetTitles, error) {
	res, err := s.like.GetUserLikes(in)
	if err != nil {
		s.logger.Error("failed to get liked tweets", err)
		return nil, err
	}
	return res, nil
}

func (s *TweetService) GetCountTweetLikes(ctx context.Context, in *pb.TweetId) (*pb.Count, error) {
	res, err := s.like.GetCountTweetLikes(in)
	if err != nil {
		s.logger.Error("failed to get liked tweets", err)
		return nil, err
	}
	return res, nil
}

func (s *TweetService) MostLikedTweets(ctx context.Context, in *pb.Void) (*pb.TweetResponse, error) {
	res, err := s.like.MostLikedTweets(in)
	if err != nil {
		s.logger.Error("failed to get liked tweets", err)
		return nil, err
	}
	return res, nil
}

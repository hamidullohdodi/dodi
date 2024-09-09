package service

import (
	"context"
	pb "twitt-service/genproto/tweet"
)

func (s *TweetService) PostComment(ctx context.Context, in *pb.Comment) (*pb.CommentRes, error) {
	res, err := s.comments.PostComment(in)
	if err != nil {
		s.logger.Error("failed to post comment", err)
		return nil, err
	}
	return res, nil
}

func (s *TweetService) UpdateComment(ctx context.Context, in *pb.UpdateAComment) (*pb.CommentRes, error) {
	res, err := s.comments.UpdateComment(in)
	if err != nil {
		s.logger.Error("failed to update comment", err)
		return nil, err
	}
	return res, nil
}

func (s *TweetService) DeleteComment(ctx context.Context, in *pb.CommentId) (*pb.Message, error) {
	res, err := s.comments.DeleteComment(in)
	if err != nil {
		s.logger.Error("failed to delete comment", err)
		return nil, err
	}
	return res, nil
}

func (s *TweetService) GetComment(ctx context.Context, in *pb.CommentId) (*pb.Comment, error) {
	res, err := s.comments.GetComment(in)
	if err != nil {
		s.logger.Error("failed to get comment", err)
		return nil, err
	}
	return res, nil
}

func (s *TweetService) GetAllComments(ctx context.Context, in *pb.CommentFilter) (*pb.Comments, error) {
	res, err := s.comments.GetAllComments(in)
	if err != nil {
		s.logger.Error("failed to get comments", err)
		return nil, err
	}
	return res, nil
}

func (s *TweetService) GetUserComments(ctx context.Context, in *pb.UserId) (*pb.Comments, error) {
	res, err := s.comments.GetUserComments(in)
	if err != nil {
		s.logger.Error("failed to get comments", err)
		return nil, err
	}
	return res, nil
}

func (s *TweetService) AddLikeToComment(ctx context.Context, in *pb.CommentLikeReq) (*pb.Message, error) {
	res, err := s.comments.AddLikeToComment(in)
	if err != nil {
		s.logger.Error("failed to add like to comment", err)
		return nil, err
	}
	return res, nil
}

func (s *TweetService) DeleteLikeComment(ctx context.Context, in *pb.CommentLikeReq) (*pb.Message, error) {
	res, err := s.comments.DeleteLikeComment(in)
	if err != nil {
		s.logger.Error("failed to delete like comment", err)
		return nil, err
	}
	return res, nil
}

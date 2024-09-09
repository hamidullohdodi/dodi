package storage

import (
	pb "twitt-service/genproto/tweet"
)

type TweetStorage interface {
	PostTweet(in *pb.Tweet) (*pb.TweetResponse, error)
	UpdateTweet(in *pb.UpdateATweet) (*pb.TweetResponse, error)
	AddImageToTweet(in *pb.Url) (*pb.Message, error)
	UserTweets(in *pb.UserId) (*pb.Tweets, error)
	GetTweet(in *pb.TweetId) (*pb.TweetResponse, error)
	GetAllTweets(in *pb.TweetFilter) (*pb.Tweets, error)
	RecommendTweets(in *pb.UserId) (*pb.Tweets, error)
	GetNewTweets(in *pb.UserId) (*pb.Tweets, error)
	AddReTweet(in *pb.ReTweetReq) (*pb.TweetResponse, error)
}

type LikesStorage interface {
	AddLike(in *pb.LikeReq) (*pb.LikeRes, error)
	DeleteLIke(in *pb.LikeReq) (*pb.DLikeRes, error)
	GetUserLikes(in *pb.UserId) (*pb.TweetTitles, error)
	GetCountTweetLikes(in *pb.TweetId) (*pb.Count, error)
	MostLikedTweets(in *pb.Void) (*pb.TweetResponse, error)
}

type CommentsStorage interface {
	PostComment(in *pb.Comment) (*pb.CommentRes, error)
	UpdateComment(in *pb.UpdateAComment) (*pb.CommentRes, error)
	DeleteComment(in *pb.CommentId) (*pb.Message, error)
	GetComment(in *pb.CommentId) (*pb.Comment, error)
	GetAllComments(in *pb.CommentFilter) (*pb.Comments, error)
	GetUserComments(in *pb.UserId) (*pb.Comments, error)
	AddLikeToComment(in *pb.CommentLikeReq) (*pb.Message, error)
	DeleteLikeComment(in *pb.CommentLikeReq) (*pb.Message, error)
}

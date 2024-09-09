package postgres

import (
	"fmt"
	"testing"
	pb "twitt-service/genproto/tweet"
)

func TestAddLike(t *testing.T) {
	db, err := ConnectTweet()
	if err != nil {
		t.Errorf("Error connecting to database: %v", err)
	}
	res := pb.LikeReq{
		UserId:  "5226669a-28e0-4382-952b-35230d233473",
		TweetId: "b8f7677e-1905-400c-ba66-24f44abb21b4",
	}

	tweet := NewLikeRepo(db)

	req, err := tweet.AddLike(&res)
	if err != nil {
		t.Errorf("Error adding like: %v", err)
	}
	fmt.Println(req)
}

func TestGetUserLikes(t *testing.T) {
	db, err := ConnectTweet()
	if err != nil {
		t.Errorf("Error connecting to database: %v", err)
	}
	res := pb.UserId{
		Id: "5226669a-28e0-4382-952b-35230d233473",
	}

	tweet := NewLikeRepo(db)

	req, err := tweet.GetUserLikes(&res)
	if err != nil {
		t.Errorf("Error getting likes: %v", err)
	}
	fmt.Println(req)
}

func TestGetCountTweetLikes(t *testing.T) {
	db, err := ConnectTweet()
	if err != nil {
		t.Errorf("Error connecting to database: %v", err)
	}

	res := pb.TweetId{
		Id: "b8f7677e-1905-400c-ba66-24f44abb21b4",
	}

	tweet := NewLikeRepo(db)

	rep, err := tweet.GetCountTweetLikes(&res)
	if err != nil {
		t.Errorf("Error getting likes: %v", err)
	}
	fmt.Println(rep)
}

func TestMostLikedTweets(t *testing.T) {
	db, err := ConnectTweet()
	if err != nil {
		t.Errorf("Error connecting to database: %v", err)
	}

	res := pb.Void{}

	tweet := NewLikeRepo(db)

	rep, err := tweet.MostLikedTweets(&res)
	if err != nil {
		t.Errorf("Error getting likes: %v", err)
	}
	fmt.Println(rep)
}

func TestDeleteLike(t *testing.T) {
	db, err := ConnectTweet()
	if err != nil {
		t.Errorf("Error connecting to database: %v", err)
	}

	res := pb.LikeReq{
		UserId:  "5226669a-28e0-4382-952b-35230d233473",
		TweetId: "b8f7677e-1905-400c-ba66-24f44abb21b4",
	}

	tweet := NewLikeRepo(db)

	req, err := tweet.DeleteLIke(&res)
	if err != nil {
		t.Errorf("Error deleting like: %v", err)
	}

	fmt.Println(req)
}

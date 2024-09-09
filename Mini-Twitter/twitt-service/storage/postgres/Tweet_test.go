package postgres

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"testing"
	pb "twitt-service/genproto/tweet"
)

func ConnectTweet() (*sqlx.DB, error) {
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		"localhost", "5432", "postgres", "dodi", "tweet")

	db, err := sqlx.Connect("postgres", psqlInfo)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func TestPostTweet(t *testing.T) {
	db, err := ConnectTweet()
	if err != nil {
		t.Fatal(err)
	}
	res := pb.Tweet{
		UserId:   uuid.New().String(),
		Hashtag:  "...",
		Title:    "....",
		Content:  "...",
		ImageUrl: "...",
	}

	tweet := NewTweetRepo(db)

	req, err := tweet.PostTweet(&res)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(req)

}

func TestUpdateTweet(t *testing.T) {
	db, err := ConnectTweet()
	if err != nil {
		t.Fatal(err)
	}

	res := pb.UpdateATweet{
		Id:      "b8f7677e-1905-400c-ba66-24f44abb21b4",
		Hashtag: "...",
		Title:   "....",
		Content: "...",
	}

	tweet := NewTweetRepo(db)

	req, err := tweet.UpdateTweet(&res)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(req)
}

func TestAddImageToTweet(t *testing.T) {
	db, err := ConnectTweet()
	if err != nil {
		t.Fatal(err)
	}

	res := pb.Url{
		TweetId: "b8f7677e-1905-400c-ba66-24f44abb21b4",
		Url:     "nimadur",
	}

	tweet := NewTweetRepo(db)

	req, err := tweet.AddImageToTweet(&res)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(req)
}

func TestUserTweets(t *testing.T) {
	db, err := ConnectTweet()
	if err != nil {
		t.Fatal(err)
	}

	res := pb.UserId{
		Id: "5226669a-28e0-4382-952b-35230d233473",
	}

	tweet := NewTweetRepo(db)

	req, err := tweet.UserTweets(&res)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(req)
}

func TestGetTweet(t *testing.T) {
	db, err := ConnectTweet()
	if err != nil {
		t.Fatal(err)
	}

	res := pb.TweetId{
		Id: "b8f7677e-1905-400c-ba66-24f44abb21b4",
	}

	tweet := NewTweetRepo(db)

	req, err := tweet.GetTweet(&res)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(req)
}

func TestRecommendTweets(t *testing.T) {
	db, err := ConnectTweet()
	if err != nil {
		t.Fatal(err)
	}

	res := pb.UserId{
		Id: "5226669a-28e0-4382-952b-35230d233473",
	}

	tweet := NewTweetRepo(db)

	req, err := tweet.RecommendTweets(&res)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(req)
}

func TestGetNewTweets(t *testing.T) {
	db, err := ConnectTweet()
	if err != nil {
		t.Fatal(err)
	}

	res := pb.UserId{
		Id: "5226669a-28e0-4382-952b-35230d233473",
	}

	tweet := NewTweetRepo(db)

	req, err := tweet.GetNewTweets(&res)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(req)
}

package postgres

import (
	"fmt"
	"github.com/google/uuid"
	_ "github.com/lib/pq"
	"testing"
	pb "twitt-service/genproto/tweet"
)

func TestPostComment(t *testing.T) {
	db, err := ConnectTweet()
	if err != nil {
		t.Fatal(err)
	}
	res := pb.Comment{
		Id:        uuid.New().String(),
		UserId:    "5226669a-28e0-4382-952b-35230d233473",
		TweetId:   "b8f7677e-1905-400c-ba66-24f44abb21b4",
		Content:   "...",
		LikeCount: 70,
	}

	tweet := NewCommentRepo(db)

	req, err := tweet.PostComment(&res)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(req)

}

func TestUpdateComment(t *testing.T) {
	db, err := ConnectTweet()
	if err != nil {
		t.Fatal(err)
	}

	res := pb.UpdateAComment{
		Id:      "ebf86f8f-c094-4676-b257-2d00f6435639",
		Content: "111",
	}

	tweet := NewCommentRepo(db)

	req, err := tweet.UpdateComment(&res)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(req)
}

func TestGetComment(t *testing.T) {
	db, err := ConnectTweet()
	if err != nil {
		t.Fatal(err)
	}

	res := pb.CommentId{
		Id: "99b6b42e-ade2-4454-8ff8-9ad6ab5db05d",
	}

	tweet := NewCommentRepo(db)

	req, err := tweet.GetComment(&res)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(req)
}

func TestGetAllComments(t *testing.T) {
	db, err := ConnectTweet()
	if err != nil {
		t.Fatal(err)
	}

	res := pb.CommentFilter{
		UserId:  "5226669a-28e0-4382-952b-35230d233473",
		TweetId: "b8f7677e-1905-400c-ba66-24f44abb21b4",
	}

	tweet := NewCommentRepo(db)

	req, err := tweet.GetAllComments(&res)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(req)
}

func TestGetUserComments(t *testing.T) {
	db, err := ConnectTweet()
	if err != nil {
		t.Fatal(err)
	}
	res := pb.UserId{
		Id: "5226669a-28e0-4382-952b-35230d233473",
	}

	tweet := NewCommentRepo(db)

	req, err := tweet.GetUserComments(&res)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(req)
}

func TestAddLikeToComment(t *testing.T) {
	db, err := ConnectTweet()
	if err != nil {
		t.Fatal(err)
	}

	res := pb.CommentLikeReq{
		CommentId: "ebf86f8f-c094-4676-b257-2d00f6435639",
	}

	tweet := NewCommentRepo(db)

	req, err := tweet.AddLikeToComment(&res)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(req)
}

func TestDeleteComment(t *testing.T) {
	db, err := ConnectTweet()
	if err != nil {
		t.Fatal(err)
	}
	res := pb.CommentId{
		Id: "ebf86f8f-c094-4676-b257-2d00f6435639",
	}

	tweet := NewCommentRepo(db)

	req, err := tweet.DeleteComment(&res)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(req)
}

func TestDeleteLikeComment(t *testing.T) {
	db, err := ConnectTweet()
	if err != nil {
		t.Fatal(err)
	}
	res := pb.CommentLikeReq{
		CommentId: "99b6b42e-ade2-4454-8ff8-9ad6ab5db05d",
	}
	tweet := NewCommentRepo(db)

	req, err := tweet.DeleteLikeComment(&res)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(req)
}

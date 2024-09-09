package postgres

import (
	"context"
	"github.com/jmoiron/sqlx"
	pb "twitt-service/genproto/tweet"
)

type LikeRepo struct {
	db *sqlx.DB
}

func NewLikeRepo(db *sqlx.DB) *LikeRepo {
	return &LikeRepo{
		db: db,
	}
}

func (l *LikeRepo) AddLike(in *pb.LikeReq) (*pb.LikeRes, error) {
	query := `INSERT INTO likes (user_id, tweet_id, created_at) 
	          VALUES ($1, $2, NOW()) 
	          RETURNING user_id, tweet_id`

	var res pb.LikeRes
	err := l.db.QueryRowContext(context.Background(), query, in.UserId, in.TweetId).Scan(
		&res.UserId, &res.TweetId)

	if err != nil {
		return nil, err
	}

	return &res, nil
}

func (l *LikeRepo) DeleteLIke(in *pb.LikeReq) (*pb.DLikeRes, error) {
	query := `DELETE FROM likes WHERE user_id = $1 AND tweet_id = $2 
	          RETURNING user_id, tweet_id`

	var res pb.DLikeRes
	err := l.db.QueryRowContext(context.Background(), query, in.UserId, in.TweetId).Scan(
		&res.UserId, &res.TweetId)

	if err != nil {
		return nil, err
	}

	return &res, nil
}

func (l *LikeRepo) GetUserLikes(in *pb.UserId) (*pb.TweetTitles, error) {
	query := `SELECT t.title 
	          FROM tweets t 
	          JOIN likes l ON t.id = l.tweet_id 
	          WHERE l.user_id = $1`

	rows, err := l.db.QueryContext(context.Background(), query, in.Id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	titles := &pb.TweetTitles{}
	for rows.Next() {
		var title string
		if err := rows.Scan(&title); err != nil {
			return nil, err
		}
		titles.Titles = append(titles.Titles, title)
	}

	return titles, nil
}

func (l *LikeRepo) GetCountTweetLikes(in *pb.TweetId) (*pb.Count, error) {
	query := `SELECT COUNT(*) FROM likes WHERE tweet_id = $1`
	var count pb.Count
	err := l.db.QueryRowContext(context.Background(), query, in.Id).Scan(&count.Count)
	if err != nil {
		return nil, err
	}

	return &count, nil
}

func (l *LikeRepo) MostLikedTweets(in *pb.Void) (*pb.TweetResponse, error) {
	query := `SELECT t.id, t.user_id, t.title, t.content, t.image_url, t.created_at, COUNT(l.tweet_id) as like_count 
	          FROM tweets t 
	          JOIN likes l ON t.id = l.tweet_id 
	          GROUP BY t.id 
	          ORDER BY like_count DESC 
	          LIMIT 1`

	var tweet pb.TweetResponse
	err := l.db.QueryRowContext(context.Background(), query).Scan(
		&tweet.Id, &tweet.UserId, &tweet.Title, &tweet.Content, &tweet.ImageUrl, &tweet.CreatedAt, &tweet.LikeCount)

	if err != nil {
		return nil, err
	}

	return &tweet, nil
}

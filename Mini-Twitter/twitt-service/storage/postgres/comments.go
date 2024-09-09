package postgres

import (
	"github.com/jmoiron/sqlx"
	"time"
	pb "twitt-service/genproto/tweet"
	"twitt-service/storage"
)

type CommentRepo struct {
	db *sqlx.DB
}

func NewCommentRepo(db *sqlx.DB) storage.CommentsStorage {
	return &CommentRepo{db: db}
}

func (c *CommentRepo) PostComment(in *pb.Comment) (*pb.CommentRes, error) {
	query := `INSERT INTO comments (id, user_id, tweet_id, context, like_count, created_at, updated_at)
			  VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id`

	var id string
	err := c.db.QueryRow(query, in.Id, in.UserId, in.TweetId, in.Content, in.LikeCount, time.Now(), time.Now()).Scan(&id)
	if err != nil {
		return nil, err
	}

	return &pb.CommentRes{Id: id, UserId: in.UserId, TweetId: in.TweetId, Content: in.Content, LikeCount: in.LikeCount, CreatedAt: time.Now().String(), UpdatedAt: time.Now().String()}, nil
}

func (c *CommentRepo) UpdateComment(in *pb.UpdateAComment) (*pb.CommentRes, error) {
	query := `UPDATE comments SET context = $1, updated_at = $2 WHERE id = $3 RETURNING id`

	var id string
	err := c.db.QueryRow(query, in.Content, time.Now(), in.Id).Scan(&id)
	if err != nil {
		return nil, err
	}

	return &pb.CommentRes{Id: id, Content: in.Content, UpdatedAt: time.Now().Format(time.RFC3339)}, nil
}

func (c *CommentRepo) DeleteComment(in *pb.CommentId) (*pb.Message, error) {
	query := `DELETE FROM comments WHERE id = $1`

	_, err := c.db.Exec(query, in.Id)
	if err != nil {
		return nil, err
	}

	return &pb.Message{Message: "Comment deleted successfully"}, nil
}

func (c *CommentRepo) GetComment(in *pb.CommentId) (*pb.Comment, error) {
	query := `SELECT id, user_id, tweet_id, context, like_count FROM comments WHERE id = $1`

	var comment pb.Comment
	err := c.db.QueryRow(query, in.Id).Scan(
		&comment.Id, &comment.UserId, &comment.TweetId, &comment.Content, &comment.LikeCount)
	if err != nil {
		return nil, err
	}

	return &comment, nil
}

func (c *CommentRepo) GetAllComments(in *pb.CommentFilter) (*pb.Comments, error) {
	query := `SELECT id, user_id, tweet_id, context, like_count, created_at, updated_at FROM comments WHERE tweet_id = $1`

	rows, err := c.db.Query(query, in.TweetId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var comments pb.Comments
	for rows.Next() {
		var comment pb.CommentRes
		if err := rows.Scan(&comment.Id, &comment.UserId, &comment.TweetId, &comment.Content, &comment.LikeCount, &comment.CreatedAt, &comment.UpdatedAt); err != nil {
			return nil, err
		}
		comments.Comments = append(comments.Comments, &comment)
	}

	return &comments, nil
}

func (c *CommentRepo) GetUserComments(in *pb.UserId) (*pb.Comments, error) {
	query := `SELECT id, user_id, tweet_id, context, like_count, created_at, updated_at FROM comments WHERE user_id = $1`

	rows, err := c.db.Query(query, in.Id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var comments pb.Comments
	for rows.Next() {
		var comment pb.CommentRes
		if err := rows.Scan(&comment.Id, &comment.UserId, &comment.TweetId, &comment.Content, &comment.LikeCount, &comment.CreatedAt, &comment.UpdatedAt); err != nil {
			return nil, err
		}
		comments.Comments = append(comments.Comments, &comment)
	}

	return &comments, nil
}

func (c *CommentRepo) AddLikeToComment(in *pb.CommentLikeReq) (*pb.Message, error) {
	query := `UPDATE comments SET like_count = like_count + 1 WHERE id = $1`

	_, err := c.db.Exec(query, in.CommentId)
	if err != nil {
		return nil, err
	}

	return &pb.Message{Message: "Like added to comment"}, nil
}

func (c *CommentRepo) DeleteLikeComment(in *pb.CommentLikeReq) (*pb.Message, error) {
	query := `UPDATE comments SET like_count = like_count - 1 WHERE id = $1`

	_, err := c.db.Exec(query, in.CommentId)
	if err != nil {
		return nil, err
	}

	return &pb.Message{Message: "Like removed from comment"}, nil
}

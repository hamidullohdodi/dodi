package handler

import (
	pb "apigateway/genproto/tweet"
	"apigateway/pkg/models"
	t "apigateway/pkg/token"
	"apigateway/service"
	"context"
	"encoding/json"
	"github.com/gin-gonic/gin"
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
	"log/slog"
	"net/http"
)

type LikeHandler interface {
	AddLike(c *gin.Context)
	DeleteLIke(c *gin.Context)
	GetUserLikes(c *gin.Context)
	GetCountTweetLikes(c *gin.Context)
	MostLikedTweets(c *gin.Context)
}

type HandlerL struct {
	LikeMQ       *service.MsgBroker
	TweetService pb.TweetServiceClient
	logger       *slog.Logger
}

func NewLikeHandler(tweerService service.Service, logger *slog.Logger, conn *amqp.Channel) LikeHandler {
	tweerClient := tweerService.TweetService()
	if tweerClient == nil {
		log.Fatalf("Error creating like handler")
		return nil
	}
	return &HandlerL{
		LikeMQ:       service.NewMsgBroker(conn, logger),
		TweetService: tweerClient,
		logger:       logger,
	}
}

// AddLike godoc
// @Summary Add a like to a tweet
// @Description Add a like to a tweet by a user
// @Security BearerAuth
// @Tags Like
// @Accept json
// @Produce json
// @Param AddLike body models.LikeReq true "Add like request"
// @Success 200 {object} models.LikeRes
// @Failure 400 {object} models.Error
// @Failure 404 {object} models.Error
// @Failure 500 {object} models.Error
// @Router /like/add [post]
func (h *HandlerL) AddLike(c *gin.Context) {
	var like models.LikeReq
	if err := c.ShouldBindJSON(&like); err != nil {
		h.logger.Error("Error occurred while posting tweet", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	token := c.GetHeader("Authorization")
	cl, err := t.ExtractClaims(token)
	if err != nil {
		h.logger.Error("Error occurred while extracting claims", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	UserID := cl["user_id"].(string)

	res := pb.LikeReq{
		UserId:  UserID,
		TweetId: like.TweetID,
	}

	body, err := json.Marshal(&res)
	if err != nil {
		h.logger.Error("Error occurred while posting tweet", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = h.LikeMQ.AddLike(body)
	if err != nil {
		h.logger.Error("Error occurred while posting tweet", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"massage": "success"})
}

// DeleteLIke godoc
// @Summary Delete a like from a tweet
// @Description Remove a like from a tweet by a user
// @Security BearerAuth
// @Tags Like
// @Accept json
// @Produce json
// @Param tweet_id path string true "Tweet ID"
// @Success 200 {object} models.LikeRes
// @Failure 400 {object} models.Error
// @Failure 404 {object} models.Error
// @Failure 500 {object} models.Error
// @Router /like/delete/{tweet_id} [delete]
func (h *HandlerL) DeleteLIke(c *gin.Context) {
	tweetID := c.Param("tweet_id")

	res := pb.LikeReq{
		UserId:  c.MustGet("user_id").(string),
		TweetId: tweetID,
	}

	// Call the service to delete the like
	req, err := h.TweetService.DeleteLike(context.Background(), &res)
	if err != nil {
		h.logger.Error("Error occurred while deleting like", err)
		// Determine the appropriate status code based on the error
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": req})
}

// GetUserLikes godoc
// @Summary Get all likes by a user
// @Description Retrieve all tweets liked by a specific user
// @Security BearerAuth
// @Tags Like
// @Accept json
// @Produce json
// @Success 200 {object} models.TweetTitles
// @Failure 400 {object} models.Error
// @Failure 404 {object} models.Error
// @Failure 500 {object} models.Error
// @Router /like/get_user [get]
func (h *HandlerL) GetUserLikes(c *gin.Context) {
	res := pb.UserId{
		Id: c.MustGet("user_id").(string),
	}
	req, err := h.TweetService.GetUserLikes(context.Background(), &res)
	if err != nil {
		h.logger.Error("Error occurred while posting tweet", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": req})
}

// GetCountTweetLikes godoc
// @Summary Get the count of likes for a tweet
// @Description Retrieve the number of likes for a specific tweet
// @Security BearerAuth
// @Tags Like
// @Accept json
// @Produce json
// @Param tweet_id path string true "Tweet ID"
// @Success 200 {object} models.Count
// @Failure 400 {object} models.Error
// @Failure 404 {object} models.Error
// @Failure 500 {object} models.Error
// @Router /like/get_count/{tweet_id} [get]
func (h *HandlerL) GetCountTweetLikes(c *gin.Context) {
	TweetId := c.Param("tweet_id")

	res := pb.TweetId{
		Id: TweetId,
	}

	req, err := h.TweetService.GetCountTweetLikes(context.Background(), &res)
	if err != nil {
		h.logger.Error("Error occurred while posting tweet", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": req})
}

// MostLikedTweets godoc
// @Summary Get the most liked tweets
// @Description Retrieve tweets with the highest number of likes
// @Security BearerAuth
// @Tags Like
// @Accept json
// @Produce json
// @Success 200 {object} models.Tweet
// @Failure 400 {object} models.Error
// @Failure 404 {object} models.Error
// @Failure 500 {object} models.Error
// @Router /like/get_most [get]
func (h *HandlerL) MostLikedTweets(c *gin.Context) {
	res := pb.Void{}

	req, err := h.TweetService.MostLikedTweets(context.Background(), &res)
	if err != nil {
		h.logger.Error("Error occurred while posting tweet", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": req})
}

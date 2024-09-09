package handler

import (
	pb "apigateway/genproto/tweet"
	"apigateway/pkg/models"
	t "apigateway/pkg/token"
	"apigateway/service"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
	"log/slog"
	"net/http"
	"strconv"
)

type TweetHandler interface {
	PostTweet(c *gin.Context)
	UpdateTweet(c *gin.Context)
	AddImageToTweet(c *gin.Context)
	UserTweets(c *gin.Context)
	GetTweet(c *gin.Context)
	GetAllTweets(c *gin.Context)
	RecommendTweets(c *gin.Context)
	GetNewTweets(c *gin.Context)
	ReTweet(c *gin.Context)
}

type tweetHandler struct {
	TweetMQ      *service.MsgBroker
	tweetService pb.TweetServiceClient
	logger       *slog.Logger
}

func NewTweetHandler(tweetService service.Service, logger *slog.Logger, conn *amqp.Channel) TweetHandler {
	tweetClient := tweetService.TweetService()
	if tweetClient == nil {
		log.Fatalf("tweet client is nil")
		return nil
	}
	return &tweetHandler{
		TweetMQ:      service.NewMsgBroker(conn, logger),
		tweetService: tweetClient,
		logger:       logger,
	}
}

// PostTweet godoc
// @Summary PostTweet Tweets
// @Description Create a new tweet
// @Security BearerAuth
// @Tags Tweet
// @Accept json
// @Produce json
// @Param PostTweet body models.Tweet true "Post tweet"
// @Success 200 {object} models.TweetResponse
// @Failure 400 {object} models.Error
// @Failure 500 {object} models.Error
// @Router /tweet/add [post]
func (h *tweetHandler) PostTweet(c *gin.Context) {
	var tweet models.Tweet
	if err := c.ShouldBindJSON(&tweet); err != nil {
		h.logger.Error("Error occurred while binding json", err)
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

	tweet.UserID = cl["user_id"].(string)

	fmt.Println(tweet.UserID)

	res := pb.Tweet{
		UserId:   tweet.UserID,
		Hashtag:  tweet.Hashtag,
		Title:    tweet.Title,
		Content:  tweet.Content,
		ImageUrl: tweet.ImageURL,
	}

	bady, err := json.Marshal(&res)
	if err != nil {
		h.logger.Error("Error occurred while marshaling json", err)
		return
	}

	log.Println(string(bady))

	err = h.TweetMQ.PostTweet(bady)
	if err != nil {
		h.logger.Error("Error occurred while posting tweet", err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"massage": "success"})
}

// UpdateTweet godoc
// @Summary UpdateTweet Tweets
// @Description Update an existing tweet
// @Security BearerAuth
// @Tags Tweet
// @Accept json
// @Produce json
// @Param UpdateTweet body models.UpdateATweet true "Update tweet"
// @Success 200 {object} models.TweetResponse
// @Failure 400 {object} models.Error
// @Failure 500 {object} models.Error
// @Router /tweet/add_up [put]
func (h *tweetHandler) UpdateTweet(c *gin.Context) {
	var tweet models.UpdateATweet
	if err := c.ShouldBindJSON(&tweet); err != nil {
		h.logger.Error("Error occurred while binding json", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	res := pb.UpdateATweet{
		Id:      tweet.ID,
		Hashtag: tweet.Hashtag,
		Title:   tweet.Title,
		Content: tweet.Content,
	}
	bady, err := json.Marshal(&res)
	if err != nil {
		h.logger.Error("Error occurred while marshaling json", err)
		return
	}
	err = h.TweetMQ.UpdateTweet(bady)
	if err != nil {
		h.logger.Error("Error occurred while updating tweet", err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"massage": "success"})
}

// AddImageToTweet godoc
// @Summary AddImageToTweet Tweets
// @Description Add an image to a tweet
// @Security BearerAuth
// @Tags Tweet
// @Accept json
// @Produce json
// @Param AddImageToTweet body models.Url true "Add image to tweet"
// @Success 200 {object} models.Message
// @Failure 400 {object} models.Error
// @Failure 500 {object} models.Error
// @Router /tweet/add_image [post]
func (h *tweetHandler) AddImageToTweet(c *gin.Context) {
	var tweet models.Url
	if err := c.ShouldBindJSON(&tweet); err != nil {
		h.logger.Error("Error occurred while binding json", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	res := pb.Url{
		TweetId: tweet.TweetID,
		Url:     tweet.URL,
	}

	req, err := h.tweetService.AddImageToTweet(c.Request.Context(), &res)
	if err != nil {
		h.logger.Error("Error occurred while adding image to tweet", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": req})
}

// UserTweets godoc
// @Summary UserTweets Tweets
// @Description Get all tweets for a specific user
// @Security BearerAuth
// @Tags Tweet
// @Accept json
// @Produce json
// @Param user_id query string false "User ID"
// @Success 200 {object} models.Tweets
// @Failure 400 {object} models.Error
// @Failure 500 {object} models.Error
// @Router /tweet/user [get]
func (h *tweetHandler) UserTweets(c *gin.Context) {
	res := pb.UserId{
		Id: c.MustGet("user_id").(string),
	}

	req, err := h.tweetService.UserTweets(c.Request.Context(), &res)
	if err != nil {
		h.logger.Error("Error occurred while getting user tweets", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": req})
}

// GetTweet godoc
// @Summary GetTweet Tweets
// @Description Get details of a specific tweet
// @Security BearerAuth
// @Tags Tweet
// @Accept json
// @Produce json
// @Param id path string true "Tweet ID"
// @Success 200 {object} models.TweetResponse
// @Failure 400 {object} models.Error
// @Failure 500 {object} models.Error
// @Router /tweet/get_tt/{id} [get]
func (h *tweetHandler) GetTweet(c *gin.Context) {
	id := c.Param("id")

	res := pb.TweetId{
		Id: id,
	}

	req, err := h.tweetService.GetTweet(c.Request.Context(), &res)
	if err != nil {
		h.logger.Error("Error occurred while getting tweet", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": req})
}

// GetAllTweets godoc
// @Summary GetAllTweets Tweets
// @Description Get a list of all tweets with optional filters
// @Security BearerAuth
// @Tags Tweet
// @Accept json
// @Produce json
// @Param limit query int false "Limit of tweets"
// @Param offset query int false "Offset of tweets"
// @Param hashtag query string false "Hashtag filter"
// @Param title query string false "Title filter"
// @Success 200 {object} models.Tweets
// @Failure 400 {object} models.Error
// @Failure 500 {object} models.Error
// @Router /tweet/get_all [get]
func (h *tweetHandler) GetAllTweets(c *gin.Context) {
	limit := c.Query("limit")
	offset := c.Query("offset")

	offsets, err := strconv.Atoi(offset)
	if err != nil {
		offsets = 1
	}

	limits, err := strconv.Atoi(limit)
	if err != nil {
		limits = 10
	}

	hashtag := c.Query("hashtag")
	title := c.Query("title")

	res := pb.TweetFilter{
		Limit:   int64(limits),
		Offset:  int64(offsets),
		Hashtag: hashtag,
		Title:   title,
	}

	req, err := h.tweetService.GetAllTweets(c.Request.Context(), &res)
	if err != nil {
		h.logger.Error("Error occurred while getting tweets", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": req})
}

// RecommendTweets godoc
// @Summary RecommendTweets Tweets
// @Description Get tweet recommendations for a specific user
// @Security BearerAuth
// @Tags Tweet
// @Accept json
// @Produce json
// @Param user_id query string false "User ID"
// @Success 200 {object} models.Tweets
// @Failure 400 {object} models.Error
// @Failure 500 {object} models.Error
// @Router /tweet/recommend [get]
func (h *tweetHandler) RecommendTweets(c *gin.Context) {
	res := pb.UserId{
		Id: c.MustGet("user_id").(string),
	}
	req, err := h.tweetService.RecommendTweets(c.Request.Context(), &res)
	if err != nil {
		h.logger.Error("Error occurred while getting tweets", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": req})
}

// GetNewTweets godoc
// @Summary GetNewTweets Tweets
// @Description Get new tweets for a specific user
// @Security BearerAuth
// @Tags Tweet
// @Accept json
// @Produce json
// @Param user_id query string false "User ID"
// @Success 200 {object} models.Tweets
// @Failure 400 {object} models.Error
// @Failure 500 {object} models.Error
// @Router /tweet/get_new [get]
func (h *tweetHandler) GetNewTweets(c *gin.Context) {
	res := pb.UserId{
		Id: c.MustGet("user_id").(string),
	}
	req, err := h.tweetService.RecommendTweets(c.Request.Context(), &res)
	if err != nil {
		h.logger.Error("Error occurred while getting tweets", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": req})
}

// ReTweet godoc
// @Summary ReTweet Tweets
// @Description Retweet a tweet by a user
// @Security BearerAuth
// @Tags Tweet
// @Accept json
// @Produce json
// @Param ReTweet body tweet.ReTweetReq true "Post retweet"
// @Success 200 {object} tweet.TweetResponse
// @Failure 400 {object} models.Error
// @Failure 500 {object} models.Error
// @Router /tweet/re_tweet [post]
func (h *tweetHandler) ReTweet(c *gin.Context) {
	req := pb.ReTweetReq{}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	req.UserId = c.MustGet("user_id").(string)

	res, err := h.tweetService.ReTweet(c.Request.Context(), &req)
	if err != nil {
		log.Println(err)
		h.logger.Error("Error occurred while retweeting tweet", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": res})
}

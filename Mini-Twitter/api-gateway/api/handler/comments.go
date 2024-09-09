package handler

import (
	pb "apigateway/genproto/tweet"
	"apigateway/pkg/models"
	t "apigateway/pkg/token"
	"apigateway/service"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
	"log/slog"
	"net/http"
)

type CommentHandler interface {
	PostComment(c *gin.Context)
	UpdateComment(c *gin.Context)
	DeleteComment(c *gin.Context)
	GetComment(c *gin.Context)
	GetAllComments(c *gin.Context)
	GetUserComments(c *gin.Context)
	AddLikeToComment(c *gin.Context)
	DeleteLikeComment(c *gin.Context)
}

type commentHandler struct {
	CommentMQ      *service.MsgBroker
	CommentService pb.TweetServiceClient
	logger         *slog.Logger
}

func NewCommentHandler(commentService service.Service, logger *slog.Logger, conn *amqp.Channel) CommentHandler {
	commentClient := commentService.TweetService()
	if commentClient == nil {
		log.Fatalf("Failed to create comment handler")
		return nil
	}
	return &commentHandler{
		CommentMQ:      service.NewMsgBroker(conn, logger),
		CommentService: commentClient,
		logger:         logger,
	}
}

// PostComment godoc
// @Summary Post a new comment
// @Description Create a new comment for a tweet
// @Security BearerAuth
// @Tags Comments
// @Accept json
// @Produce json
// @Param comment body models.Comment true "Comment to be created"
// @Success 200 {object} models.CommentRes
// @Failure 400 {object} models.Error
// @Failure 500 {object} models.Error
// @Router /comment/post [post]
func (h *commentHandler) PostComment(c *gin.Context) {
	var tweet models.Comment
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

	UserID := cl["user_id"].(string)

	res := pb.Comment{
		Id:        uuid.NewString(),
		UserId:    UserID,
		TweetId:   tweet.TweetID,
		Content:   tweet.Content,
		LikeCount: tweet.LikeCount,
	}

	log.Printf("%+v", res)

	bady, err := json.Marshal(&res)
	if err != nil {
		h.logger.Error("Error occurred while marshaling json", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err = h.CommentMQ.PostComment(bady)
	if err != nil {
		h.logger.Error("Error occurred while posting json", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"massage": "success"})
}

// UpdateComment godoc
// @Summary Update an existing comment
// @Description Update the content of a comment
// @Security BearerAuth
// @Tags Comments
// @Accept json
// @Produce json
// @Param comment body models.UpdateAComment true "Updated comment details"
// @Success 200 {object} models.CommentRes
// @Failure 400 {object} models.Error
// @Failure 500 {object} models.Error
// @Router /comment/update [put]
func (h *commentHandler) UpdateComment(c *gin.Context) {
	var tweet models.UpdateAComment
	if err := c.ShouldBindJSON(&tweet); err != nil {
		h.logger.Error("Error occurred while binding json", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	res := pb.UpdateAComment{
		Id:      tweet.ID,
		Content: tweet.Content,
	}
	bady, err := json.Marshal(&res)
	if err != nil {
		h.logger.Error("Error occurred while marshaling json", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err = h.CommentMQ.UpdateComment(bady)
	if err != nil {
		h.logger.Error("Error occurred while updating json", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"massage": "success"})
}

// DeleteComment godoc
// @Summary Delete a comment
// @Description Remove a comment by ID
// @Security BearerAuth
// @Tags Comments
// @Accept json
// @Produce json
// @Param id path string true "ID of the comment to delete"
// @Success 200 {object} models.Message
// @Failure 400 {object} models.Error
// @Failure 500 {object} models.Error
// @Router /comment/delete/{id} [delete]
func (h *commentHandler) DeleteComment(c *gin.Context) {
	id := c.Param("id")

	res := pb.CommentId{
		Id: id,
	}

	req, err := h.CommentService.DeleteComment(c.Request.Context(), &res)
	if err != nil {
		h.logger.Error("Error occurred while deleting comment", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": req})
}

// GetComment godoc
// @Summary Get a specific comment
// @Description Retrieve a comment by ID
// @Security BearerAuth
// @Tags Comments
// @Accept json
// @Produce json
// @Param id path string true "ID of the comment to retrieve"
// @Success 200 {object} models.CommentRes
// @Failure 400 {object} models.Error
// @Failure 500 {object} models.Error
// @Router /comment/get/{id} [get]
func (h *commentHandler) GetComment(c *gin.Context) {
	id := c.Param("id")

	res := pb.CommentId{
		Id: id,
	}

	req, err := h.CommentService.GetComment(c.Request.Context(), &res)
	if err != nil {
		h.logger.Error("Error occurred while deleting comment", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": req})
}

// GetAllComments godoc
// @Summary Get all comments for a tweet
// @Description Retrieve all comments for a specific tweet
// @Security BearerAuth
// @Tags Comments
// @Accept json
// @Produce json
// @Param tweet_id path string true "ID of the tweet to retrieve comments for"
// @Success 200 {object} models.Comments
// @Failure 400 {object} models.Error
// @Failure 500 {object} models.Error
// @Router /comment/get_all/{tweet_id} [get]
func (h *commentHandler) GetAllComments(c *gin.Context) {
	userId := c.MustGet("user_id").(string)
	tweetId := c.Param("tweet_id")

	res := pb.CommentFilter{
		UserId:  userId,
		TweetId: tweetId,
	}
	req, err := h.CommentService.GetAllComments(c.Request.Context(), &res)
	if err != nil {
		h.logger.Error("Error occurred while getting comments", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": req})
}

// GetUserComments godoc
// @Summary Get all comments by a user
// @Description Retrieve all comments made by a specific user
// @Security BearerAuth
// @Tags Comments
// @Accept json
// @Produce json
// @Success 200 {object} models.Comments
// @Failure 400 {object} models.Error
// @Failure 500 {object} models.Error
// @Router /comment/get_user [get]
func (h *commentHandler) GetUserComments(c *gin.Context) {

	token := c.GetHeader("Authorization")
	cl, err := t.ExtractClaims(token)
	if err != nil {
		h.logger.Error("Error occurred while extracting claims", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res := pb.UserId{
		Id: cl["user_id"].(string),
	}
	req, err := h.CommentService.GetUserComments(c.Request.Context(), &res)
	if err != nil {
		h.logger.Error("Error occurred while getting comments", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": req})
}

// AddLikeToComment godoc
// @Summary Add a like to a comment
// @Description Increment the like count for a comment
// @Security BearerAuth
// @Tags Comments
// @Accept json
// @Produce json
// @Param comment_id path string true "ID of the comment to like"
// @Success 200 {object} models.Message
// @Failure 400 {object} models.Error
// @Failure 500 {object} models.Error
// @Router /comment/add_like/{comment_id} [get]
func (h *commentHandler) AddLikeToComment(c *gin.Context) {
	id := c.Param("comment_id")
	res := pb.CommentLikeReq{
		CommentId: id,
	}
	req, err := h.CommentService.AddLikeToComment(c.Request.Context(), &res)
	if err != nil {
		h.logger.Error("Error occurred while adding like to comment", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": req})
}

// DeleteLikeComment godoc
// @Summary Remove a like from a comment
// @Description Decrement the like count for a comment
// @Security BearerAuth
// @Tags Comments
// @Accept json
// @Produce json
// @Param comment_id path string true "ID of the comment to unlike"
// @Success 200 {object} models.Message
// @Failure 400 {object} models.Error
// @Failure 500 {object} models.Error
// @Router /comment/remove_like/{comment_id} [delete]
func (h *commentHandler) DeleteLikeComment(c *gin.Context) {
	id := c.Param("comment_id")
	res := pb.CommentLikeReq{
		CommentId: id,
	}
	req, err := h.CommentService.DeleteLikeComment(c.Request.Context(), &res)
	if err != nil {
		h.logger.Error("Error occurred while deleting like comment", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": req})
}

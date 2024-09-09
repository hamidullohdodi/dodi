package handler

import (
	pb "apigateway/genproto/user"
	"apigateway/pkg/models"
	"apigateway/service"
	"context"
	"database/sql"
	"errors"
	"github.com/gin-gonic/gin"
	"log"
	"log/slog"
	"net/http"
	"strconv"
)

type UserHandler interface {
	Create(c *gin.Context)
	GetProfile(c *gin.Context)
	UpdateProfile(c *gin.Context)
	ChangePassword(c *gin.Context)
	ChangeProfileImage(c *gin.Context)
	FetchUsers(c *gin.Context)
	ListOfFollowing(c *gin.Context)
	ListOfFollowers(c *gin.Context)
	ListOfFollowingByUsername(c *gin.Context)
	ListOfFollowersByUsername(c *gin.Context)
	DeleteUser(c *gin.Context)

	Follow(c *gin.Context)
	Unfollow(c *gin.Context)
	GetUserFollowers(c *gin.Context)
	GetUserFollows(c *gin.Context)
	MostPopularUser(c *gin.Context)
}

type userHandler struct {
	userService pb.UserServiceClient
	logger      *slog.Logger
}

func NewUserHandler(userService service.Service, logger *slog.Logger) UserHandler {
	userClient := userService.UserService()
	if userClient == nil {
		log.Fatalf("failed to create user client")
		return nil
	}
	return &userHandler{
		userService: userClient,
		logger:      logger,
	}
}

// Create godoc
// @Summary Create User
// @Description Create a new user
// @Security BearerAuth
// @Tags Admin
// @Accept json
// @Produce json
// @Param Create body models.CreateRequest true "Create user"
// @Success 201 {object} models.UserResponse
// @Failure 400 {object} models.Error
// @Failure 500 {object} models.Error
// @Router /admin/create [post]
func (h *userHandler) Create(c *gin.Context) {
	var user models.CreateRequest

	if err := c.ShouldBindJSON(&user); err != nil {
		h.logger.Error("Error occurred while binding json", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	req := pb.CreateRequest{
		Email:       user.Email,
		Password:    user.Password,
		Phone:       user.Phone,
		FirstName:   user.FirstName,
		LastName:    user.LastName,
		Username:    user.Username,
		Nationality: user.Nationality,
		Bio:         user.Bio,
	}

	res, err := h.userService.Create(context.Background(), &req)
	if err != nil {
		h.logger.Error("Error occurred while creating user", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, res)
}

// GetProfile godoc
// @Summary Get User Profile
// @Description Retrieve the profile of a user
// @Security BearerAuth
// @Tags User
// @Accept json
// @Produce json
// @Success 200 {object} models.GetProfileResponse
// @Failure 400 {object} models.Error
// @Failure 500 {object} models.Error
// @Router /user/get_profile [get]
func (h *userHandler) GetProfile(c *gin.Context) {
	req := pb.Id{
		UserId: c.MustGet("user_id").(string),
	}

	res, err := h.userService.GetProfile(context.Background(), &req)
	if err != nil {
		h.logger.Error("Error occurred while getting user", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

// GetProfileById godoc
// @Summary Get User Profile
// @Description Retrieve the profile of a user
// @Security BearerAuth
// @Tags Admin
// @Accept json
// @Produce json
// @Param user_id path string true "User ID"
// @Success 200 {object} models.GetProfileResponse
// @Failure 400 {object} models.Error
// @Failure 500 {object} models.Error
// @Router /admin/get_profile_by_id/{user_id} [get]
func (h *userHandler) GetProfileById(c *gin.Context) {
	req := pb.Id{
		UserId: c.Param("user_id"),
	}

	res, err := h.userService.GetProfile(context.Background(), &req)
	if err != nil {
		h.logger.Error("Error occurred while getting user", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

// UpdateProfile godoc
// @Summary Update User Profile
// @Description Update user profile details
// @Security BearerAuth
// @Tags User
// @Accept json
// @Produce json
// @Param UpdateProfile body models.UpdateProfileRequest true "Update user profile"
// @Success 200 {object} models.UserResponse
// @Failure 400 {object} models.Error
// @Failure 500 {object} models.Error
// @Router /user/update_profile [put]
func (h *userHandler) UpdateProfile(c *gin.Context) {
	var user models.UpdateProfileRequest
	if err := c.ShouldBindJSON(&user); err != nil {
		h.logger.Error("Error occurred while binding json", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	res := pb.UpdateProfileRequest{
		UserId:       c.MustGet("user_id").(string),
		FirstName:    user.FirstName,
		LastName:     user.LastName,
		PhoneNumber:  user.PhoneNumber,
		Username:     user.Username,
		Nationality:  user.Nationality,
		Bio:          user.Bio,
		ProfileImage: user.ProfileImage,
	}

	req, err := h.userService.UpdateProfile(context.Background(), &res)
	if err != nil {
		h.logger.Error("Error occurred while updating user", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, req)
}

// UpdateProfileById godoc
// @Summary Update User Profile
// @Description Update user profile details
// @Security BearerAuth
// @Tags Admin
// @Accept json
// @Produce json
// @Param user_id path string true "User ID"
// @Param UpdateProfileById body models.UpdateProfileRequest true "Update user profile"
// @Success 200 {object} models.UserResponse
// @Failure 400 {object} models.Error
// @Failure 500 {object} models.Error
// @Router /admin/update_profile_by_id/{user_id} [put]
func (h *userHandler) UpdateProfileById(c *gin.Context) {
	var user models.UpdateProfileRequest
	if err := c.ShouldBindJSON(&user); err != nil {
		h.logger.Error("Error occurred while binding json", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	res := pb.UpdateProfileRequest{
		UserId:       c.Param("user_id"),
		FirstName:    user.FirstName,
		LastName:     user.LastName,
		PhoneNumber:  user.PhoneNumber,
		Username:     user.Username,
		Nationality:  user.Nationality,
		Bio:          user.Bio,
		ProfileImage: user.ProfileImage,
	}

	req, err := h.userService.UpdateProfile(context.Background(), &res)
	if err != nil {
		h.logger.Error("Error occurred while updating user", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, req)
}

// ChangePassword godoc
// @Summary Change User Password
// @Description Change the password of a user
// @Security BearerAuth
// @Tags User
// @Accept json
// @Produce json
// @Param ChangePassword body models.ChangePasswordRequest true "Change password"
// @Success 200 {object} models.ChangePasswordResponse
// @Failure 400 {object} models.Error
// @Failure 500 {object} models.Error
// @Router /user/change_password [put]
func (h *userHandler) ChangePassword(c *gin.Context) {
	var user models.ChangePasswordRequest
	if err := c.ShouldBindJSON(&user); err != nil {
		h.logger.Error("Error occurred while binding json", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Validate request body
	if user.CurrentPassword == "" || user.NewPassword == "" {
		h.logger.Error("Invalid request body")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	res := pb.ChangePasswordRequest{
		UserId:          c.MustGet("user_id").(string),
		CurrentPassword: user.CurrentPassword,
		NewPassword:     user.NewPassword,
	}

	req, err := h.userService.ChangePassword(context.Background(), &res)
	if err != nil {
		h.logger.Error("Error occurred while changing user", err)
		if errors.Is(err, sql.ErrNoRows) {
			c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	if req == nil {
		h.logger.Error("Nil response from userService")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	c.JSON(http.StatusOK, req)
}

// ChangeProfileImage godoc
// @Summary Change User Profile Image
// @Description Update the profile image of a user
// @Security BearerAuth
// @Tags User
// @Accept json
// @Produce json
// @Param ChangeProfileImage body models.URL true "Change profile image"
// @Success 200 {object} models.Void
// @Failure 400 {object} models.Error
// @Failure 500 {object} models.Error
// @Router /user/change_profile_image [put]
func (h *userHandler) ChangeProfileImage(c *gin.Context) {
	var user models.URL
	if err := c.ShouldBindJSON(&user); err != nil {
		h.logger.Error("Error occurred while binding json", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	res := pb.URL{
		UserId: c.MustGet("user_id").(string),
		Url:    user.URL,
	}
	req, err := h.userService.ChangeProfileImage(context.Background(), &res)
	if err != nil {
		h.logger.Error("Error occurred while changing user", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, req)
}

// ChangeProfileImageById godoc
// @Summary Change User Profile Image
// @Description Update the profile image of a user
// @Security BearerAuth
// @Tags Admin
// @Accept json
// @Produce json
// @Param user_id path string true "User ID"
// @Param ChangeProfileImageById body models.URL true "Change profile image"
// @Success 200 {object} models.Void
// @Failure 400 {object} models.Error
// @Failure 500 {object} models.Error
// @Router /admin/change_profile_image_by_id/{user_id} [put]
func (h *userHandler) ChangeProfileImageById(c *gin.Context) {
	var user models.URL
	if err := c.ShouldBindJSON(&user); err != nil {
		h.logger.Error("Error occurred while binding json", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	res := pb.URL{
		UserId: c.Param("user_id"),
		Url:    user.URL,
	}
	req, err := h.userService.ChangeProfileImage(context.Background(), &res)
	if err != nil {
		h.logger.Error("Error occurred while changing user", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, req)
}

// FetchUsers godoc
// @Summary Fetch Users
// @Description Retrieve a list of users with filtering options
// @Security BearerAuth
// @Tags User
// @Accept json
// @Produce json
// @Param page query int false "Page number"
// @Param limit query int false "Number of users per page"
// @Param name query string false "Username"
// @Success 200 {object} user.UserResponses
// @Failure 400 {object} models.Error
// @Failure 500 {object} models.Error
// @Router /user/fetch_users [get]
func (h *userHandler) FetchUsers(c *gin.Context) {
	pageStr := c.Query("page")
	limitStr := c.Query("limit")
	name := c.Query("name")

	page, err := strconv.Atoi(pageStr)
	if err != nil {
		page = 1
	}

	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		limit = 10
	}

	res := pb.Filter{
		Page:      int32(page),
		Limit:     int32(limit),
		FirstName: name,
	}

	req, err := h.userService.FetchUsers(context.Background(), &res)
	if err != nil {
		h.logger.Error("Error occurred while fetching users", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, req)

}

// ListOfFollowing godoc
// @Summary List of Following Users
// @Description Retrieve the list of users that a specific user is following
// @Security BearerAuth
// @Tags User
// @Accept json
// @Produce json
// @Success 200 {object} user.Followings
// @Failure 400 {object} models.Error
// @Failure 500 {object} models.Error
// @Router /user/list_of_following [get]
func (h *userHandler) ListOfFollowing(c *gin.Context) {
	res := pb.Id{
		UserId: c.MustGet("user_id").(string),
	}
	req, err := h.userService.ListOfFollowing(context.Background(), &res)
	if err != nil {
		h.logger.Error("Error occurred while listing following", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, req)
}

// ListOfFollowingByUsername godoc
// @Summary List of Following Users
// @Description Retrieve the list of users that a specific user is following
// @Security BearerAuth
// @Tags User
// @Accept json
// @Produce json
// @Param username path string true "Username"
// @Success 200 {object} user.Followings
// @Failure 400 {object} models.Error
// @Failure 500 {object} models.Error
// @Router /user/list_of_following_by_username/{username} [get]
func (h *userHandler) ListOfFollowingByUsername(c *gin.Context) {
	res := pb.Id{
		UserId: c.Param("username"),
	}
	req, err := h.userService.ListOfFollowingByUsername(context.Background(), &res)
	if err != nil {
		h.logger.Error("Error occurred while listing following", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, req)
}

// ListOfFollowersByUsername godoc
// @Summary List of Followers
// @Description Retrieve the list of followers for a user
// @Security BearerAuth
// @Tags User
// @Accept json
// @Produce json
// @Param username path string true "Username"
// @Success 200 {object} user.Followers
// @Failure 400 {object} models.Error
// @Failure 500 {object} models.Error
// @Router /user/list_of_followers_by_username/{username} [get]
func (h *userHandler) ListOfFollowersByUsername(c *gin.Context) {
	res := pb.Id{
		UserId: c.Param("username"),
	}
	req, err := h.userService.ListOfFollowersByUsername(context.Background(), &res)
	if err != nil {
		h.logger.Error("Error occurred while listing followers", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, req)
}

// ListOfFollowers godoc
// @Summary List of Followers
// @Description Retrieve the list of followers for a user
// @Security BearerAuth
// @Tags User
// @Accept json
// @Produce json
// @Success 200 {object} user.Followers
// @Failure 400 {object} models.Error
// @Failure 500 {object} models.Error
// @Router /user/list_of_followers [get]
func (h *userHandler) ListOfFollowers(c *gin.Context) {
	res := pb.Id{
		UserId: c.MustGet("user_id").(string),
	}
	req, err := h.userService.ListOfFollowers(context.Background(), &res)
	if err != nil {
		h.logger.Error("Error occurred while listing followers", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, req)
}

// DeleteUser godoc
// @Summary Delete User
// @Description Delete a user account
// @Security BearerAuth
// @Tags Admin
// @Accept json
// @Produce json
// @Param user_id path string true "User ID"
// @Success 200 {object} models.Message
// @Failure 400 {object} models.Error
// @Failure 500 {object} models.Error
// @Router /admin/delete/{user_id} [delete]
func (h *userHandler) DeleteUser(c *gin.Context) {
	res := pb.Id{
		UserId: c.Param("user_id"),
	}
	req, err := h.userService.DeleteUser(context.Background(), &res)
	if err != nil {
		h.logger.Error("Error occurred while deleting user", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, models.Message{Message: "Successfully deleted user" + req.String()})
}

// Follow godoc
// @Summary Follow User
// @Description Follow another user
// @Security BearerAuth
// @Tags User
// @Accept json
// @Produce json
// @Param Follow body models.FollowReq true "post user"
// @Success 200 {object} models.FollowRes
// @Failure 400 {object} models.Error
// @Failure 500 {object} models.Error
// @Router /user/follow [post]
func (h *userHandler) Follow(c *gin.Context) {
	var user models.FollowReq
	if err := c.ShouldBindJSON(&user); err != nil {
		h.logger.Error("Error occurred while binding json", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res := pb.FollowReq{
		FollowingId: user.FollowingID,
		FollowerId:  c.MustGet("user_id").(string),
	}

	req, err := h.userService.Follow(context.Background(), &res)
	if err != nil {
		h.logger.Error("Error occurred while following", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, req)
}

// Unfollow godoc
// @Summary Unfollow User
// @Description Unfollow a user
// @Security BearerAuth
// @Tags User
// @Accept json
// @Produce json
// @Param Unfollow body models.FollowReq true "put user"
// @Success 200 {object} models.DFollowRes
// @Failure 400 {object} models.Error
// @Failure 500 {object} models.Error
// @Router /user/unfollow [delete]
func (h *userHandler) Unfollow(c *gin.Context) {
	var user models.FollowReq
	if err := c.ShouldBindJSON(&user); err != nil {
		h.logger.Error("Error occurred while binding json", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	res := pb.FollowReq{
		FollowingId: user.FollowingID,
		FollowerId:  c.MustGet("user_id").(string),
	}

	req, err := h.userService.Unfollow(context.Background(), &res)
	if err != nil {
		h.logger.Error("Error occurred while following", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, req)
}

// GetUserFollowers godoc
// @Summary Get User Followers
// @Description Retrieve a list of followers for a specific user
// @Security BearerAuth
// @Tags User
// @Accept json
// @Produce json
// @Success 200 {object} models.Count
// @Failure 400 {object} models.Error
// @Failure 500 {object} models.Error
// @Router /user/get_user_followers [get]
func (h *userHandler) GetUserFollowers(c *gin.Context) {
	res := pb.Id{
		UserId: c.MustGet("user_id").(string),
	}
	req, err := h.userService.GetUserFollowers(context.Background(), &res)
	if err != nil {
		h.logger.Error("Error occurred while following", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, req)
}

// GetUserFollows godoc
// @Summary Get User Follows
// @Description Retrieve a list of users that a specific user is following
// @Security BearerAuth
// @Tags User
// @Accept json
// @Produce json
// @Success 200 {object} models.Count
// @Failure 400 {object} models.Error
// @Failure 500 {object} models.Error
// @Router /user/get_user_follows [get]
func (h *userHandler) GetUserFollows(c *gin.Context) {
	res := pb.Id{
		UserId: c.MustGet("user_id").(string),
	}
	req, err := h.userService.GetUserFollows(context.Background(), &res)
	if err != nil {
		h.logger.Error("Error occurred while following", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, req)
}

// MostPopularUser godoc
// @Summary Most Popular User
// @Description Retrieve the most popular user based on criteria
// @Security BearerAuth
// @Tags User
// @Accept json
// @Produce json
// @Success 200 {object} models.UserResponse
// @Failure 400 {object} models.Error
// @Failure 500 {object} models.Error
// @Router /user/most_popular [get]
func (h *userHandler) MostPopularUser(c *gin.Context) {
	res := pb.Void{}
	req, err := h.userService.MostPopularUser(context.Background(), &res)
	if err != nil {
		h.logger.Error("Error occurred while most popular user", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, req)
}

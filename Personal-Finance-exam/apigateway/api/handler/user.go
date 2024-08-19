package handler

import (
	pb "api/genproto/user"
	_ "api/models"
	"api/service"
	"context"
	"github.com/gin-gonic/gin"
	"log/slog"
	"net/http"
)

type UserHandler interface {
	GetProfile(c *gin.Context)
	UpdateProfile(c *gin.Context)
	DeleteUser(c *gin.Context)
}

type UserHandlerIml struct {
	userClient pb.UserServiceClient
	logger     *slog.Logger
}

func NewUserHandler(serviceManager service.ServiceManager, logger *slog.Logger) UserHandler {
	userClient := serviceManager.UserService()
	if userClient == nil {
		logger.Error("userClient is nil")
		return nil
	}
	return &UserHandlerIml{
		userClient: userClient,
		logger:     logger,
	}
}

// GetProfile godoc
// @Summary Get user profile by ID
// @Description Get user profile details by ID
// @Security BearerAuth
// @Tags user
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Success 200 {object} models.UserResp
// @Failure 500 {object} string "Error while retrieving user profile"
// @Router /user/profile/{id} [get]
func (u *UserHandlerIml) GetProfile(c *gin.Context) {
	id := c.Param("id")
	rep := &pb.UserId{
		Id: id,
	}
	resp, err := u.userClient.GetProfile(context.Background(), rep)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, resp)
}

// UpdateProfile godoc
// @Summary Update user profile
// @Description Update user profile details
// @Security BearerAuth
// @Tags user
// @Accept json
// @Produce json
// @Param profile body models.UpdateProfileReq true "Update Profile Request"
// @Success 200 {object} models.UserResp
// @Failure 400 {object} string "Invalid data"
// @Failure 500 {object} string "Error while updating user profile"
// @Router /user/profile [put]
func (u *UserHandlerIml) UpdateProfile(c *gin.Context) {
	req := &pb.EditProfileReq{}
	if err := c.ShouldBind(req); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	resp, err := u.userClient.UpdateProfile(context.Background(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, resp)
}

// DeleteUser godoc
// @Summary Delete user by ID
// @Description Delete a user by ID
// @Security BearerAuth
// @Tags user
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Success 200 {object} models.Void
// @Failure 500 {object} string "Error while deleting user"
// @Router /user/profile/{id} [delete]
func (u *UserHandlerIml) DeleteUser(c *gin.Context) {
	id := c.Param("id")
	rep := &pb.UserId{
		Id: id,
	}
	resp, err := u.userClient.DeleteUser(context.Background(), rep)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, resp)
}

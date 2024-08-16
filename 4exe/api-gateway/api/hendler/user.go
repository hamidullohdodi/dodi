package hendler

import (
	"api-geteway/genproto/user"
	"api-geteway/service"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserHandler interface {


	GetProfile(c *gin.Context)
	UpdateProfile(c *gin.Context)
	DeleteProfile(c *gin.Context)
	GetAllProfile(c *gin.Context)
}

type userHendler struct {
	userService user.AuthClient
	logger      *slog.Logger
}

func NewUserHendler(serviceManager service.ServiceManager, logger *slog.Logger) UserHandler {
	return &userHendler{
		userService: serviceManager.UserService(),
		logger:      logger,
	}
}




// @Summary GetProfile
// @Description GetProfile
// @Tags auth
// @ID get-profile
// @Accept  json
// @Produce  json
// @Param id query string true "User ID"
// @Success 200 {object} user.GetProfileResponse
// @Failure 400 {object} string "Invalid Argument"
// @Failure 404 {object} string "Not Found"
// @Failure 500 {object} string "Internal Server Error"
// @Router /auth/{id} [get]
func (h *userHendler) GetProfile(c *gin.Context) {
	Id := c.Query("id")
	if Id == "" {
		h.logger.Error("id is empty")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Id is required"})
		return
	}

	res, err := h.userService.GetByIdProfile(c.Request.Context(), &user.GetProfileRequest{Id: Id})
	if err != nil {
		h.logger.Error(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, res)
}

// @Summary UpdateProfile
// @Description UpdateProfile
// @Tags auth
// @ID update-profile
// @Accept  json
// @Produce  json
// @Param user body user.UpdateProfileRequest true "User"
// @Success 200 {object} user.UpdateProfileResponse
// @Failure 400 {object} string "Invalid Argument"
// @Failure 404 {object} string "Not Found"
// @Failure 500 {object} string "Internal Server Error"
// @Router /auth/update-profile [put]
func (h *userHendler) UpdateProfile(c *gin.Context) {
	req := user.UpdateProfileRequest{}
	err := c.BindJSON(&req)
	if err != nil {
		h.logger.Error(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	res, err := h.userService.UpdateUserProfile(c.Request.Context(), &req)
	if err != nil {
		h.logger.Error(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, res)
}

// @Summary DeleteProfile
// @Description DeleteProfile
// @Tags auth
// @ID delete-profile
// @Accept  json
// @Produce  json
// @Param user body user.DeleteProfileRequest true "User"
// @Success 200 {object} user.DeleteProfileResponse
// @Failure 400 {object} string "Invalid Argument"
// @Failure 404 {object} string "Not Found"
// @Failure 500 {object} string "Internal Server Error"
// @Router /auth/delete-profile [delete]
func (h *userHendler) DeleteProfile(c *gin.Context) {
	req := user.DeleteProfileRequest{}
	err := c.BindJSON(&req)
	if err != nil {
		h.logger.Error(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	res, err := h.userService.DeleteUserProfile(c.Request.Context(), &req)
	if err != nil {
		h.logger.Error(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, res)
}

// @Summary GetAllProfile
// @Description GetAllProfile
// @Tags auth
// @ID get-all-profile
// @Accept  json
// @Produce  json
// @Param offset query int false "Offset"
// @Param limit query int false "Limit"
// @Success 200 {object} user.GetProfilesResponse
// @Failure 400 {object} string "Invalid Argument"
// @Failure 404 {object} string "Not Found"
// @Failure 500 {object} string "Internal Server Error"
// @Router /auth/profiles [get]
func (h *userHendler) GetAllProfile(c *gin.Context) {
	req := user.GetProfilesRequest{}
	LimitSrt := c.Query("limit")
	if LimitSrt != "" {
		limit, err := strconv.Atoi(LimitSrt)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		req.Limit = int32(limit)
	}
	OffsetSrt := c.Query("offset")
	if OffsetSrt != "" {
		offset, err := strconv.Atoi(OffsetSrt)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		req.Offset = int32(offset)
	}

	res, err := h.userService.GetAllProfile(c.Request.Context(), &req)
	if err != nil {
		h.logger.Error(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, res)
}

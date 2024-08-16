package handler

import (
	"auth_service/config"
	"auth_service/genproto/user"
	"auth_service/model"
	"auth_service/service"
	"context"

	// "auth_service/storage/redis"
	"auth_service/api/token"
	"fmt"
	"log"
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

type AuthenticaionHandler interface {
	Register(c *gin.Context)
	Login(c *gin.Context)
	LogOut(c *gin.Context)
}

type AuthenticaionHandlerImpl struct {
	userService service.AuthServiceI
	logger      *slog.Logger
	Config      *config.Config
	ContextTimeout time.Duration
}

func handleError(c *gin.Context, h *AuthenticaionHandlerImpl, err error, msg string, code int) {
	er := errors.Wrap(err, msg).Error()
	c.AbortWithStatusJSON(code, gin.H{"error": er})

	h.logger.Error(er)
}

func NewAuthenticaionHandlerImpl(userService service.AuthServiceI, logger *slog.Logger) *AuthenticaionHandlerImpl {
	return &AuthenticaionHandlerImpl{
		userService: userService,
		logger:      logger,
	}
}

// Register implements AuthenticaionHandler
// @Summary Register
// @Description Register
// @ID register
// @Accept  json
// @Produce  json
// @Param user body user.RegisterRequest true "User"
// @Success 200 {object} user.RegisterResponse
// @Failure 400 {object} string "Invalid Argument"
// @Failure 404 {object} string "Not Found"
// @Failure 500 {object} string "Internal Server Error"
// @Router /auth/register [post]
func (a *AuthenticaionHandlerImpl) Register(c *gin.Context) {

	req := user.RegisterRequest{}
	err := c.BindJSON(&req)
	if err != nil {
		a.logger.Error(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	res, err := a.userService.Register(c, &req)
	if err != nil {
		a.logger.Error(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	log.Println(res)
	c.JSON(http.StatusOK, res)
}

// @Summary Login
// @Description Login
// @ID login
// @Accept  json
// @Produce  json
// @Param user body model.LoginRequest true "User"
// @Success 200 {object} model.Tokens
// @Failure 400 {object} string "Invalid Argument"
// @Failure 404 {object} string "Not Found"
// @Failure 500 {object} string "Internal Server Error"
// @Router /auth/login [post]
func (a *AuthenticaionHandlerImpl) Login(c *gin.Context) {

	req := model.LoginRequest{}
	err := c.ShouldBindJSON(&req)
	if err != nil {
		a.logger.Error(err.Error())
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	fmt.Println(req)

	res, err := a.userService.Login(c.Request.Context(), &req)
	if err != nil {
		a.logger.Error(err.Error())
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	accessToken, err := token.GenerateAccessToken(a.Config, res.Id, res.Role)
	if err != nil {
		handleError(c, a, err, "error generating access token", http.StatusInternalServerError)
		return
	}

	refreshToken, err := token.GenerateRefreshToken(a.Config, res.Id)
	if err != nil {
		handleError(c, a, err, "error generating refresh token", http.StatusInternalServerError)
		return
	}

	a.logger.Info("user logged in successfully")
	c.JSON(http.StatusOK, model.Tokens{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	})
}

// @Summary LogOut a user
// @Description LogOut a user
// @Tags auth
// @Accept  json
// @Produce  json
// @Param user body model.LogoutRequest true "User"
// @Success 200 {object} model.LogoutResponse
// @Failure 400 {object} string "Invalid Argument"
// @Failure 404 {object} string "Not Found"
// @Failure 500 {object} string "Internal Server Error"
// @Router /auth/logout [post]
func (a *AuthenticaionHandlerImpl) LogOut(c *gin.Context) {
	req := model.LogoutRequest{}

	res, err := a.userService.LogOut(c.Request.Context(), &req)
	if err != nil {
		a.logger.Error(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, res)

}

// RefreshToken godoc
// @Summary Refreshes token
// @Description Refreshes refresh token
// @Tags auth
// @Param data body models.RefreshToken true "Refresh token"
// @Success 200 {object} models.Tokens
// @Failure 400 {object} string "Invalid data"
// @Failure 500 {object} string "Server error while processing request"
// @Router /refresh [post]
func (h *AuthenticaionHandlerImpl) RefreshToken(c *gin.Context) {
	h.logger.Info("RefreshToken handler is invoked")
	req := model.RefreshToken{}
	if err := c.ShouldBindJSON(&req); err != nil {
		handleError(c, h, err, "invalid data", http.StatusBadRequest)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), h.ContextTimeout)
	defer cancel()

	valid, err := token.ValidateRefreshToken(h.Config, req.RefreshToken)
	if !valid || err != nil {
		handleError(c, h, err, "error validating refresh token", http.StatusInternalServerError)
		return
	}

	id, err := token.ExtractRefreshUserID(h.Config, req.RefreshToken)
	if err != nil {
		handleError(c, h, err, "error extracting user id from refresh token", http.StatusInternalServerError)
		return
	}
	user, err := h.userService.(ctx, id)
	if err != nil {
		handleError(c, h, err, "error getting user", http.StatusInternalServerError)
		return
	}
	accessToken, err := token.GenerateAccessToken(h.Config, id, user.Role)
	if err != nil {
		handleError(c, h, err, "error generating access token", http.StatusInternalServerError)
		return
	}

	h.logger.Info("RefreshToken handler is completed successfully")
	c.JSON(http.StatusOK, model.Tokens{
		AccessToken:  accessToken,
		RefreshToken: req.RefreshToken,
	})
}

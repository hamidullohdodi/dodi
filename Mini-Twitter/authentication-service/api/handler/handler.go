package handler

import (
	"auth-service/api/email"
	"auth-service/pkg/models"
	"auth-service/service"
	"auth-service/storage/redis"
	"context"
	"github.com/badoux/checkmail"
	"github.com/gin-gonic/gin"
	"log/slog"
	"time"

	"net/http"

	_ "auth-service/api/docs"
)

type AuthHandler interface {
	Register(c *gin.Context)
	LoginEmail(c *gin.Context)
	LoginUsername(c *gin.Context)
	AcceptCodeToRegister(c *gin.Context)
}

type authHandler struct {
	srv   service.AuthService
	log   *slog.Logger
	redis *redis.RedisStorage
}

func NewAuthHandler(log *slog.Logger, sr service.AuthService, redis *redis.RedisStorage) AuthHandler {
	return &authHandler{log: log, srv: sr, redis: redis}
}

// Register godoc
// @Summary Register Users
// @Description create users
// @Tags Auth
// @Accept json
// @Produce json
// @Param Register body models.RegisterRequest true "register user"
// @Success 200 {object} models.RegisterResponse
// @Failure 400 {object} models.Error
// @Failure 404 {object} models.Error
// @Failure 500 {object} models.Error
// @Router /register [post]
func (h *authHandler) Register(c *gin.Context) {
	var auth models.RegisterRequest

	if err := c.ShouldBindJSON(&auth); err != nil {
		h.log.Error("Error occurred while binding json", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := checkmail.ValidateFormat(auth.Email)
	if err != nil {
		h.log.Error("Invalid email provided", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid email provided"})
		return
	}
	code, err := email.Email(auth.Email)
	if err != nil {
		h.log.Error("Invalid email provided", "error", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid email provided: " + err.Error()})
		return
	}
	req1 := models.RegisterRequest1{
		FirstName:   auth.FirstName,
		LastName:    auth.LastName,
		Email:       auth.Email,
		Phone:       auth.Phone,
		Username:    auth.Username,
		Nationality: auth.Nationality,
		Bio:         auth.Bio,
		Password:    auth.Password,
	}
	req1.Code = code

	err = h.redis.SetRegister(c, req1)
	if err != nil {
		h.log.Error("Failed to register user", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to register user"})
		return
	}
	h.log.Info("Successfully saved to redis")

	c.JSON(http.StatusOK, gin.H{"info": "code sent to this email " + req1.Email})
}

// AcceptCodeToRegister godoc
// @Summary Accept code to register
// @Description it accepts code to register
// @Tags Auth
// @Param token body models.AcceptCode true "enough"
// @Success 200 {object} models.RegisterResponse
// @Failure 400 {object} models.Error
// @Failure 500 {object} models.Error
// @Router /accept-code [post]
func (h *authHandler) AcceptCodeToRegister(c *gin.Context) {
	var req models.AcceptCode
	if err := c.ShouldBindJSON(&req); err != nil {
		h.log.Error("Invalid data provided", "error", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx, cancel := context.WithTimeout(c, 10*time.Second)
	defer cancel()

	register, err := h.redis.GetRegister(ctx, req.Email)
	if err != nil {
		h.log.Error("Failed to get register from redis", "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get register from redis; " + err.Error()})
		return
	}

	if register.Code != req.Code {
		h.log.Error("Invalid code", "code", req.Code)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid code"})
		return
	}

	response, err := h.srv.Register(models.RegisterRequest{
		FirstName:   register.FirstName,
		LastName:    register.LastName,
		Email:       register.Email,
		Phone:       register.Phone,
		Username:    register.Username,
		Nationality: register.Nationality,
		Bio:         register.Bio,
		Password:    register.Password,
	})
	if err != nil {
		h.log.Error("Failed to register student", "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to register student; " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, response)
}

// @Summary LoginEmail Users
// @Description sign in user
// @Tags Auth
// @Accept json
// @Produce json
// @Param LoginEmail body models.LoginEmailRequest true "register user"
// @Success 200 {object} models.Tokens
// @Failure 400 {object} models.Error
// @Failure 404 {object} models.Error
// @Failure 500 {object} models.Error
// @Router /login/email [post]
func (h *authHandler) LoginEmail(c *gin.Context) {
	var auth models.LoginEmailRequest

	if err := c.ShouldBindJSON(&auth); err != nil {
		h.log.Error("Error occurred while binding json", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := h.srv.LoginEmail(auth)
	if err != nil {
		h.log.Error("Error occurred while login", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.SetCookie("access_token", res.AccessToken, 3600, "", "", false, true)
	c.SetCookie("refresh_token", res.RefreshToken, 3600, "", "", false, true)

	c.JSON(http.StatusOK, res)
}

// @Summary LoginUsername Users
// @Description sign in user
// @Tags Auth
// @Accept json
// @Produce json
// @Param LoginUsername body models.LoginUsernameRequest true "register user"
// @Success 200 {object} models.Tokens
// @Failure 400 {object} models.Error
// @Failure 404 {object} models.Error
// @Failure 500 {object} models.Error
// @Router /login/username [post]
func (h *authHandler) LoginUsername(c *gin.Context) {
	var auth models.LoginUsernameRequest
	if err := c.ShouldBindJSON(&auth); err != nil {
		h.log.Error("Error occurred while binding json", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := h.srv.LoginUsername(auth)
	if err != nil {
		h.log.Error("Error occurred while login", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.SetCookie("access_token", res.AccessToken, 3600, "", "", false, true)
	c.SetCookie("refresh_token", res.RefreshToken, 3600, "", "", false, true)

	c.JSON(http.StatusOK, res)
}

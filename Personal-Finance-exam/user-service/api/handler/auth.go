package handler

import (
	t "user/api/token"
	"user/genproto/auth"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"regexp"
	"user/models"
)

func isValidEmail(email string) bool {
	const emailRegexPattern = `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	re := regexp.MustCompile(emailRegexPattern)
	return re.MatchString(email)
}

// RegisterUser handles user registration
// @Summary Register a new user
// @Description Register a new user
// @Tags auth
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param user body auth.RegisterUserReq true "Register User Request"
// @Success 200 {object} string "User registered successfully"
// @Failure 400 {string} string "invalid request"
// @Failure 500 {string} string "internal server error"
// @Router /register [post]
func (h *Handlers) RegisterUser(c *gin.Context) {
	var req auth.RegisterUserReq
	if err := c.BindJSON(&req); err != nil {
		log.Printf("failed to bind JSON: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	password, err := t.HashPassword(req.Password) // h ni ishlatish kerak bo'lishi mumkin
	if err != nil {
		log.Printf("failed to hash password: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	req.Password = password

	if isValidEmail(req.Email) {
		fmt.Println("Valid email")
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid email format"})
		return
	}

	resp, err := h.Auth.Register(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"Register": resp})
}

// LoginUser handles user login
// @Summary Login a user
// @Description Login a user
// @Tags auth
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param user body auth.LoginUserReq true "Login Request"
// @Success 200 {string} auth.User
// @Failure 400 {string} string "invalid request"
// @Failure 500 {string} string "internal server error"
// @Router /login [post]
func (h *Handlers) LoginUser(c *gin.Context) {
	var req auth.LoginUserReq
	if err := c.BindJSON(&req); err != nil {
		log.Printf("failed to bind JSON: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	res, err := h.Auth.Login(context.Background(), &req)
	if err != nil {
		log.Printf("failed to login user: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	token, _ := t.GenerateJWTToken(res)
	c.JSON(http.StatusOK, token)
}

// LogoutUser handles user LogoutUser
// @Summary LogoutUser a user
// @Description LogoutUser a user
// @Tags auth
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param user body models.LoginRequest true "LogoutUser Request"
// @Success 200 {string} string "logout success"
// @Failure 400 {string} string "invalid request"
// @Failure 500 {string} string "internal server error"
// @Router /LogoutUser [put]
func (h *Handlers) LogoutUser(c *gin.Context) {
	var req models.LoginRequest
	if err := c.BindJSON(&req); err != nil {
		log.Printf("failed to bind JSON: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, "logout success")
}

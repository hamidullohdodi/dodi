package hendler

import (
	pb "api_getway/genproto"
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
)

// RegisterUser registers a new user.
// @Summary Register User
// @Description Register a new user
// @Tags register
// @Accept json
// @Produce json
// @Param user body pb.RegisterUserReq true "Register User Request"
// @Success 200 {object} pb.Register
// @Failure 400 {object} error "Error"
// @Failure 500 {object} error "Error"
// @Router /registers/register [post]
func (h *Handler) RegisterUser(c *gin.Context) {
	var req pb.RegisterUserReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	r, err := h.Auth.RegisterUser(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"user": r})
}

// Login logs in a user.
// @Summary Login User
// @Description Log in a user
// @Tags register
// @Accept json
// @Produce json
// @Param user body pb.LoginReq true "Login Request"
// @Success 200 {object} pb.Token
// @Failure 400 {object} error "Error"
// @Failure 500 {object} error "Error"
// @Router /registers/login [post]
func (h *Handler) Login(c *gin.Context) {
	var req pb.LoginReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	r, err := h.Auth.LoginUser(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"user": r})
}

// GetUser retrieves user information.
// @Summary Get User
// @Description Get user information
// @Tags user
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "User ID"
// @Success 200 {object} pb.Register
// @Failure 400 {object} error "Error"
// @Failure 500 {object} error "Error"
// @Router /users/get/{id} [get]
func (h *Handler) GetUser(c *gin.Context) {
	id := c.Param("id")

	req := pb.ById{Id: id}

	r, err := h.Auth.GetUser(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"user": gin.H{
			"id":        r.Id,
			"username":  r.Username,
			"email":     r.Email,
			"password":  r.Password,
			"full_name": r.FullName,
			"user_type": r.UserType,
			"bio":       r.Bio,
		},
	})
}

// UpdateUser updates user information.
// @Summary Update User
// @Description Update user information
// @Tags user
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param user body pb.UpdateUserReq true "Update User Request"
// @Success 200 {object} pb.Register
// @Failure 400 {object} error "Error"
// @Failure 500 {object} error "Error"
// @Router /users/update [put]
func (h *Handler) UpdateUser(c *gin.Context) {
	var req pb.UpdateUserReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	r, err := h.Auth.UpdateUser(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"user": gin.H{
			"id":        r.Id,
			"username":  r.Username,
			"email":     r.Email,
			"password":  r.Password,
			"full_name": r.FullName,
			"user_type": r.UserType,
			"bio":       r.Bio,
		},
	})
}

// GetAllUsers retrieves all users.
// @Summary Get All Users
// @Description Get all users
// @Tags user
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param page query int false "Page number"
// @Param limit query int false "Limit per page"
// @Success 200 {object} pb.GetAllUsersRes
// @Failure 400 {object} error "Error"
// @Failure 500 {object} error "Error"
// @Router /users/all/{id} [get]
func (h *Handler) GetAllUsers(c *gin.Context) {
	var req pb.PageLimit

	r, err := h.Auth.GetAllUsers(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"users": r})
}

// DeleteUser deletes a user.
// @Summary Delete User
// @Description Delete a user
// @Tags user
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "User ID"
// @Success 200 {object} pb.Void
// @Failure 400 {object} error "Error"
// @Failure 500 {object} error "Error"
// @Router /users/delete/{id} [delete]
func (h *Handler) DeleteUser(c *gin.Context) {
	id := c.Param("id")
	rep := &pb.ById{Id: id}
	_, err := h.Auth.DeleteUser(context.Background(), rep)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, "Delete")
}

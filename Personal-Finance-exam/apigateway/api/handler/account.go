package handler

import (
	"api/api/token"
	pb "api/genproto/account"
	_ "api/models"
	"api/service"
	"context"
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AccountHandler interface { //methodlar
	CreateAccount(c *gin.Context)
	GetAccount(c *gin.Context)
	UpdateAccount(c *gin.Context)
	DeleteAccount(c *gin.Context)
	ListAccounts(c *gin.Context)
}

type AccountHandlerIml struct {
	AccountClient pb.AccountServiceClient
	logger        *slog.Logger
}

func NewAccountHandler(serviceManger service.ServiceManager, logger *slog.Logger) AccountHandler {
	return &AccountHandlerIml{
		AccountClient: serviceManger.AccountService(),
		logger:        logger,
	}
}

// CreateAccount godoc
// @Summary Create a new Account
// @Description Create a new device with the provided details
// @Security BearerAuth
// @Tags account
// @Accept json
// @Produce json
// @Param account body models.CreateAccountReq true "account Creation Data"
// @Success 201 {object} models.AccountResp
// @Failure 400 {object} string "Invalid data"
// @Failure 500 {object} string "Error while creating account "
// @Router /account/create [post]
func (h *AccountHandlerIml) CreateAccount(c *gin.Context) {
	req := pb.CreateAccountReq{}
	if err := c.ShouldBind(&req); err != nil {
		h.logger.Error("failed to bind request", "error", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	tokens := c.GetHeader("Authorization")
	cl, err := token.ExtractClaims(tokens)
	if err != nil {
		h.logger.Error("failed to extract claims", "error", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	req.UserId = cl["user_id"].(string)

	resp, err := h.AccountClient.CreateAccount(context.Background(), &req)
	if err != nil {
		h.logger.Error("failed to create account", "error", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, resp)
}

// GetAccount godoc
// @Summary Get  account by ID
// @Description Get account by ID
// @Security BearerAuth
// @Tags account
// @Accept json
// @Produce json
// @Param id path string true "Account ID"
// @Success 200 {object} models.AccountResp
// @Failure 400 {object} string "Invalid data"
// @Failure 500 {object} string "Error while retrieving account"
// @Router /account/get/{id} [get]
func (h *AccountHandlerIml) GetAccount(c *gin.Context) {
	id := c.Param("id")
	req := &pb.GetAccountReq{
		UserId: id,
	}
	resp, err := h.AccountClient.GetAccount(context.Background(), req)
	if err != nil {
		h.logger.Error("failed to get account", "error", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}

// UpdateAccount godoc
// @Summary Update account by ID
// @Description Update account details by ID
// @Security BearerAuth
// @Tags account
// @Accept json
// @Produce json
// @Param account body models.UpdateAccountReq true "Update Account Request"
// @Success 200 {object} models.AccountResp
// @Failure 400 {object} string "Invalid data"
// @Failure 500 {object} string "Error while updating account"
// @Router /account/update [put]
func (h *AccountHandlerIml) UpdateAccount(c *gin.Context) {
	req := pb.UpdateAccountReq{}
	if err := c.ShouldBind(&req); err != nil {
		h.logger.Error("failed to bind request", "error", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	resp, err := h.AccountClient.UpdateAccount(context.Background(), &req)
	if err != nil {
		h.logger.Error("failed to update account", "error", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}

// DeleteAccount godoc
// @Summary Delete account by ID
// @Description Delete an account by ID
// @Security BearerAuth
// @Tags account
// @Accept json
// @Produce json
// @Param id path string true "Account ID"
// @Success 200 {object} string "Delete"
// @Failure 400 {object} string "Invalid data"
// @Failure 500 {object} string "Error while deleting account"
// @Router /account/delete/{id} [delete]
func (h *AccountHandlerIml) DeleteAccount(c *gin.Context) {
	id := c.Param("id")
	req := &pb.DeleteAccountReq{
		UserId: id,
	}
	resp, err := h.AccountClient.DeleteAccount(context.Background(), req)
	if err != nil {
		h.logger.Error("failed to delete account", "error", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}

// ListAccounts godoc
// @Summary List all accounts
// @Description Get a list of all accounts
// @Security BearerAuth
// @Tags account
// @Accept json
// @Produce json
// @Param limit query string true "Account ID"
// @Param offset query string true "Account ID"
// @Success 200 {object} models.ListAccountsResp
// @Failure 400 {object} string "Invalid data"
// @Failure 500 {object} string "Error while listing accounts"
// @Router /account/list [get]
func (h *AccountHandlerIml) ListAccounts(c *gin.Context) {
	req := &pb.ListAccountsReq{}
	limit := c.Query("limit")
	offset := c.Query("offset")

	req.Limit = limit
	req.Paid = offset

	resp, err := h.AccountClient.ListAccounts(context.Background(), req)
	if err != nil {
		h.logger.Error("failed to list accounts", "error", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}

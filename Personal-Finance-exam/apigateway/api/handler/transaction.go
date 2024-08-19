package handler

import (
	"api/api/token"
	pb "api/genproto/transaction"
	_ "api/models"
	"api/service"
	"context"
	"encoding/json"
	"github.com/gin-gonic/gin"
	amqp "github.com/rabbitmq/amqp091-go"
	"log/slog"
	"net/http"
)

type TransactionHandler interface {
	CreateTransaction(c *gin.Context)
	GetTransaction(c *gin.Context)
	UpdateTransaction(c *gin.Context)
	DeleteTransaction(c *gin.Context)
	ListTransactions(c *gin.Context)
}

type TransactionHandlerIml struct {
	Produser          *service.MsgBroker
	TransactionClient pb.TransactionServiceClient
	logger            *slog.Logger
}

func NewTransactionHandler(serviceManger service.ServiceManager, logger *slog.Logger, conn *amqp.Channel) TransactionHandler {
	return &TransactionHandlerIml{
		Produser:          service.NewMsgBroker(conn, logger),
		TransactionClient: serviceManger.TransactionService(),
		logger:            logger,
	}
}

// CreateTransaction godoc
// @Summary Create a new transaction
// @Description Create a new transaction
// @Security BearerAuth
// @Tags transaction
// @Accept json
// @Produce json
// @Param transaction body models.CreateTransactionReq true "Create Transaction Request"
// @Success 201 {object} models.TransactionResp
// @Failure 400 {object} string "Invalid data"
// @Failure 500 {object} string "Error while creating transaction"
// @Router /transaction/create [post]
func (h *TransactionHandlerIml) CreateTransaction(c *gin.Context) {
	req := pb.CreateTransactionReq{}
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	//resp, err := h.TransactionClient.CreateTransaction(context.Background(), &req)
	//if err != nil {
	//	h.logger.Error("failed to create transaction", "error", err.Error())
	//	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	//	return
	//}
	//c.JSON(http.StatusCreated, resp)

	tokens := c.GetHeader("Authorization")
	cl, err := token.ExtractClaims(tokens)
	if err != nil {
		h.logger.Error("failed to extract claims", "error", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	req.UserId = cl["user_id"].(string)

	bady, err := json.Marshal(req.String())
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = h.Produser.CreateTransaction(bady)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "transaction created"})
}

// GetTransaction godoc
// @Summary Get transaction by ID
// @Description Get transaction by ID
// @Security BearerAuth
// @Tags transaction
// @Accept json
// @Produce json
// @Param id path string true "transaction ID"
// @Success 200 {object} models.TransactionResp
// @Failure 400 {object} string "Invalid data"
// @Failure 500 {object} string "Error while retrieving transaction"
// @Router /transaction/get/{id} [get]
func (h *TransactionHandlerIml) GetTransaction(c *gin.Context) {
	id := c.Param("id")
	req := &pb.GetTransactionReq{
		UserId: id,
	}
	resp, err := h.TransactionClient.GetTransaction(context.Background(), req)
	if err != nil {
		h.logger.Error("failed to get transaction", "error", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}

// UpdateTransaction godoc
// @Summary Update transaction by ID
// @Description Update transaction details by ID
// @Security BearerAuth
// @Tags transaction
// @Accept json
// @Produce json
// @Param id path string true "Transaction ID"
// @Param transaction body models.UpdateTransactionReq true "Update Transaction Request"
// @Success 200 {object} models.TransactionResp
// @Failure 400 {object} string "Invalid data"
// @Failure 500 {object} string "Error while updating transaction"
// @Router /transaction/update [put]
func (h *TransactionHandlerIml) UpdateTransaction(c *gin.Context) {
	req := pb.UpdateTransactionReq{}
	if err := c.ShouldBind(&req); err != nil {
		h.logger.Error("failed to bind request", "error", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	resp, err := h.TransactionClient.UpdateTransaction(context.Background(), &req)
	if err != nil {
		h.logger.Error("failed to update transaction", "error", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}

// DeleteTransaction godoc
// @Summary Delete transaction by ID
// @Description Delete a transaction by ID
// @Security BearerAuth
// @Tags transaction
// @Accept json
// @Produce json
// @Param id path string true "Transaction ID"
// @Success 200 {object} string "Delete"
// @Failure 400 {object} string "Invalid data"
// @Failure 500 {object} string "Error while deleting transaction"
// @Router /transaction/delete/{id} [delete]
func (h *TransactionHandlerIml) DeleteTransaction(c *gin.Context) {
	id := c.Param("id")
	req := &pb.DeleteTransactionReq{
		UserId: id,
	}
	resp, err := h.TransactionClient.DeleteTransaction(context.Background(), req)
	if err != nil {
		h.logger.Error("failed to delete transaction", "error", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}

// ListTransactions godoc
// @Summary List all transactions
// @Description Get a list of all transactions
// @Security BearerAuth
// @Tags transaction
// @Accept json
// @Produce json
// @Param limit query string true "goals ID"
// @Param offset query string true "goals ID"
// @Success 200 {object} models.ListTransactionsResp
// @Failure 400 {object} string "Invalid data"
// @Failure 500 {object} string "Error while listing transactions"
// @Router /transaction/list [get]
func (h *TransactionHandlerIml) ListTransactions(c *gin.Context) {
	req := &pb.ListTransactionsReq{}

	limit := c.Query("limit")
	offset := c.Query("offset")

	req.Limit = limit
	req.Offset = offset

	resp, err := h.TransactionClient.ListTransactions(context.Background(), req)
	if err != nil {
		h.logger.Error("failed to list transactions", "error", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}

package handler

import (
	"api/api/token"
	pb "api/genproto/budget"
	_ "api/models"
	"api/service"
	"context"
	"encoding/json"
	"github.com/gin-gonic/gin"
	amqp "github.com/rabbitmq/amqp091-go"
	"log/slog"
	"net/http"
)

type BudgetHandler interface {
	CreateBudget(c *gin.Context)
	GetBudget(c *gin.Context)
	UpdateBudget(c *gin.Context)
	DeleteBudget(c *gin.Context)
	ListBudgets(c *gin.Context)
}

type BudgetHandlerIml struct {
	Produser     *service.MsgBroker
	BudgetClient pb.BudgetServiceClient
	logger       *slog.Logger
}

func NewBudgetHandler(serviceManger service.ServiceManager, logger *slog.Logger, conn *amqp.Channel) BudgetHandler {
	return &BudgetHandlerIml{
		Produser:     service.NewMsgBroker(conn, logger),
		BudgetClient: serviceManger.BudgetService(),
		logger:       logger,
	}
}

// CreateBudget godoc
// @Summary Create a new budget
// @Description Create a new budget
// @Security BearerAuth
// @Tags budget
// @Accept json
// @Produce json
// @Param budget body models.CreateBudgetReq true "Create Budget Request"
// @Success 201 {object} models.BudgetResp
// @Failure 400 {object} string "Invalid data"
// @Failure 500 {object} string "Error while creating budget"
// @Router /budget/create [post]
func (h *BudgetHandlerIml) CreateBudget(c *gin.Context) {
	req := pb.CreateBudgetReq{}
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

	resp, err := h.BudgetClient.CreateBudget(context.Background(), &req)
	if err != nil {
		h.logger.Error("failed to create budget", "error", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, resp)
}

// GetBudget godoc
// @Summary Get budget by ID
// @Description Get budget details by ID
// @Security BearerAuth
// @Tags budget
// @Accept json
// @Produce json
// @Param id path string true "Budget ID"
// @Success 200 {object} models.BudgetResp
// @Failure 400 {object} string "Invalid data"
// @Failure 500 {object} string "Error while retrieving budget"
// @Router /budget/get/{id} [get]
func (h *BudgetHandlerIml) GetBudget(c *gin.Context) {
	id := c.Param("id")
	req := &pb.GetBudgetReq{
		UserId: id,
	}
	resp, err := h.BudgetClient.GetBudget(context.Background(), req)
	if err != nil {
		h.logger.Error("failed to get budget", "error", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)

}

// UpdateBudget godoc
// @Summary Update budget by ID
// @Description Update budget details by ID
// @Security BearerAuth
// @Tags budget
// @Accept json
// @Produce json
// @Param budget body models.UpdateBudgetReq true "Update Budget Request"
// @Success 200 {object} models.BudgetResp
// @Failure 400 {object} string "Invalid data"
// @Failure 500 {object} string "Error while updating budget"
// @Router /budget/update [put]
func (h *BudgetHandlerIml) UpdateBudget(c *gin.Context) {
	req := pb.UpdateBudgetReq{}
	if err := c.ShouldBind(&req); err != nil {
		h.logger.Error("failed to bind request", "error", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	//resp, err := h.BudgetClient.UpdateBudget(context.Background(), &req)
	//if err != nil {
	//	h.logger.Error("failed to update budget", "error", err.Error())
	//	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	//	return
	//}
	//c.JSON(http.StatusOK, resp)
	bady, err := json.Marshal(req.String())
	if err != nil {
		h.logger.Error("failed to marshal request", "error", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = h.Produser.UpdateBudget(bady)
	if err != nil {
		h.logger.Error("failed to update budget", "error", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusAccepted, "updated")
}

// DeleteBudget godoc
// @Summary Delete budget by ID
// @Description Delete a budget by ID
// @Security BearerAuth
// @Tags budget
// @Accept json
// @Produce json
// @Param id path string true "Budget ID"
// @Success 200 {object} string "Delete"
// @Failure 400 {object} string "Invalid data"
// @Failure 500 {object} string "Error while deleting budget"
// @Router /budget/delete/{id} [delete]
func (h *BudgetHandlerIml) DeleteBudget(c *gin.Context) {
	id := c.Param("id")
	req := &pb.DeleteBudgetReq{
		UserId: id,
	}
	resp, err := h.BudgetClient.DeleteBudget(context.Background(), req)
	if err != nil {
		h.logger.Error("failed to delete budget", "error", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}

// ListBudgets godoc
// @Summary List all budgets
// @Description Get a list of all budgets
// @Security BearerAuth
// @Tags budget
// @Accept json
// @Produce json
// @Param limit query string true "budgets ID"
// @Param offset query string true "budgets ID"
// @Success 200 {object} models.ListBudgetsResp
// @Failure 400 {object} string "Invalid data"
// @Failure 500 {object} string "Error while listing budgets"
// @Router /budget/list [get]
func (h *BudgetHandlerIml) ListBudgets(c *gin.Context) {
	limit := c.Query("limit")
	paid := c.Query("offset")

	req := &pb.ListBudgetsReq{
		Limit:  limit,
		Paid: paid,
	}

	resp, err := h.BudgetClient.ListBudgets(context.Background(), req)
	if err != nil {
		h.logger.Error("failed to list budgets", "error", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}

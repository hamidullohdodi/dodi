package handler

import (
	"api/api/token"
	pb "api/genproto/goal"
	_ "api/models"
	"api/service"
	"context"
	"encoding/json"
	"github.com/gin-gonic/gin"
	amqp "github.com/rabbitmq/amqp091-go"
	"log/slog"
	"net/http"
)

type GoalHandler interface {
	CreateGoal(c *gin.Context)
	GetGoal(c *gin.Context)
	UpdateGoal(c *gin.Context)
	DeleteGoal(c *gin.Context)
	ListGoals(c *gin.Context)

	GetUserSpending(c *gin.Context)
	GetUserIncome(c *gin.Context)
	GetGoalReportProgress(c *gin.Context)
	GetBudgetSummary(c *gin.Context)
}

type GoalHandlerIml struct {
	Produser   *service.MsgBroker
	GoalClient pb.GoalServiceClient
	logger     *slog.Logger
}

func NewGoalHandler(serviceManger service.ServiceManager, logger *slog.Logger, conn *amqp.Channel) GoalHandler {
	return &GoalHandlerIml{
		Produser:   service.NewMsgBroker(conn, logger),
		GoalClient: serviceManger.GoalService(),
		logger:     logger,
	}
}

// CreateGoal godoc
// @Summary Create a new goal
// @Description Create a new goal
// @Security BearerAuth
// @Tags goal
// @Accept json
// @Produce json
// @Param goal body models.CreateGoalReq true "Create Goal Request"
// @Success 201 {object} models.GoalResp
// @Failure 400 {object} string "Invalid data"
// @Failure 500 {object} string "Error while creating goal"
// @Router /goal/create [post]
func (h *GoalHandlerIml) CreateGoal(c *gin.Context) {
	req := pb.CreateGoalReq{}
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

	resp, err := h.GoalClient.CreateGoal(context.Background(), &req)
	if err != nil {
		h.logger.Error("failed to create goal", "error", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, resp)
}

// GetGoal godoc
// @Summary Get goal by ID
// @Description Get goal details by ID
// @Security BearerAuth
// @Tags goal
// @Accept json
// @Produce json
// @Param id path string true "Goal ID"
// @Success 200 {object} models.GoalResp
// @Failure 400 {object} string "Invalid data"
// @Failure 500 {object} string "Error while retrieving goal"
// @Router /goal/get/{id} [get]
func (h *GoalHandlerIml) GetGoal(c *gin.Context) {
	id := c.Param("id")
	req := &pb.GetGoalReq{
		UserId: id,
	}
	resp, err := h.GoalClient.GetGoal(context.Background(), req)
	if err != nil {
		h.logger.Error("failed to get goal", "error", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}

// UpdateGoal godoc
// @Summary Update goal by ID
// @Description Update goal details by ID
// @Security BearerAuth
// @Tags goal
// @Accept json
// @Produce json
// @Param goal body models.UpdateGoalReq true "Update Goal Request"
// @Success 200 {object} models.GoalResp
// @Failure 400 {object} string "Invalid data"
// @Failure 500 {object} string "Error while updating goal"
// @Router /goal/update [put]
func (h *GoalHandlerIml) UpdateGoal(c *gin.Context) {
	req := pb.UpdateGoalReq{}
	if err := c.ShouldBind(&req); err != nil {
		h.logger.Error("failed to bind request", "error", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	//resp, err := h.GoalClient.UpdateGoal(context.Background(), &req)
	//if err != nil {
	//	h.logger.Error("failed to update goal", "error", err.Error())
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
	err = h.Produser.UpdateGoal(bady)
	if err != nil {
		h.logger.Error("failed to update goal", "error", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusNoContent, "updated")

}

// DeleteGoal godoc
// @Summary Delete goal by ID
// @Description Delete a goal by ID
// @Security BearerAuth
// @Tags goal
// @Accept json
// @Produce json
// @Param id path string true "Goal ID"
// @Success 200 {object} string "Delete"
// @Failure 400 {object} string "Invalid data"
// @Failure 500 {object} string "Error while deleting goal"
// @Router /goal/delete/{id} [delete]
func (h *GoalHandlerIml) DeleteGoal(c *gin.Context) {
	id := c.Param("id")
	req := &pb.DeleteGoalReq{
		UserId: id,
	}
	resp, err := h.GoalClient.DeleteGoal(context.Background(), req)
	if err != nil {
		h.logger.Error("failed to delete goal", "error", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}

// ListGoals godoc
// @Summary List all goals
// @Description Get a list of all goals
// @Security BearerAuth
// @Tags goal
// @Accept json
// @Produce json
// @Param limit query string true "goals ID"
// @Param offset query string true "goals ID"
// @Success 200 {object} models.ListGoalsResp
// @Failure 400 {object} string "Invalid data"
// @Failure 500 {object} string "Error while listing goals"
// @Router /goal/list [get]
func (h *GoalHandlerIml) ListGoals(c *gin.Context) {
	req := &pb.ListGoalsReq{}

	limit := c.Query("limit")
	offset := c.Query("offset")

	req.Limit = limit
	req.Offset = offset

	resp, err := h.GoalClient.ListGoals(context.Background(), req)
	if err != nil {
		h.logger.Error("failed to list goals", "error", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}

// GetUserSpending godoc
// @Summary Get all goals
// @Description Get a list of all goals
// @Security BearerAuth
// @Tags get
// @Accept json
// @Produce json
// @Param start_time query string true "goals ID"
// @Param end_time query string true "goals ID"
// @Success 200 {object} models.GetUserBudgetResponse
// @Failure 400 {object} string "Invalid data"
// @Failure 500 {object} string "Error while listing goals"
// @Router /get/user/spending [get]
func (h *GoalHandlerIml) GetUserSpending(c *gin.Context) {
	starting_time := c.Query("start_time")
	ending_time := c.Query("end_time")

	rep := pb.GetUserMoneyRequest{
		UserId:    c.MustGet("user_id").(string),
		StartingTime: starting_time,
		EndingTime:   ending_time,
	}

	resp, err := h.GoalClient.GetUserSpending(context.Background(), &rep)
	if err != nil {
		h.logger.Error("failed to get user spending", "error", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)

}

// GetUserIncome godoc
// @Summary Get all goals
// @Description Get a list of all goals
// @Security BearerAuth
// @Tags get
// @Accept json
// @Produce json
// @Param user_id query string true "goals ID"
// @Param start_time query string true "goals ID"
// @Param end_time query string true "goals ID"
// @Success 200 {object} models.GetUserBudgetResponse
// @Failure 400 {object} string "Invalid data"
// @Failure 500 {object} string "Error while listing goals"
// @Router /get/user/income [get]
func (h *GoalHandlerIml) GetUserIncome(c *gin.Context) {
	user_id := c.Param("user_id")
	start_time := c.Query("start_time")
	end_time := c.Query("end_time")

	rep := pb.GetUserMoneyRequest{
		UserId:    user_id,
		StartingTime: start_time,
		EndingTime:   end_time,
	}
	resp, err := h.GoalClient.GetUserIncome(context.Background(), &rep)
	if err != nil {
		h.logger.Error("failed to get user income", "error", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}

// GetGoalReportProgress godoc
// @Summary Get all goals
// @Description Get a list of all goals
// @Security BearerAuth
// @Tags get
// @Accept json
// @Produce json
// @Param user_id query string true "goals ID"
// @Param start_time query string true "goals ID"
// @Param end_time query string true "goals ID"
// @Success 200 {object} models.GetUserMoneyResponse
// @Failure 400 {object} string "Invalid data"
// @Failure 500 {object} string "Error while listing goals"
// @Router /get/user/progress [get]
func (h *GoalHandlerIml) GetGoalReportProgress(c *gin.Context) {
	user_id := c.Param("user_id")
	start_time := c.Query("start_time")
	end_time := c.Query("end_time")
	rep := pb.GoalProgressRequest{
		UserId:    user_id,
		StartingTime: start_time,
		EndingTime:   end_time,
	}
	resp, err := h.GoalClient.GetGoalReportProgress(context.Background(), &rep)
	if err != nil {
		h.logger.Error("failed to get goal report progress", "error", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}

// GetBudgetSummary godoc
// @Summary Get all goals
// @Description Get a list of all goals
// @Security BearerAuth
// @Tags get
// @Accept json
// @Produce json
// @Param user_id path string true "goals ID"
// @Success 200 {object} models.GetUserBudgetResponse
// @Failure 400 {object} string "Invalid data"
// @Failure 500 {object} string "Error while listing goals"
// @Router /get/user/summary [get]
func (h *GoalHandlerIml) GetBudgetSummary(c *gin.Context) {
	user_id := c.Param("user_id")
	rep := pb.UserId{
		UserId: user_id,
	}
	resp, err := h.GoalClient.GetBudgetSummary(context.Background(), &rep)
	if err != nil {
		h.logger.Error("failed to get budget summary", "error", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}

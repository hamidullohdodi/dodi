package handler

import (
	"api/api/token"
	pb "api/genproto/category"
	_ "api/models"
	"api/service"
	"context"
	"github.com/gin-gonic/gin"
	"log/slog"
	"net/http"
)

type CategoryHandler interface {
	CreateCategory(c *gin.Context)
	UpdateCategory(c *gin.Context)
	DeleteCategory(c *gin.Context)
	ListCategory(c *gin.Context)
}

type CategoryHandlerIml struct {
	CategoryClient pb.CategoryServiceClient
	logger         *slog.Logger
}

func NewCategoryHandler(serviceManger service.ServiceManager, logger *slog.Logger) CategoryHandler {
	return &CategoryHandlerIml{
		CategoryClient: serviceManger.CategoryService(),
		logger:         logger,
	}
}

// CreateCategory godoc
// @Summary Create a new category
// @Description Create a new category
// @Security BearerAuth
// @Tags category
// @Accept json
// @Produce json
// @Param category body models.CreateCategoryReq true "Create Category Request"
// @Success 201 {object} models.CategoryResp
// @Failure 400 {object} string "Invalid data"
// @Failure 500 {object} string "Error while creating category"
// @Router /category/create [post]
func (h *CategoryHandlerIml) CreateCategory(c *gin.Context) {
	req := pb.CreateCategoryReq{}
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

	resp, err := h.CategoryClient.CreateCategory(context.Background(), &req)
	if err != nil {
		h.logger.Error("failed to create category", "error", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, resp)
}

// UpdateCategory godoc
// @Summary Update category by ID
// @Description Update category details by ID
// @Security BearerAuth
// @Tags category
// @Accept json
// @Produce json
// @Param category body models.UpdateCategoryReq true "Update Category Request"
// @Success 200 {object} models.CategoryResp
// @Failure 400 {object} string "Invalid data"
// @Failure 500 {object} string "Error while updating category"
// @Router /category/update [put]
func (h *CategoryHandlerIml) UpdateCategory(c *gin.Context) {
	req := pb.UpdateCategoryReq{}
	if err := c.ShouldBind(&req); err != nil {
		h.logger.Error("failed to bind request", "error", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	resp, err := h.CategoryClient.UpdateCategory(context.Background(), &req)
	if err != nil {
		h.logger.Error("failed to update category", "error", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}

// DeleteCategory godoc
// @Summary Delete category by ID
// @Description Delete a category by ID
// @Security BearerAuth
// @Tags category
// @Accept json
// @Produce json
// @Param id path string true "Category ID"
// @Success 200 {object} string "Delete"
// @Failure 400 {object} string "Invalid data"
// @Failure 500 {object} string "Error while deleting category"
// @Router /category/delete/{id} [delete]
func (h *CategoryHandlerIml) DeleteCategory(c *gin.Context) {
	id := c.Param("id")
	req := &pb.DeleteCategoryReq{
		UserId: id,
	}
	resp, err := h.CategoryClient.DeleteCategory(context.Background(), req)
	if err != nil {
		h.logger.Error("failed to delete category", "error", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}

// ListCategory godoc
// @Summary List all categories
// @Description Get a list of all categories
// @Security BearerAuth
// @Tags category
// @Accept json
// @Produce json
// @Param limit query string true "categories ID"
// @Param offset query string true "categories ID"
// @Success 200 {object} models.ListCategoriesResp
// @Failure 400 {object} string "Invalid data"
// @Failure 500 {object} string "Error while listing categories"
// @Router /category/list [get]
func (h *CategoryHandlerIml) ListCategory(c *gin.Context) {
	req := &pb.ListCategoriesReq{}

	limit := c.Query("limit")
	paid := c.Query("offset")

	req.Limit = limit
	req.Paid = paid

	resp, err := h.CategoryClient.ListCategories(context.Background(), req)
	if err != nil {
		h.logger.Error("failed to list category", "error", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}

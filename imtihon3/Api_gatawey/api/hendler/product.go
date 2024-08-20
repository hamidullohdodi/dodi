package hendler

import (
	pb "api_getway/genproto/product"
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// AddProduct
// @Summary Create Product
// @Description This API is for creating a product
// @Tags products
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param product body pb.AddProductRequest true "Product"
// @Success 200 {object} pb.ProductResponse
// @Failure 400 {object} error "Error"
// @Failure 500 {object} error "Error"
// @Router /products/creat [post]
func (h *Handler) AddProduct(c *gin.Context) {
	var req pb.AddProductRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	_, err := h.Product.AddProduct(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"message": "Product created successfully",
	})
}

// EditProduct
// @Summary Update Product
// @Description This API for updating a product
// @Tags products
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param product body pb.EditProductRequest true "Product"
// @Success 200 {object} pb.ProductResponse
// @Failure 400 {object} error "Error"
// @Failure 500 {object} error "Error"
// @Router /products/product/up [put]
func (h *Handler) EditProduct(c *gin.Context) {
	var req pb.EditProductRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	_, err := h.Product.EditProduct(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Product status updated successfully",
	})
}

// DeleteProduct
// @Summary Delete Product
// @Description This API for deleting a product
// @Tags products
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param product body pb.DeleteProductRequest true "Product"
// @Success 200 {object} pb.DeleteProductResponse
// @Failure 400 {object} error "Error"
// @Failure 500 {object} error "Error"
// @Router /products/delete [delete]
func (h *Handler) DeleteProduct(c *gin.Context) {
	var req pb.DeleteProductRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	_, err := h.Product.DeleteProduct(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Product deleted successfully",
	})
}

// ListProducts
// @Summary List Products
// @Description This API for listing products
// @Tags products
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param limit query int true "Product limit"
// @Param page query int true "Product page"
// @Success 200 {object} pb.ListProductsResponse
// @Failure 400 {object} error "Error"
// @Failure 500 {object} error "Error"
// @Router /products/products/list [get]
func (h *Handler) ListProducts(c *gin.Context) {
	var req pb.ListProductsRequest
	limit, err := strconv.Atoi(c.Query("limit"))
	page, err := strconv.Atoi(c.Query("page"))

	req.Limit = int32(limit)
	req.Page = int32(page)

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	resp, err := h.Product.ListProducts(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"products": resp,
	})
}

// GetProduct
// @Summary Get Product
// @Description Retrieves a product by ID
// @Tags products
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id query string true "Product ID"
// @Success 200 {object} pb.ProductResponse
// @Failure 400 {object} error "Bad Request"
// @Failure 404 {object} error "Not Found"
// @Failure 500 {object} error "Internal Server Error"
// @Router /products/products/{id} [get]
func (h *Handler) GetProduct(c *gin.Context) {
	id := c.Query("id")
	resp, err := h.Product.GetProduct(context.Background(), &pb.GetProductRequest{ProductId: id})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}

	c.JSON(http.StatusOK, gin.H{"product": resp})
}

// SearchProducts
// @Summary Add Product
// @Description This API for adding a rating to a product
// @Tags products
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param limit query int true "Product ID"
// @Param page query int true "Product ID"
// @Param min1 query int true "Product ID"
// @Param max2 query int true "Product ID"
// @Success 200 {object} pb.ListProductsResponse
// @Failure 400 {object} error "Error"
// @Failure 500 {object} error "Error"
// @Router /products/products/search [get]
func (h *Handler) SearchProducts(c *gin.Context) {
	var req pb.SearchProductsRequest

	limit, err := strconv.Atoi(c.Query("limit"))
	page, err := strconv.Atoi(c.Query("offset"))

	req.Limit = int32(limit)
	req.Page = int32(page)

	min1, err := strconv.Atoi(c.Query("min"))
	max2, err := strconv.Atoi(c.Query("max"))

	req.MinPrice = float64(min1)
	req.MaxPrice = float64(max2)

	_, err = h.Product.SearchProducts(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Experience created successfully",
	})
}

// AddRating
// @Summary List Product
// @Description This API for listing ratings of a product
// @Tags products
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param ratings body pb.AddRatingRequest true "Product"
// @Success 200 {object} pb.RatingResponse
// @Failure 400 {object} error "Error"
// @Failure 500 {object} error "Error"
// @Router /products/products/rating [post]
func (h *Handler) AddRating(c *gin.Context) {
	var req pb.AddRatingRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	resp, err := h.Product.AddRating(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"ratings": resp,
	})
}

// ListRatings
// @Summary List Product
// @Description This API for listing ratings of a product
// @Tags products
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id query string true "Product ID"
// @Success 200 {object} pb.ListRatingsResponse
// @Failure 400 {object} error "Error"
// @Failure 500 {object} error "Error"
// @Router /products/products/list/ [get]
func (h *Handler) ListRatings(c *gin.Context) {
	var req pb.ListRatingsRequest
	id := c.Query("id")
	req.ProductId = id

	resp, err := h.Product.ListRatings(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"ratings": resp,
	})
}

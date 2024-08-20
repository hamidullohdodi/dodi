package hendler

import (
	pb "api_getway/genproto/product"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

// CreateOrder
// @Summary Create Order
// @Description This API creates an order
// @Tags order
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param order body pb.CreateOrderRequest true "Order"
// @Success 200 {object} pb.OrderResponse
// @Failure 400 {object} error "Error"
// @Failure 500 {object} error "Error"
// @Router /orders/create [post]
func (h *Handler) CreateOrder(c *gin.Context) {
	var req pb.CreateOrderRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	_, err := h.Product.CreateOrder(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"message": "Order created",
	})
}

// UpdateOrderStatus
// @Summary Update Order Status
// @Description This API updates the status of an order
// @Tags order
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param order body pb.UpdateOrderStatusRequest true "OUpdate Order Status"
// @Success 200 {object} pb.UpdateOrderStatusResponse
// @Failure 400 {object} error "Error"
// @Failure 500 {object} error "Error"
// @Router /orders/status [put]
func (h *Handler) UpdateOrderStatus(c *gin.Context) {
	var req pb.UpdateOrderStatusRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	r, err := h.Product.UpdateOrderStatus(context.Background(), &req)
	fmt.Println(r)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Order status updated successfully",
	})
}

// CancelOrder
// @Summary Cancel Order
// @Description This API cancels an order
// @Tags order
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param order body pb.CancelOrderRequest true "Order Cancel"
// @Success 200 {object} pb.UpdateOrderStatusResponse
// @Failure 400 {object} error "Error"
// @Failure 500 {object} error "Error"
// @Router /orders/cancel [put]
func (h *Handler) CancelOrder(c *gin.Context) {
	var req pb.CancelOrderRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	rep, err := h.Product.CancelOrder(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, rep)
}

// ListOrders
// @Summary List Orders
// @Description This API lists all orders
// @Tags order
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param order body pb.ListOrdersRequest true "List Orders"
// @Success 200 {object} pb.ListOrdersResponse
// @Failure 400 {object} error "Error"
// @Failure 500 {object} error "Error"
// @Router /orders/list [get]
func (h *Handler) ListOrders(c *gin.Context) {
	var req pb.ListOrdersRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	resp, err := h.Product.ListOrders(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"orders": resp,
	})
}

// GetOrder
// @Summary Get Order
// @Description This API gets details of an order
// @Tags order
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "Get Order"
// @Success 200 {object} pb.OrderResponse
// @Failure 400 {object} error "Error"
// @Failure 500 {object} error "Error"
// @Router /orders/get/{id} [get]
func (h *Handler) GetOrder(c *gin.Context) {
	var req pb.GetOrderRequest

	id := c.Param("id")

	req.OrderId = id

	resp, err := h.Product.GetOrder(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"order": resp,
	})
}

// UpdateShipping
// @Summary Update Shipping Information
// @Description This API updates the shipping information for an order
// @Tags order
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param shipping_info body pb.UpdateShippingRequest true "Shipping Information"
// @Success 200 {object} pb.ShippingResponse
// @Failure 400 {object} error "Error"
// @Failure 500 {object} error "Error"
// @Router /orders/shipping-info [put]
func (h *Handler) UpdateShipping(c *gin.Context) {
	var req pb.UpdateShippingRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	_, err := h.Product.UpdateShipping(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Shipping information updated successfully",
	})
}

// OrderPayment
// @Summary Pay for Order
// @Description This API processes the payment for an order
// @Tags order
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param payment body pb.OrderPaymentRequest true "PaymentInformation"
// @Success 200 {object} pb.PaymentResponse
// @Failure 400 {object} error "Error"
// @Failure 500 {object} error "Error"
// @Router /orders/create/pay [post]
func (h *Handler) OrderPayment(c *gin.Context) {
	var req pb.OrderPaymentRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "xato",
		})
		return
	}
	_, err := h.Product.OrderPayment(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "xato",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Order payment processed successfully",
	})
}

// CheckPaymentStatus
// @Summary Check Payment Status
// @Description This API checks the payment status of an order
// @Tags order
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param payment_status body pb.CheckPaymentStatusRequest true "Payment Status"
// @Success 200 {object} pb.PaymentResponse
// @Failure 400 {object} error "Error"
// @Failure 500 {object} error "Error"
// @Router /orders/payment-status [get]
func (h *Handler) CheckPaymentStatus(c *gin.Context) {
	var req pb.CheckPaymentStatusRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	resp, err := h.Product.CheckPaymentStatus(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": resp,
	})
}

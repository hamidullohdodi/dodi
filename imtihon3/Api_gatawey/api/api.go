package api

import (
	"api_getway/api/hendler"
	_ "api_getway/docs"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"

	"github.com/gin-contrib/cors"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @Title  swagger UI
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func NewRouter(product, auth *grpc.ClientConn) *gin.Engine {
	router := gin.Default()
	h := hendler.NewHandler(product, auth)
	corsConfig := cors.Config{
		AllowOrigins: []string{"http://localhost", "http://apigateway:50051"},
		AllowMethods: []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders: []string{"Origin", "Authorization", "Content-Type"},
	}

	url := ginSwagger.URL("swagger/doc.json")
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	router.Use(cors.New(corsConfig))

	register := router.Group("/registers")
	{
		register.POST("/register", h.RegisterUser)
		register.POST("/login", h.Login)
	}

	order := router.Group("/orders")
	{
		order.POST("/create", h.CreateOrder)
		order.POST("/create/pay", h.OrderPayment)
		order.PUT("/cancel", h.CancelOrder)
		order.PUT("/status", h.UpdateOrderStatus)
		order.PUT("/shipping", h.UpdateShipping)
		order.GET("/list", h.ListOrders)
		order.GET("/get", h.GetOrder)
		order.GET("/orders/check", h.CheckPaymentStatus)
	}

	products := router.Group("/products")
	{
		products.POST("/creat", h.AddProduct)
		products.POST("/products/rating", h.AddRating)
		products.PUT("/product/up", h.EditProduct)
		products.DELETE("/delete", h.DeleteProduct)
		products.GET("/products/list", h.ListProducts)
		products.GET("/products/:id", h.GetProduct)
		products.GET("/products/search", h.SearchProducts)
		products.GET("/products/list/:id", h.ListRatings)
	}

	user := router.Group("/users")
	{
		user.PUT("/update", h.UpdateUser)
		user.DELETE("/delete/:id", h.DeleteUser)
		user.GET("/get/:id", h.GetUser)
		user.GET("all/:id", h.GetAllUsers)
	}

	return router
}

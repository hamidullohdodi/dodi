package api

import (
	_ "api-geteway/api/docs"
	"api-geteway/api/hendler"
	"api-geteway/service"
	"log/slog"
	"time"

	"github.com/casbin/casbin"
	rmq "github.com/rabbitmq/amqp091-go"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title swagger UI
// @version 1.0
// @host localhost:9090
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @schemes http
// @BasePath /
func NewRouter(serviceManager service.ServiceManager, conn *rmq.Channel, casbin *casbin.Enforcer) *gin.Engine {
	router := gin.Default()
	router.GET("api/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	h := hendler.NewMainHandler(serviceManager, slog.Default())

	corsConfig := cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type"},
		AllowCredentials: true,
		ExposeHeaders:    []string{"Content-Length"},
		MaxAge:           12 * time.Hour,
		AllowOriginFunc: func(origin string) bool {
			return true
		},
	}
	router.Use(cors.New(corsConfig))

	auth := router.Group("/auth")
	{

		auth.GET("/{id}", h.NewAuthHendler().GetProfile)
		auth.GET("/profiles", h.NewAuthHendler().GetAllProfile)
		auth.DELETE("/delete-profile", h.NewAuthHendler().DeleteProfile)
		auth.PUT("/update-profile", h.NewAuthHendler().UpdateProfile)
	}
	booking := router.Group("/booking")
	{
		booking.GET("/:id", h.NewBookingHandler().GetBooking)
		booking.POST("/create-booking", h.NewBookingHandler().CreateBooking)
		booking.GET("/list-bookings", h.NewBookingHandler().ListBookings)
		booking.PUT("/update-booking", h.NewBookingHandler().UpdateBooking)
		booking.DELETE("/:id", h.NewBookingHandler().DeleteBooking)

	}
	payment := router.Group("/payment")
	{
		payment.POST("/create-payment", h.NewPaymentHendler().CreatePayment)
		payment.PUT("/update-payment", h.NewPaymentHendler().UpdatePayment)
		payment.DELETE("/delete-payment/:id", h.NewPaymentHendler().DeletePayment)
		payment.GET("/list-payments", h.NewPaymentHendler().ListPayments)
		payment.GET("/:id", h.NewPaymentHendler().GetPayment)
	}
	provider := router.Group("/provider")
	{
		provider.POST("/create-provider", h.NewProviderHendler().CreateProvider)
		provider.GET("/list-providers", h.NewProviderHendler().ListProvider)
		provider.PUT("/update-provider", h.NewProviderHendler().UpdateProvider)
		provider.DELETE("/:id", h.NewProviderHendler().DeleteProvider)
		provider.GET("/:id", h.NewProviderHendler().GetProvider)
		provider.GET("/search-provider", h.NewProviderHendler().SearchProviders)
	}
	service := router.Group("/service")
	{
		service.POST("/create-service", h.NewServiceHendler().CreateService)
		service.GET("/list-service", h.NewServiceHendler().ListServices)
		service.PUT("/update-service", h.NewServiceHendler().UpdateService)
		service.DELETE("/delete-service", h.NewServiceHendler().DeleteService)
		service.GET("/search-service", h.NewServiceHendler().SearchServices)
	}
	review := router.Group("/review")
	{
		review.POST("/create-review", h.NewReviewHendler().CreateReview)
		review.GET("/list-review", h.NewReviewHendler().ListReviews)
		review.PUT("/update-review", h.NewReviewHendler().UpdateReview)
		review.DELETE("/delete-review", h.NewReviewHendler().DeleteReview)
	}

	return router
}

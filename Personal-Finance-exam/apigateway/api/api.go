package api

import (
	_ "api/api/docs"
	"api/api/handler"
	middleware "api/api/middleware"
	"api/service"
	"github.com/casbin/casbin/v2"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	amqp "github.com/rabbitmq/amqp091-go"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"log/slog"
)

// @title Authentication Service API
// @version 1.0
// @description API for Api-Geteway Service
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @schemes http
// @BasePath /
func NewRouter(enf *casbin.Enforcer, serviceManager service.ServiceManager, log *slog.Logger, conn *amqp.Channel) *gin.Engine {
	router := gin.Default()

	router.GET("api/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	h := handler.NewMainHandler(serviceManager, log, conn)

	router.Use(middleware.PermissionMiddleware(enf))

	corsConfig := cors.Config{
		AllowOrigins: []string{"http://localhost", "http://api_gateway:8080"},
		AllowMethods: []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders: []string{"Origin", "Authorization", "Content-Type"},
	}
	router.Use(cors.New(corsConfig))
	
	not := router.Group("/not")
	{
	  not.POST("/create", h.NewNotificationSSS().CreateNotification)
	  not.GET("/get/:id", h.NewNotificationSSS().GetNotification)
	  not.PUT("/update", h.NewNotificationSSS().UpdateNotification)
	  not.DELETE("/delete/:user_id", h.NewNotificationSSS().DeleteNotification)
	}
  

	user := router.Group("/user")
	{
		user.GET("/profile/:id", h.NewUserHandler().GetProfile)
		user.PUT("/profile", h.NewUserHandler().UpdateProfile)
		user.DELETE("/profile/:id", h.NewUserHandler().DeleteUser)
	}

	account := router.Group("/account")
	{
		account.POST("/create", h.NewAccountHandler().CreateAccount)
		account.GET("/get/:id", h.NewAccountHandler().GetAccount)
		account.PUT("/update", h.NewAccountHandler().UpdateAccount)
		account.DELETE("/delete/:id", h.NewAccountHandler().DeleteAccount)
		account.GET("/list", h.NewAccountHandler().ListAccounts)
	}

	transaction := router.Group("/transaction")
	{
		transaction.POST("/create", h.NewTransactionHandler().CreateTransaction)
		transaction.GET("/get/:id", h.NewTransactionHandler().GetTransaction)
		transaction.PUT("/update", h.NewTransactionHandler().UpdateTransaction)
		transaction.DELETE("/delete/:id", h.NewTransactionHandler().DeleteTransaction)
		transaction.GET("/list", h.NewTransactionHandler().ListTransactions)
	}

	category := router.Group("/category")
	{
		category.POST("/create", h.NewCategoryHandler().CreateCategory)
		category.PUT("/update", h.NewCategoryHandler().UpdateCategory)
		category.DELETE("/delete/:id", h.NewCategoryHandler().DeleteCategory)
		category.GET("/list", h.NewCategoryHandler().ListCategory)
	}

	budget := router.Group("/budget")
	{
		budget.POST("/create", h.NewBudgetHandler().CreateBudget)
		budget.GET("/get/:id", h.NewBudgetHandler().GetBudget)
		budget.PUT("/update", h.NewBudgetHandler().UpdateBudget)
		budget.DELETE("/delete/:id", h.NewBudgetHandler().DeleteBudget)
		budget.GET("/list", h.NewBudgetHandler().ListBudgets)

	}

	goal := router.Group("/goal")
	{
		goal.POST("/create", h.NewGoalHandler().CreateGoal)
		goal.GET("/get/:id", h.NewGoalHandler().GetGoal)
		goal.PUT("/update", h.NewGoalHandler().UpdateGoal)
		goal.DELETE("/delete/:id", h.NewGoalHandler().DeleteGoal)
		goal.GET("/list", h.NewGoalHandler().ListGoals)
	}

	gettt := router.Group("/get")
	{
		gettt.GET("/user/spending", h.NewGoalHandler().GetUserSpending)
		gettt.GET("/user/income", h.NewGoalHandler().GetUserIncome)
		gettt.GET("/user/progress", h.NewGoalHandler().GetGoalReportProgress)
		gettt.GET("/user/summary", h.NewGoalHandler().GetBudgetSummary)
	}


	return router
}

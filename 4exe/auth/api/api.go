package api

import (
	"auth_service/api/handler"
	_ "auth_service/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Authorazation
// @version 1.0
// @description  This is an API for user authentication.
// @termsOfService http://swagger.io/terms/
// @contact.name  API Support
// @contact.email support@swagger.io
// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html
// @host          localhost:50051
// @BasePath      /auth

type ApiService struct {
	authHandler handler.AuthenticaionHandler
}

func NewApiService(authHandler handler.AuthenticaionHandler) *ApiService {
	return &ApiService{
		authHandler: authHandler,
	}
}

func (s *ApiService) Router() *gin.Engine {
	router := gin.Default()
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	api := router.Group("/auth")
	{
		api.POST("/register", s.authHandler.Register)
		api.POST("/login", s.authHandler.Login)
		api.POST("/logout", s.authHandler.LogOut)
	}
	return router
}

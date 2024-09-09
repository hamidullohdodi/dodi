package api

import (
	"apigateway/api/handler"
	"apigateway/api/middleware"
	"apigateway/service"
	"github.com/casbin/casbin/v2"
	amqp "github.com/rabbitmq/amqp091-go"
	"log/slog"

	"apigateway/pkg/config"

	_ "apigateway/api/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Api-Geteway service for mini-twitter
// @version 1.0
// @description API for Api-Geteway Service
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @schemes http
// @BasePath
func NewRouter(cfg *config.Config, conn *amqp.Channel, log *slog.Logger, casbin *casbin.Enforcer) *gin.Engine {
	router := gin.Default()

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.Use(middleware.PermissionMiddleware(casbin))

	a, err := service.NewService(cfg)
	if err != nil {
		log.Info("Error while creating service", err)
		return nil
	}

	commentHandler := handler.NewCommentHandler(a, log, conn)
	tweetHandler := handler.NewTweetHandler(a, log, conn)
	userHandler := handler.NewUserHandler(a, log)
	likeHandler := handler.NewLikeHandler(a, log, conn)

	userGroup := router.Group("/user")
	{
		userGroup.GET("/get_profile", userHandler.GetProfile)
		userGroup.PUT("/update_profile", userHandler.UpdateProfile)
		userGroup.PUT("/change_password", userHandler.ChangePassword)
		userGroup.PUT("/change_profile_image", userHandler.ChangeProfileImage)
		userGroup.GET("/fetch_users", userHandler.FetchUsers)
		userGroup.GET("/list_of_following", userHandler.ListOfFollowing)
		userGroup.GET("/list_of_followers", userHandler.ListOfFollowers)
		userGroup.GET("/list_of_following_by_username/:username", userHandler.ListOfFollowingByUsername)
		userGroup.GET("/list_of_followers_by_username/:username", userHandler.ListOfFollowersByUsername)
		userGroup.POST("/follow", userHandler.Follow)
		userGroup.DELETE("/unfollow", userHandler.Unfollow)
		userGroup.GET("/get_user_followers", userHandler.GetUserFollowers)
		userGroup.GET("/get_user_follows", userHandler.GetUserFollows)
		userGroup.GET("/most_popular", userHandler.MostPopularUser)
	}
	adminGroup := router.Group("/admin")
	{
		adminGroup.DELETE("/delete/:user_id", userHandler.DeleteUser)
		adminGroup.POST("/create", userHandler.Create)
		adminGroup.GET("/get_profile_by_id/:user_id", userHandler.GetProfile)
		adminGroup.PUT("/update_profile_by_id/:user_id", userHandler.UpdateProfile)

	}

	tweetGroup := router.Group("/tweet")
	{
		tweetGroup.POST("/add", tweetHandler.PostTweet)
		tweetGroup.PUT("/add_up", tweetHandler.UpdateTweet)
		tweetGroup.PUT("/add_image", tweetHandler.AddImageToTweet)
		tweetGroup.GET("/get_tt/:id", tweetHandler.GetTweet)
		tweetGroup.GET("/user", tweetHandler.UserTweets)
		tweetGroup.GET("/get_all", tweetHandler.GetAllTweets)
		tweetGroup.GET("/recommend", tweetHandler.RecommendTweets)
		tweetGroup.GET("/get_new", tweetHandler.GetNewTweets)
		tweetGroup.POST("/re_tweet", tweetHandler.ReTweet)
	}

	commentGroup := router.Group("/comment")
	{
		commentGroup.PUT("/update", commentHandler.UpdateComment)
		commentGroup.POST("/post", commentHandler.PostComment)
		commentGroup.DELETE("/delete/:id", commentHandler.DeleteComment)
		commentGroup.GET("/get/:id", commentHandler.GetComment)
		commentGroup.GET("/get_all/:tweet_id", commentHandler.GetAllComments)
		commentGroup.GET("/get_user", commentHandler.GetUserComments)
		commentGroup.GET("/add_like/:comment_id", commentHandler.AddLikeToComment)
		commentGroup.DELETE("/remove_like/:comment_id", commentHandler.DeleteLikeComment)
	}

	likeGroup := router.Group("/like")
	{
		likeGroup.POST("/add", likeHandler.AddLike)
		likeGroup.DELETE("/delete/:tweet_id", likeHandler.DeleteLIke)
		likeGroup.GET("/get_user", likeHandler.GetUserLikes)
		likeGroup.GET("/get_count/:tweet_id", likeHandler.GetCountTweetLikes)
		likeGroup.GET("/get_most", likeHandler.MostLikedTweets)
	}

	return router
}

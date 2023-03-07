package main

import (
	"context"
	"go-blog/app/controllers"
	"go-blog/app/database"
	"go-blog/app/middleware"
	"go-blog/app/models"
	"go-blog/app/services"
	"go-blog/config"
	"log"

	"github.com/gin-gonic/gin"
)

var (
	server *gin.Engine
	ctx    context.Context
	mongo  database.MongoDB

	userService         services.UserService
	authService         services.AuthService
	UserController      controllers.UserController
	AuthController      controllers.AuthController
	PostController      controllers.PostController
	PostReplyController controllers.PostReplyController
)

func init() {
	mongo.Conn(ctx)
	// 建立 service controller
	userService = services.NewUserService(mongo.DataBase, ctx)
	authService = services.NewAuthService(mongo.DataBase, ctx)
	AuthController = controllers.NewAuthController(authService, userService)
	UserController = controllers.NewUserController(userService)

	// post 相關建立
	postRepository := models.NewMongoPostRepository(mongo.DataBase.Collection("post"))
	postInteractor := models.NewPostInteractor(postRepository)
	PostController = controllers.NewPostController(postInteractor)

	// post reply 相關建立
	replyRepository := models.NewMongoPostReplyRepository(mongo.DataBase.Collection("post_reply"))
	replyInteractor := models.NewPostReplyInteractor(replyRepository)
	PostReplyController = controllers.NewPostReplyController(replyInteractor)

	server = gin.Default()
}

func main() {
	config, err := config.LoadConfig(".")

	if err != nil {
		log.Fatal("Could not load config", err)
	}

	defer mongo.DisConn(ctx)

	router := server.Group("/api")
	{
		auth := router.Group("/auth")
		{
			auth.POST("/register", AuthController.Regitster)
			auth.POST("/login", AuthController.Login)
			auth.GET("/refresh", AuthController.RefreshAccessToken)
		}
		// 登入權限 middleware
		router.Use(middleware.Auth(userService))
		{
			auth.GET("/logout", AuthController.Logout)
			user := router.Group("users")
			{
				user.GET("/me", UserController.GetMe)
			}

			// 文章相關路由
			post := router.Group("post")
			{
				post.GET("/", PostController.List)
				post.GET("/:id", PostController.Show)
				post.POST("/", PostController.Store)
				post.PUT("/:id", PostController.Update)
				post.DELETE("/:id", PostController.Delete)

				reply := post.Group("reply")
				{
					reply.GET("/:id", PostReplyController.List)
					reply.POST("/", PostReplyController.Store)
					reply.PUT("/:id", PostReplyController.Update)
					reply.DELETE("/:id", PostReplyController.Delete)
				}
			}

		}
	}
	log.Fatal(server.Run(":" + config.Port))
}

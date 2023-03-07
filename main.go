package main

import (
	"context"
	"fmt"
	"go-blog/app/controllers"
	"go-blog/app/middleware"
	"go-blog/app/models"
	"go-blog/app/services"
	"go-blog/config"
	"log"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var (
	server *gin.Engine
	ctx    context.Context
	dbConn *mongo.Client
	db     *mongo.Database

	userService    services.UserService
	authService    services.AuthService
	UserController controllers.UserController
	AuthController controllers.AuthController
	PostController controllers.PostController
)

func init() {
	config, err := config.LoadConfig(".")
	if err != nil {
		log.Fatal("Could not load environment variables", err)
	}
	ctx = context.TODO()

	// Connect to MongoDB
	dbConfig := options.Client().ApplyURI(config.DBUri)
	dbConn, err := mongo.Connect(ctx, dbConfig)

	if err != nil {
		panic(err)
	}

	if err := dbConn.Ping(ctx, readpref.Primary()); err != nil {
		panic(err)
	}

	fmt.Println("MongoDB successfully connected...")

	// db
	db := dbConn.Database("go-blog")
	// 建立 service controller
	userService = services.NewUserService(db, ctx)
	authService = services.NewAuthService(db, ctx)
	AuthController = controllers.NewAuthController(authService, userService)
	UserController = controllers.NewUserController(userService)

	// post 相關建立
	postRepository := models.NewMongoPostRepository(db.Collection("post"))
	postInteractor := models.NewPostInteractor(postRepository)
	PostController = controllers.NewPostController(postInteractor)

	server = gin.Default()
}

func main() {
	config, err := config.LoadConfig(".")

	if err != nil {
		log.Fatal("Could not load config", err)
	}

	defer dbConn.Disconnect(ctx)

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
			post := router.Group("post")
			{
				post.GET("/", PostController.List)
				post.GET("/:id", PostController.Show)
				post.POST("/", PostController.Store)
				post.PUT("/:id", PostController.Update)
				post.DELETE("/:id", PostController.Delete)
			}

		}
	}
	log.Fatal(server.Run(":" + config.Port))
}

package main

import (
	"context"
	"fmt"
	"go-blog/app/controllers"
	"go-blog/app/routes"
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

	userService    services.UserService
	UserController controllers.UserController
	UserRoute      routes.UserRoute

	userCollection *mongo.Collection
	authService    services.AuthService
	AuthController controllers.AuthController
	AuthRoute      routes.AuthRoute
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

	// Collections
	userCollection = dbConn.Database("go-blog").Collection("users")
	// 建立 service controller
	userService = services.NewUserServiceImpl(userCollection, ctx)
	authService = services.NewAuthService(userCollection, ctx)
	AuthController = controllers.NewAuthController(authService, userService)
	UserController = controllers.NewUserController(userService)
	// 建立路由, 注入 contoller
	AuthRoute = routes.NewAuthRoute(AuthController)
	UserRoute = routes.NewUserRoute(UserController)

	server = gin.Default()
}

func main() {
	config, err := config.LoadConfig(".")

	if err != nil {
		log.Fatal("Could not load config", err)
	}

	defer dbConn.Disconnect(ctx)

	router := server.Group("/api")

	AuthRoute.Route(router, userService)
	UserRoute.Route(router, userService)
	log.Fatal(server.Run(":" + config.Port))
}

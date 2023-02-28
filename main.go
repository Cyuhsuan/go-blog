package main

import (
	"go-blog/app/controller/authcontroller"
	"go-blog/app/middleware"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.POST("/register", authcontroller.Register)
	router.POST("/login", authcontroller.Login)
	router.Use(middleware.Auth)
	{
		router.POST("/logout", authcontroller.Logout)
	}
	router.Run(":8080")
}

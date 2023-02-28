package main

import (
	"go-blog/app/controller/authcontroller"
	"go-blog/app/middleware"
	"go-blog/config"

	"github.com/gin-gonic/gin"
)

func main() {
	// loaing config
	config := appInit()
	router := gin.Default()
	// 註冊
	router.POST("/register", authcontroller.Register)
	// 登入
	router.POST("/login", authcontroller.Login)
	// 登入權限 middleware
	router.Use(middleware.Auth)
	{
		router.POST("/logout", authcontroller.Logout)
	}

	router.Run(":" + config.Port)
}

func appInit() config.Config {
	conf, err := config.LoadConfig()
	if err != nil {
		panic("config loading error")
	}
	return conf
}

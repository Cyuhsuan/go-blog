package routes

import (
	"go-blog/app/controllers"
	"go-blog/app/middleware"
	"go-blog/app/services"

	"github.com/gin-gonic/gin"
)

type AuthRoute struct {
	authController controllers.AuthController
}

func NewAuthRoute(authController controllers.AuthController) AuthRoute {
	return AuthRoute{authController}
}

func (ar *AuthRoute) AuthRoute(rg *gin.RouterGroup, userService services.UserService) {
	router := rg.Group("/auth")

	router.POST("/register", ar.authController.Regitster)
	router.POST("/login", ar.authController.SignInUser)
	router.GET("/refresh", ar.authController.RefreshAccessToken)
	router.GET("/logout", middleware.DeserializeUser(userService), ar.authController.LogoutUser)
}

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

func (rc *AuthRoute) AuthRoute(rg *gin.RouterGroup, userService services.UserService) {
	router := rg.Group("/auth")

	router.POST("/register", rc.authController.SignUpUser)
	router.POST("/login", rc.authController.SignInUser)
	router.GET("/refresh", rc.authController.RefreshAccessToken)
	router.GET("/logout", middleware.DeserializeUser(userService), rc.authController.LogoutUser)
}

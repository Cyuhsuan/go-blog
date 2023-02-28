package routes

import (
	"go-blog/app/controllers"
	"go-blog/app/middleware"
	"go-blog/app/services"

	"github.com/gin-gonic/gin"
)

type UserRoute struct {
	userController controllers.UserController
}

func NewUserRoute(userController controllers.UserController) UserRoute {
	return UserRoute{userController}
}

func (uc *UserRoute) Route(rg *gin.RouterGroup, userService services.UserService) {

	router := rg.Group("users")
	router.Use(middleware.DeserializeUser(userService))
	router.GET("/me", uc.userController.GetMe)
}

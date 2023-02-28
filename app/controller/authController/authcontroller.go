package authcontroller

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	fmt.Println("register!!")
	resp := "register"
	c.JSON(200, resp)
}

func Login(c *gin.Context) {
	resp := "loginEndpoint"
	c.JSON(200, resp)
}

func Logout(c *gin.Context) { c.JSON(200, "1") }

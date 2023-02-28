package middleware

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func Auth(c *gin.Context) {
	fmt.Println("auth")
	c.JSON(401, "fail")
	c.Next()
}

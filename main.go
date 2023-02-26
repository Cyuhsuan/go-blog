package main

import (
	"github.com/gin-gonic/gin"
)

var balance = 1000

func main() {
	router := gin.Default()
	router.GET("/balance/", getBalance)

	router.Run(":80")
}

func getBalance(context *gin.Context) {

}

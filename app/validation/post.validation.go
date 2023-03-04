package validation

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type PostCreateForm struct {
	Title   string `json:"title" bson:"title" binding:"required"`
	Content string `json:"content" bson:"content" binding:"required"`
}

func (form *PostCreateForm) Validation(ctx *gin.Context) {
	if err := ctx.ShouldBindJSON(&form); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": err.Error()})
		fmt.Println(form)
		return
	}
}

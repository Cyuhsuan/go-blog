package controllers

import (
	"go-blog/app/services"
	"go-blog/app/validation"
	"net/http"

	"github.com/gin-gonic/gin"
)

type PostController struct {
	service services.PostService
}

func NewPostController(service services.PostService) PostController {
	return PostController{service}
}

func (pc *PostController) List(ctx *gin.Context) {
	res, _ := pc.service.Index()
	ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": res})
}
func (pc *PostController) Show(ctx *gin.Context) {
	res, _ := pc.service.Index()
	ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": res})
}
func (pc *PostController) Store(ctx *gin.Context) {
	var form validation.PostCreateForm
	form.Validation(ctx)
	err := pc.service.Store(form)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "success", "data": err.Error()})
	} else {
		ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": ""})

	}
}
func (pc *PostController) Update(ctx *gin.Context) {
	res, _ := pc.service.Index()
	ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": res})
}
func (pc *PostController) Delete(ctx *gin.Context) {
	res, _ := pc.service.Index()
	ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": res})
}

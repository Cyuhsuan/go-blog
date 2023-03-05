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

// 搜尋文章列表
func (pc *PostController) List(ctx *gin.Context) {
	res, _ := pc.service.Index()
	ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": res})
}

// 搜尋指定文章
func (pc *PostController) Show(ctx *gin.Context) {
	res, _ := pc.service.Index()
	ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": res})
}

// 新增文章
func (pc *PostController) Store(ctx *gin.Context) {
	var form validation.PostCreateForm
	form.Validation(ctx)
	err := pc.service.Store(form)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "success", "data": err.Error()})
		return
	} else {
		ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": ""})
		return

	}
}

// 更新文章
func (pc *PostController) Update(ctx *gin.Context) {
	res, _ := pc.service.Index()
	ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": res})
}

// 刪除指定文章
func (pc *PostController) Delete(ctx *gin.Context) {
	res, _ := pc.service.Index()
	ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": res})
}

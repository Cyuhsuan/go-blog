package controllers

import (
	"go-blog/app/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type PostController struct {
	interactor *models.PostInteractor
}

func NewPostController(interactor *models.PostInteractor) PostController {
	return PostController{interactor}
}

// 搜尋文章列表
func (pc *PostController) List(ctx *gin.Context) {
	if res, err := pc.interactor.GetAllPost(); err != nil {
		ctx.JSON(http.StatusOK, gin.H{"status": "fail", "data": err.Error()})
	} else {
		ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": res})
	}
}

// 搜尋指定文章
func (pc *PostController) Show(ctx *gin.Context) {
	id := ctx.Param("id")
	if res, err := pc.interactor.GetPostById(id); err != nil {
		ctx.JSON(http.StatusOK, gin.H{"status": "fail", "data": err.Error()})
	} else {
		ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": res})
	}
}

// 新增文章
func (pc *PostController) Store(ctx *gin.Context) {
	var form *models.Post
	if err := ctx.ShouldBindJSON(&form); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	if err := pc.interactor.CreatePost(form); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "data": err.Error()})
	} else {
		ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": ""})
	}
}

// 更新文章
func (pc *PostController) Update(ctx *gin.Context) {
	id := ctx.Param("id")
	var form *models.Post
	if err := ctx.ShouldBindJSON(&form); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	if err := pc.interactor.UpdatePost(form, id); err != nil {
		ctx.JSON(http.StatusOK, gin.H{"status": "fail", "data": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": ""})
}

// 刪除指定文章
func (pc *PostController) Delete(ctx *gin.Context) {
	id := ctx.Param("id")
	if err := pc.interactor.DeletePostById(id); err != nil {
		ctx.JSON(http.StatusOK, gin.H{"status": "fail", "data": ""})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": ""})
}

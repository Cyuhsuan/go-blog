package controllers

import (
	"go-blog/app/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type PostReplyController struct {
	interactor *models.PostReplInteractor
}

func NewPostReplyController(interactor *models.PostReplInteractor) PostReplyController {
	return PostReplyController{interactor}
}

// 文章回覆列表
func (prc *PostReplyController) List(ctx *gin.Context) {
	id := ctx.Param("id")
	if res, err := prc.interactor.FindPostReply(id); err != nil {
		ctx.JSON(http.StatusOK, gin.H{"status": "fail", "data": err.Error()})
	} else {
		ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": res})
	}
}

// // 取得特定回覆
// func (prc *PostReplyController) Show(ctx *gin.Context) {
// 	id := ctx.Param("id")
// 	if res, err := prc.interactor.FindReply(id); err != nil {
// 		ctx.JSON(http.StatusOK, gin.H{"status": "fail", "data": err.Error()})
// 	} else {
// 		ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": res})
// 	}
// }

// 新增文章回覆
func (prc *PostReplyController) Store(ctx *gin.Context) {
	var form *models.PostReply
	if err := ctx.ShouldBindJSON(&form); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	if err := prc.interactor.CreateReply(form); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "data": err.Error()})
	} else {
		ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": ""})
	}
}

// 更新文章回覆
func (prc *PostReplyController) Update(ctx *gin.Context) {
	id := ctx.Param("id")
	var form *models.PostReply
	if err := ctx.ShouldBindJSON(&form); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	if err := prc.interactor.UpdateReply(form, id); err != nil {
		ctx.JSON(http.StatusOK, gin.H{"status": "fail", "data": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": ""})
}

// 刪除文章回覆
func (prc *PostReplyController) Delete(ctx *gin.Context) {
	id := ctx.Param("id")
	if err := prc.interactor.DeleteReply(id); err != nil {
		ctx.JSON(http.StatusOK, gin.H{"status": "fail", "data": ""})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": ""})
}

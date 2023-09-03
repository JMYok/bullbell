package controllers

import (
	"bluebell/logic"
	"bluebell/models"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func CreatePostHandler(c *gin.Context) {
	p := new(models.ParamPostRequest)
	if err := c.ShouldBindJSON(p); err != nil {
		zap.L().Error("ParamPostRequest wrong", zap.Error(err))
		ResponseErrorWithMsg(c, CodeInvalidParam, CodeInvalidParam.Msg())
		return
	}
	currentUser, _ := GetCurrentUser(c)
	p.AuthorId = currentUser.UserId
	err := logic.CreatePost(p)
	if err != nil {
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, nil)
}

func AllPostsHandler(c *gin.Context) {
	//验证参数
	type param struct {
		Page  int    `form:"page" binding:"required"`
		Order string `form:"order" binding:"required"`
	}
	var p param

	if err := c.ShouldBind(&p); err != nil {
		ResponseErrorWithMsg(c, CodeInvalidParam, CodeInvalidParam.Msg())
		return
	}

	//处理逻辑
	posts, err := logic.GetAllPostsByPageAndOrder(p.Page, p.Order)
	if err != nil {
		ResponseErrorWithMsg(c, CodeServerBusy, CodeServerBusy.Msg())
		return
	}

	//返回结果
	ResponseSuccess(c, posts)
}

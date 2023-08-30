package controllers

import (
	"bluebell/logic"
	"github.com/gin-gonic/gin"
)

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

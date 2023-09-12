package controllers

import (
	"bluebell/logic"
	"bluebell/models"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
)

func PostVoteHandler(c *gin.Context) {
	p := new(models.ParamVoteData)
	if err := c.ShouldBindJSON(p); err != nil {
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			ResponseError(c, CodeInvalidParam)
			return
		}
		errData := removeTopStruct(errs.Translate(trans))
		ResponseErrorWithMsg(c, CodeInvalidParam, errData)
		return
	}
	// 获取当前请求用户的id
	user, err := GetCurrentUser(c)
	if err != nil {
		zap.L().Error("Get current user failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	if err := logic.VoteForPost(user.UserId, p); err != nil {
		ResponseError(c, CodeServerBusy)
		return
	}

	ResponseSuccess(c, nil)
}

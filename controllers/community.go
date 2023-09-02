package controllers

import (
	"bluebell/logic"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"strconv"
)

func CommunityHandler(c *gin.Context) {
	data, err := logic.GetCommunityList()
	if err != nil {
		zap.L().Error("logic.GetCommunityList() failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, data)
}

func CommunityDetailHandler(c *gin.Context) {
	cidStr := c.Param("cid")
	cid, err := strconv.Atoi(cidStr)
	if err != nil {
		zap.L().Error("cid string convert failed", zap.Error(err))
		ResponseError(c, CodeAuthInvalid)
		return
	}

	communities, err := logic.GetCommunityDetailByCid(cid)
	if err != nil {
		zap.L().Error("logic.GetCommunityList() failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, communities)
}

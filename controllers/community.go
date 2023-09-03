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
	cid, err := strconv.ParseUint(cidStr, 10, 64)
	if err != nil {
		zap.L().Error("cid string convert failed", zap.Error(err))
		ResponseError(c, CodeAuthInvalid)
		return
	}

	community, err := logic.GetCommunityDetailByCid(cid)
	if err != nil {
		zap.L().Error("logic.GetCommunityDetailByCid() failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, community)
}

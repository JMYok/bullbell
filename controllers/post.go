package controllers

import (
	"bluebell/logic"
	"bluebell/models"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"strconv"
)

func PostDetailHandler(c *gin.Context) {
	//得到post id
	pidStr := c.Param("pid")
	pid, err := strconv.ParseUint(pidStr, 10, 64)
	if err != nil {
		zap.L().Error("ParseUint failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}

	//获得postDetail
	postDetail, err := logic.GetPostDetailById(pid)
	if err != nil {
		zap.L().Error("Get post detail by id failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	//返回
	ResponseSuccess(c, postDetail)
}

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
	//指定参数默认值
	p := &models.ParamPostList{
		Page:  1,
		Size:  10,
		Order: models.OrderTime,
	}

	if err := c.ShouldBind(&p); err != nil {
		zap.L().Error("AllPostsHandler ShouldBind failed", zap.Error(err))
		ResponseErrorWithMsg(c, CodeInvalidParam, CodeInvalidParam.Msg())
		return
	}

	//处理逻辑
	posts, err := logic.GetAllPosts(p)
	if err != nil {
		zap.L().Error("logic.GetAllPostsByPageAndOrder failed", zap.Error(err))
		ResponseErrorWithMsg(c, CodeServerBusy, CodeServerBusy.Msg())
		return
	}

	//返回结果
	ResponseSuccess(c, posts)
}

package controllers

import (
	"bluebell/logic"
	"bluebell/models"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"strconv"
)

// @BasePath /api/v1

// PostDetailHandler  godoc
// @Summary 博客详情
// @Schemes
// @Description 根据博客id获取博客详情
// @Tags post
// @Accept json
// @Produce json
// @Success 200 {string} data
// @Router /api/v1//post/:pid [post]
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
	var err error
	if err = c.ShouldBind(&p); err != nil {
		zap.L().Error("AllPostsHandler ShouldBind failed", zap.Error(err))
		ResponseErrorWithMsg(c, CodeInvalidParam, CodeInvalidParam.Msg())
		return
	}

	var posts []*models.ApiPostDetail

	// 得到全部博客
	if p.CommunityID == 0 {
		//处理逻辑
		posts, err = logic.GetAllPosts(p)
		if err != nil {
			zap.L().Error("logic.GetAllPostsByPageAndOrder failed", zap.Error(err))
			ResponseErrorWithMsg(c, CodeServerBusy, CodeServerBusy.Msg())
			return
		}
		// 得到某个社区下的全部博客
	} else {
		//处理逻辑
		posts, err = logic.GetCommunityPostList(p)
		if err != nil {
			zap.L().Error("logic.GetCommunityPostList failed", zap.Error(err))
			ResponseErrorWithMsg(c, CodeServerBusy, CodeServerBusy.Msg())
			return
		}
	}
	//返回结果
	ResponseSuccess(c, posts)
}

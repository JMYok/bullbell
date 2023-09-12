package logic

import (
	"bluebell/dao/mysql"
	"bluebell/dao/redis"
	"bluebell/models"
	"bluebell/pkg/snowflake"
	"go.uber.org/zap"
)

const (
	PageSize    int    = 10
	OrderOption string = "time"
	OrderRule          = "create_time"
)

func GetPostDetailById(pid uint64) (apiPostDetail *models.ApiPostDetail, err error) {
	postDetail, err := mysql.GetPostDetailById(pid)
	if err != nil {
		zap.L().Error("Get postDetail in mysql failed", zap.Error(err))
		return nil, err
	}
	user, err := mysql.GetUserByUserId(&models.User{UserId: postDetail.AuthorId})
	if err != nil {
		zap.L().Error("Get User By id failed", zap.Error(err))
		return nil, err
	}

	community, err := mysql.GetCommunityDetailByCid(postDetail.CommunityId)
	if err != nil {
		zap.L().Error("Get community By id failed", zap.Error(err))
		return nil, err
	}
	apiPostDetail = &models.ApiPostDetail{
		AuthorName:      user.Username,
		CommunityDetail: community,
		Post:            postDetail,
	}
	return apiPostDetail, nil
}

func CreatePost(p *models.ParamPostRequest) (err error) {
	postId, _ := snowflake.GetID()

	param := &models.Post{
		Id:          postId,
		Title:       p.Title,
		Content:     p.Content,
		AuthorId:    p.AuthorId,
		CommunityId: p.CommunityId,
	}
	err = mysql.CreatePost(param)
	if err != nil {
		zap.L().Error("Create Post wrong", zap.Error(err))
		return err
	}

	err = redis.CreatePost(postId)
	if err != nil {
		zap.L().Error("Create Post in redis wrong", zap.Error(err))
		return err
	}
	return nil
}

func GetAllPosts(p *models.ParamPostList) (posts []*models.Post, err error) {
	posts, err = mysql.GetAllPostsByPageAndOrder(PageSize, p.Page, p.Order)
	if err != nil {
		return nil, err
	}
	return posts, nil
}

package logic

import (
	"bluebell/dao/mysql"
	"bluebell/models"
	"bluebell/pkg/snowflake"
	"go.uber.org/zap"
)

const (
	PageSize    int    = 10
	OrderOption string = "time"
	OrderRule          = "create_time"
)

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
	return nil
}

func GetAllPostsByPageAndOrder(page int, order string) (posts []models.Post, err error) {
	if order == OrderOption {
		order = OrderRule
	}
	posts, err = mysql.GetAllPostsByPageAndOrder(PageSize, page, order)
	if err != nil {
		return nil, err
	}
	return posts, nil
}

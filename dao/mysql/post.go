package mysql

import (
	"bluebell/models"
	"errors"
	"go.uber.org/zap"
)

/*--------------------------Post------------------------*/

func GetAllPostsByPageAndOrder(pageSize int, page int, order string) (posts []models.Post, err error) {
	sqlStr := "select post_id,title,content,author_id,community_id,status,create_time,update_time from post order by ? limit ?,? "
	err = db.Select(&posts, sqlStr, order, (page-1)*pageSize, pageSize)
	if err != nil {
		zap.L().Error("数据库查询出错", zap.Error(err))
		return nil, errors.New("数据库查询失败")
	}
	return posts, nil
}

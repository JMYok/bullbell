package mysql

import (
	"bluebell/models"
	"errors"
	"go.uber.org/zap"
)

/*--------------------------Post------------------------*/

func CreatePost(p *models.Post) (err error) {
	sqlStr := "insert into post (post_id,title,content,author_id,community_id) values (?,?,?,?,?)"
	_, err = db.Exec(sqlStr, p.Id, p.Title, p.Content, p.AuthorId, p.CommunityId)
	if err != nil {
		zap.L().Error("insert wrong", zap.Error(err))
		err = errors.New("插入数据出错")
		return err
	}
	return nil
}

func GetAllPostsByPageAndOrder(pageSize int, page int, order string) (posts []models.Post, err error) {
	sqlStr := "select post_id,title,content,author_id,community_id,status,create_time,update_time from post order by ? limit ?,? "
	err = db.Select(&posts, sqlStr, order, (page-1)*pageSize, pageSize)
	if err != nil {
		zap.L().Error("数据库查询出错", zap.Error(err))
		return nil, errors.New("数据库查询失败")
	}
	return posts, nil
}

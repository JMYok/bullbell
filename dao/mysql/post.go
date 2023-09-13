package mysql

import (
	"bluebell/models"
	"database/sql"
	"errors"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
	"strings"
)

/*--------------------------Post------------------------*/

func GetPostDetailById(pid uint64) (postDetail *models.Post, err error) {
	sqlSql := "select post_id,title,content,author_id,community_id,create_time,update_time from post where post_id = ?"
	postDetail = new(models.Post)
	err = db.Get(postDetail, sqlSql, pid)
	if err == sql.ErrNoRows {
		return nil, err
	} else if err != nil {
		zap.L().Error("Db failure", zap.Error(err))
		return nil, err
	}
	return postDetail, nil
}

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

func GetAllPostsByPageAndOrder(pageSize, page int64, order string) (posts []*models.Post, err error) {
	sqlStr := "select post_id,title,content,author_id,community_id,status,create_time,update_time from post order by ? DESC limit ?,? "
	posts = make([]*models.Post, 0, pageSize)
	err = db.Select(&posts, sqlStr, order, (page-1)*pageSize, pageSize)
	if err != nil {
		zap.L().Error("数据库查询出错", zap.Error(err))
		return nil, errors.New("数据库查询失败")
	}
	return posts, nil
}

// GetPostListByIDs 根据给定ids查询查询帖子数据
func GetPostListByIDs(ids []string) (postList []*models.Post, err error) {
	// FIND_IN_SET 提供排序序号
	sqlStr := `select post_id,title,content,author_id,community_id,create_time from post where post_id in(?)
	order by FIND_IN_SET(post_id,?)`
	query, args, err := sqlx.In(sqlStr, ids, strings.Join(ids, ","))

	if err != nil {
		zap.L().Error("sqlx.In failed", zap.Error(err))
		return nil, err
	}

	query = db.Rebind(query)
	err = db.Select(&postList, query, args...)

	if err != nil {
		zap.L().Error("db.Select failed", zap.Error(err))
		return nil, err
	}

	return postList, nil
}

package logic

import (
	"bluebell/dao/mysql"
	"bluebell/dao/redis"
	"bluebell/models"
	"bluebell/pkg/snowflake"
	"go.uber.org/zap"
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

	err = redis.CreatePost(postId, p.CommunityId)
	if err != nil {
		zap.L().Error("Create Post in redis wrong", zap.Error(err))
		return err
	}
	return nil
}

func GetAllPosts(p *models.ParamPostList) (postList []*models.ApiPostDetail, err error) {
	// 在redis中查询id列表
	ids, err := redis.GetPostIDsInOrder(p)
	if err != nil {
		zap.L().Error("redis.GetPostIDsInOrder 查询id失败", zap.Error(err))
		return
	}
	return getPostListByIds(ids)
}

func GetCommunityPostList(p *models.ParamCommunityPostList) (data []*models.ApiPostDetail, err error) {
	ids, err := redis.GetCommunityPostIDsInOrder(p)
	if err != nil {
		zap.L().Error("redis.GetCommunityPostIDsInOrder failed", zap.Error(err))
		return nil, err
	}
	return getPostListByIds(ids)
}

func getPostListByIds(ids []string) (postList []*models.ApiPostDetail, err error) {
	if len(ids) == 0 {
		zap.L().Warn("ids is empty")
		return
	}

	// 根据id去mysql数据库查询帖子详细信息
	posts, err := mysql.GetPostListByIDs(ids)
	if err != nil {
		zap.L().Error("mysql.GetPostListByIDs failed", zap.Error(err))
		return
	}

	//提前查询好每篇帖子的投票数
	voteData, err := redis.GetPostVoteData(ids)

	// 将帖子作者及分区信息查询出来填充到帖子中
	postList = make([]*models.ApiPostDetail, 0, len(posts))
	for idx, post := range posts {
		user, err := mysql.GetUserByUserId(&models.User{UserId: post.AuthorId})
		if err != nil {
			zap.L().Error("Get User By id failed", zap.Error(err))
			return nil, err
		}

		community, err := mysql.GetCommunityDetailByCid(post.CommunityId)
		if err != nil {
			zap.L().Error("Get community By id failed", zap.Error(err))
			return nil, err
		}
		postDetail := &models.ApiPostDetail{
			AuthorName:      user.Username,
			VoteNum:         voteData[idx],
			CommunityDetail: community,
			Post:            post,
		}
		postList = append(postList, postDetail)
	}
	return postList, nil
}

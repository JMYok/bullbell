package mysql

import (
	"bluebell/models"
	"database/sql"
	"go.uber.org/zap"
)

/*--------------------------Community------------------------*/

func GetCommunityList() ([]*models.Community, error) {
	sqlStr := "select community_id,community_name,introduction from community"
	var communities []*models.Community
	err := db.Select(&communities, sqlStr)
	if err == sql.ErrNoRows {
		zap.L().Warn("there is no community in db", zap.Error(err))
		return nil, err
	}
	return communities, nil
}

func GetCommunityDetailByCid(cid uint64) (community *models.CommunityDetail, err error) {
	sqlStr := "select community_id,community_name,introduction,create_time,update_time from community where community_id = ?"
	community = new(models.CommunityDetail)
	err = db.Get(community, sqlStr, cid)
	if err == sql.ErrNoRows {
		zap.L().Warn("there is no community in db", zap.Error(err))
		return nil, err
	} else if err != nil {
		zap.L().Error("there is error in db", zap.Error(err))
		return nil, err
	}
	return community, nil
}

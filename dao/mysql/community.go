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

package logic

import (
	"bluebell/dao/mysql"
	"bluebell/models"
)

func GetCommunityList() ([]*models.Community, error) {
	data, err := mysql.GetCommunityList()
	if err != nil {
		return nil, err
	}
	return data, nil
}

func GetCommunityDetailByCid(cid uint64) (community *models.CommunityDetail, err error) {
	community, err = mysql.GetCommunityDetailByCid(cid)
	if err != nil {
		return nil, err
	}
	return community, nil
}

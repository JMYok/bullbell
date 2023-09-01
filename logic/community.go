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

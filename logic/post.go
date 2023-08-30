package logic

import (
	"bluebell/dao/mysql"
	"bluebell/models"
)

const (
	PageSize    int    = 10
	OrderOption string = "time"
	OrderRule          = "create_time"
)

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

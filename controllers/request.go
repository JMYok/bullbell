package controllers

import (
	"bluebell/dao/mysql"
	"bluebell/models"
	"errors"
	"github.com/gin-gonic/gin"
)

const CtxUserIdKey = "userId"

var ErrorUserNotLogin = errors.New("用户未登录")

func GetCurrentUser(c *gin.Context) (user *models.User, err error) {
	userId, exists := c.Get(CtxUserIdKey)
	if !exists {
		err = ErrorUserNotLogin
		return
	}
	user = &models.User{
		UserId: userId.(uint64),
	}
	user, err = mysql.GetUserByUserId(user)
	if err != nil {
		return
	}
	return user, nil
}

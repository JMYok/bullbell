package logic

import (
	"bluebell/dao/mysql"
	"bluebell/models"
	"bluebell/pkg/jwt"
	"bluebell/pkg/snowflake"
	"errors"
)

func SignUp(p *models.ParamSignUp) (err error) {
	//查询用户是否存在
	var exist bool
	exist, err = mysql.CheckUserExistByUsername(p.Username)
	if err != nil {
		return err
	}

	if exist == true {
		return errors.New("用户已存在")
	}

	//生成UID
	uid, err := snowflake.GetID()
	if err != nil {
		return errors.New("生产UID失败")
	}

	//构造User对象
	user := &models.User{
		UserId:   uid,
		Username: p.Username,
		Password: p.Password,
	}

	//用户入库
	if err = mysql.InsertUser(user); err != nil {
		return err
	}

	//返回结果
	return nil
}

func Login(p *models.ParamLogin) (token string, err error) {
	user := &models.User{
		Username: p.Username,
		Password: p.Password,
	}
	//检测用户密码是否匹配
	if err = mysql.Login(user); err != nil {
		return "", err
	}

	return jwt.GenToken(user.UserId, user.Username)
}

package mysql

import (
	"bluebell/models"
	"crypto/md5"
	"encoding/hex"
	"errors"
	"fmt"
	"go.uber.org/zap"
)

/*--------------------------User------------------------*/

func CheckUserExistByUsername(username string) (bool, error) {
	sqlStr := "select count(*) from user where username=?"
	var cnt int
	err := db.Get(&cnt, sqlStr, username)
	if err != nil {
		return true, err
	}
	return cnt > 0, nil
}

// InsertUser 插入用户数据
func InsertUser(user *models.User) (err error) {
	//对密码进行加密
	password := encryptPassword(user.Password)
	//sql操作
	sqlStr := "insert into user(user_id,username,password) values(?,?,?)"
	_, err = db.Exec(sqlStr, user.UserId, user.Username, password)
	if err != nil {
		fmt.Printf("Error:%v", err)
		return errors.New("插入数据出错")
	}
	return nil
}

// GetUserByUserId 根据userId获得user
func GetUserByUserId(user *models.User) (err error) {
	sqlStr := "select user_id,username,email from user where user_id = ?"
	err = db.Get(&user, sqlStr, user.UserId)
	if err != nil {
		return errors.New("用户不存在")
	}
	return nil
}

func Login(user *models.User) (err error) {
	username := user.Username
	password := encryptPassword(user.Password)
	sqlStr := "select user_id from user where username=? and password = ?"
	var uid uint64
	err = db.Get(&uid, sqlStr, username, password)
	if err != nil {
		zap.L().Error("mysql Login error message", zap.Error(err))
		return err
	}

	user.UserId = uid
	if user.UserId == 0 {
		zap.L().Error("数据库查询结果为0", zap.String("username", username))
		return errors.New("密码不匹配")
	}
	return nil
}

func encryptPassword(oPassword string) string {
	h := md5.New()
	h.Write([]byte(secret))
	return hex.EncodeToString(h.Sum([]byte(oPassword)))
}

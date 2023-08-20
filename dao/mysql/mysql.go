package mysql

import (
	"bluebell/models"
	"bluebell/settings"
	"crypto/md5"
	"encoding/hex"
	"errors"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

const secret = "bobJiang"

var db *sqlx.DB

func Init() (err error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True",
		settings.Conf.MySQLConfig.User,
		settings.Conf.MySQLConfig.Password,
		settings.Conf.MySQLConfig.Host,
		settings.Conf.MySQLConfig.Port,
		settings.Conf.MySQLConfig.DbName,
	)
	// 也可以使用MustConnect连接不成功就panic
	db, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		zap.L().Error("connect DB failed", zap.Error(err))
		return err
	}
	db.SetMaxOpenConns(settings.Conf.MySQLConfig.MaxOpenConns)
	db.SetMaxIdleConns(settings.Conf.MySQLConfig.MaxIdleConns)
	return
}

func Close() {
	_ = db.Close()
}

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

func encryptPassword(oPassword string) string {
	h := md5.New()
	h.Write([]byte(secret))
	return hex.EncodeToString(h.Sum([]byte(oPassword)))
}

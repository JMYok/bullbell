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

func Init(cfg *settings.MySQLConfig) (err error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True",
		cfg.User,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.DbName,
	)
	// 也可以使用MustConnect连接不成功就panic
	db, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		zap.L().Error("connect DB failed", zap.Error(err))
		return err
	}
	db.SetMaxOpenConns(cfg.MaxOpenConns)
	db.SetMaxIdleConns(cfg.MaxIdleConns)
	return
}

func Close() {
	_ = db.Close()
}

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

/*--------------------------Post------------------------*/

func GetAllPostsByPageAndOrder(pageSize int, page int, order string) (posts []models.Post, err error) {
	sqlStr := "select post_id,title,content,author_id,community_id,status,create_time,update_time from post order by ? limit ?,? "
	err = db.Select(&posts, sqlStr, order, (page-1)*pageSize, pageSize)
	if err != nil {
		zap.L().Error("数据库查询出错", zap.Error(err))
		return nil, errors.New("数据库查询失败")
	}
	return posts, nil
}

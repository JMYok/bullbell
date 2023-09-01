package jwt

import (
	"bluebell/settings"
	"errors"
	"github.com/dgrijalva/jwt-go"
	"time"
)

const InvalidToken = "invalid token"

var MySecret = []byte("BobJiang的高盐值")

// MyClaims 自定义声明结构体并内嵌jwt.StandardClaims
// jwt包自带的jwt.StandardClaims只包含了官方字段
// 我们这里需要额外记录一个userid字段，所以要自定义结构体
// 如果想要保存更多信息，都可以添加到这个结构体中
type MyClaims struct {
	UserId   uint64 `json:"user_id"`
	UserName string `json:"user_name"`
	jwt.StandardClaims
}

// GenToken 生成JWT
func GenToken(userid uint64, username string) (accessToken, refreshToken string, err error) {
	// 创建一个我们自己的声明
	c := MyClaims{
		UserId:   userid, // 自定义字段
		UserName: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Duration(settings.Conf.AccessTokenExpire) * time.Minute).Unix(), // 过期时间
			Issuer:    "BobJiang",                                                                          // 签发人
		},
	}
	// 使用指定的签名方法创建签名对象
	// 使用指定的secret签名并获得完整的编码后的字符串token
	accessToken, err = jwt.NewWithClaims(jwt.SigningMethodHS256, c).SignedString(MySecret)
	refreshToken, err = jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		ExpiresAt: time.Now().Add(time.Duration(settings.Conf.RefreshTokenExpire) * time.Minute).Unix(), //过期时间
		Issuer:    "BobJiang",
	}).SignedString(MySecret)
	if err != nil {
		return "", "", errors.New(InvalidToken)
	}

	return accessToken, refreshToken, nil
}

// ParseToken 解析JWT
func ParseToken(tokenString string) (*MyClaims, error) {
	// 解析token:tokenString,解析目的变量,解析所需密钥
	token, err := jwt.ParseWithClaims(tokenString, &MyClaims{}, func(token *jwt.Token) (i interface{}, err error) {
		return MySecret, nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*MyClaims); ok && token.Valid { // 校验token
		return claims, nil
	}
	return nil, errors.New(InvalidToken)
}

func RefreshToken(accessToken, refreshToken string) (newAToken, newRToken string, err error) {
	//refresh token是否有效
	if _, err = jwt.Parse(refreshToken, func(token *jwt.Token) (i interface{}, err error) {
		return MySecret, nil
	}); err != nil {
		return
	}

	//从旧access token中解析出claims数据
	var claims MyClaims
	_, err = jwt.ParseWithClaims(accessToken, &claims, func(token *jwt.Token) (i interface{}, err error) {
		return MySecret, nil
	})
	v, _ := err.(*jwt.ValidationError)

	//access token过期
	if v.Errors == jwt.ValidationErrorExpired {
		return GenToken(claims.UserId, claims.UserName)
	}
	return
}

package routes

import (
	"bluebell/controllers"
	"bluebell/logger"
	"bluebell/middleware"
	"github.com/gin-gonic/gin"
)

func Setup(mode string) *gin.Engine {
	if mode == gin.ReleaseMode {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.New()
	r.Use(logger.GinLogger(), logger.GinRecovery(true))

	//注册
	r.POST("/signup", controllers.SignUpHandler)
	//登录
	r.POST("/login", controllers.LoginHandler)

	//刷新token
	r.POST("/refresh_token", controllers.RefreshTokenHandler)

	r.GET("/ping", middleware.JWTAuthMiddleware(), func(c *gin.Context) {
		c.String(200, "pong")
	})
	return r
}

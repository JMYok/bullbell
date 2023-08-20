package routes

import (
	"bluebell/controllers"
	"bluebell/logger"
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

	r.GET("/", func(c *gin.Context) {
		c.String(200, "server work nicely")
	})
	return r
}

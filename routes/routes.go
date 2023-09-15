package routes

import (
	"bluebell/controllers"
	_ "bluebell/docs"
	"bluebell/logger"
	"bluebell/middleware"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Setup(mode string) *gin.Engine {
	if mode == gin.ReleaseMode {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.New()
	r.Use(logger.GinLogger(), logger.GinRecovery(true))

	v1 := r.Group("/api/v1")

	//注册
	v1.POST("/signup", controllers.SignUpHandler)
	//登录
	v1.POST("/login", controllers.LoginHandler)

	v1.Use(middleware.JWTAuthMiddleware())
	{
		//-----------------------community------------------------

		v1.GET("/community", controllers.CommunityHandler)
		v1.GET("/community/:cid", controllers.CommunityDetailHandler)

		//-----------------------post------------------------------

		v1.POST("/post", controllers.CreatePostHandler)
		v1.GET("/post/:pid", controllers.PostDetailHandler)

		v1.POST("/vote", controllers.PostVoteHandler)

	}

	//所有博客
	v1.GET("/posts", controllers.AllPostsHandler)
	v1.GET("/community_posts", controllers.GetPostListByCommunityIDHandler)

	//刷新token
	r.POST("/refresh_token", controllers.RefreshTokenHandler)

	r.GET("/ping", middleware.JWTAuthMiddleware(), func(c *gin.Context) {
		c.String(200, "pong")
	})

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	return r
}

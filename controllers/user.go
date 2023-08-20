package controllers

import (
	"bluebell/logic"
	"bluebell/models"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
)

// SignUpHandler 处理注册请求的函数
func SignUpHandler(c *gin.Context) {
	//1. 参数处理
	p := new(models.ParamSignUp)

	//shouldBindJSON只能校验数据格式、数据类型
	if err := c.ShouldBindJSON(&p); err != nil {
		//判断是否为验证型错误
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			c.JSON(http.StatusOK, gin.H{
				"msg": "invalid param",
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"msg": removeTopStruct(errs.Translate(trans)),
			})
		}
		return
	}

	//2.业务处理
	if err := logic.SignUp(p); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"msg": "注册失败",
		})
		return
	}

	//3.返回响应
	c.JSON(http.StatusOK, gin.H{
		"msg": "request success",
	})
}

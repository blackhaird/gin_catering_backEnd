package controller

import "github.com/gin-gonic/gin"

func User_login(ctx *gin.Context) {

	ctx.JSON(200, gin.H{
		"code":       200,
		"user_level": 1,
		"msg":        "登陆成功",
	})
}

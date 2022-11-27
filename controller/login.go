package controller

import "github.com/gin-gonic/gin"

func User_login(ctx *gin.Context) {

	ctx.JSON(200, "login")
}

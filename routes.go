package main

import (
	"gin_catering_backEnd/controller"
	"github.com/gin-gonic/gin"
)

func CollectRoute(gin_server *gin.Engine) *gin.Engine {
	gin_server.GET("/login", controller.User_login)

	return gin_server
}

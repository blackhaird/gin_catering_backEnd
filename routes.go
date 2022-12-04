package main

import (
	"gin_catering_backEnd/controller"
	"github.com/gin-gonic/gin"
)

func CollectRoute(gin_server *gin.Engine) *gin.Engine {
	gin_server.POST("/user_login", controller.User_login)
	//gin_server.POST("/order_show", controller.Order_show())
	return gin_server
}

package main

import (
	"gin_catering_backEnd/controller"
	"github.com/gin-gonic/gin"
)

func CollectRoute(gin_server *gin.Engine) *gin.Engine {
	gin_server.POST("/user_login", controller.Admin_login)

	gin_server.GET("/order_show", controller.Order_show)

	gin_server.POST("/employee_chancePost", controller.Employee_chancePost)
	gin_server.POST("/employee_delete", controller.Employee_delete)
	gin_server.POST("/employee_add", controller.Employee_add)

	gin_server.GET("/dish_show", controller.Dish_show)
	gin_server.POST("/dish_chanceName", controller.Dish_chanceName)
	gin_server.POST("/dish_delete", controller.Dish_delete)
	gin_server.POST("/dish_add", controller.Dish_add)
	return gin_server
}

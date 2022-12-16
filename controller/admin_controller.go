package controller

import (
	"gin_catering_backEnd/model"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func Admin_login(ctx *gin.Context) {
	var jsoninfor model.PostAdmin

	//ctx.ShouldBind(&jsoninfor)
	//fmt.Println(jsoninfor.Admin[0].AdminName)

	if err := ctx.ShouldBind(&jsoninfor); err != nil {
		sqlStr := "select * from admin where Admin_name = ? and Admin_password = ? "
		row := db.QueryRow(sqlStr, jsoninfor.Admin[0].AdminName, jsoninfor.Admin[0].AdminPassword)
		var Admin model.Admin
		if err := row.Scan(&Admin.AdminID, &Admin.AdminName, &Admin.AdminPassword, &Admin.AdminLevel); err != nil {
			log.Println(err)

			ctx.JSON(http.StatusOK, gin.H{
				"code": 200,
				"msg":  "登录成功",
				"data": Admin,
			},
			)

			//ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			//	"code": 422,
			//	"data": []interface{}{gin.H{
			//		"msg": "账号或密码出错，找不到该用户",
			//	}},
			//})
		}
	}
}

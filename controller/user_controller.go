package controller

import (
	"gin_catering_backEnd/model"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func User_login(ctx *gin.Context) {
	var jsoninfor model.PostUser

	//ctx.ShouldBind(&jsoninfor)
	//fmt.Println(jsoninfor.User[0].UserName)

	if err := ctx.ShouldBind(&jsoninfor); err != nil {
		sqlStr := "select * from user where user_name = ? and user_password = ? "
		row := db.QueryRow(sqlStr, jsoninfor.User[0].UserName, jsoninfor.User[0].UserPassword)
		var user model.User
		if err := row.Scan(&user.UserID, &user.UserName, &user.UserPassword, &user.UserLevel); err != nil {
			log.Println(err)

			ctx.JSON(http.StatusOK, gin.H{
				"code": 200,
				"msg":  "登录成功",
				"data": user,
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

func User_register(ctx *gin.Context) {
	var jsoninfor model.PostUser
	if err := ctx.ShouldBind(&jsoninfor); err != nil {
		log.Print(&jsoninfor)
		//sqlstr := "insert into user(user_name,user_password,user_level) values (?,?,?,?)"
		//row := db.QueryRow(sqlstr, jsoninfor.User[0].UserName, jsoninfor.User[0].UserPassword)
	}
}

package controller

import (
	"gin_catering_backEnd/model"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func User_login(ctx *gin.Context) {
	jsoninfor := model.PostUser{}
	ctx.ShouldBindJSON(&jsoninfor)
	log.Println(&jsoninfor)

	//sqlStr := "select * from user where user_name = ? and user_password = ? "
	//row := db.QueryRow(sqlStr, jsoninfor.User.UserName, json.UserPassword)
	//var user model.User
	//if err := row.Scan(&user.UserID, &user.UserName, &user.UserPassword, &user.UserLevel); err != nil {
	//	log.Println(err)
	//	ctx.JSON(http.StatusUnprocessableEntity, gin.H{
	//		"code": 422,
	//		"data": []interface{}{gin.H{
	//			"msg": "账号或密码出错，找不到该用户",
	//		}},
	//	})
	//}

	ctx.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": jsoninfor.User,
	})
}

func User_() {

}

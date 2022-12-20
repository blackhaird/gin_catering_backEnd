package controller

import (
	"fmt"
	"gin_catering_backEnd/model"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func Admin_login(ctx *gin.Context) {
	fmt.Println("Admin_login is running")
	var jsoninfor model.PostAdmin
	if err := ctx.ShouldBind(&jsoninfor); err != nil {
		sqlStr := "select * from admin where userid = ? and password = ? "
		fmt.Println(jsoninfor.Admin[0].UserID)
		row := db.QueryRow(sqlStr, jsoninfor.Admin[0].UserID, jsoninfor.Admin[0].AdminPassword)
		if row == nil {
			fmt.Println("sql:sql:select is error")
			log.Fatalln("sql:sql:select is error")
			ctx.JSON(http.StatusUnprocessableEntity, gin.H{
				"code": 422,
				"msg":  "select is error",
			})
			return
		}
		var admin model.Admin
		row.Scan(&admin.AdminNo, &admin.UserID, &admin.AdminPassword)
		if admin.UserID != jsoninfor.Admin[0].UserID && admin.AdminPassword != jsoninfor.Admin[0].AdminPassword {
			ctx.JSON(http.StatusOK, gin.H{
				"code": 400,
				"msg":  "账号不存在或密码错误,请检查",
			})
		}
		ctx.JSON(http.StatusOK, gin.H{
			"code": 200,
			"data": []interface{}{
				gin.H{
					"admin_no":       admin.AdminNo,
					"user_id":        admin.UserID,
					"admin_password": admin.AdminPassword,
				},
			},
		})
	}
}

func User_add(ctx *gin.Context) {
	var jsoninfo model.PostAdmin
	if err := ctx.ShouldBind(&jsoninfo); err != nil {
		sql_select_now_data := "select * from admin where userid = ? ;"
		row := db.QueryRow(sql_select_now_data, jsoninfo.Admin[0].UserID)
		var now_admin model.Admin
		row.Scan(&now_admin.AdminNo,
			&now_admin.UserID,
			&now_admin.AdminPassword,
		)
		if now_admin.UserID != jsoninfo.Admin[0].UserID {
			if jsoninfo.Admin[0].AdminPassword != 0 && len(string(jsoninfo.Admin[0].AdminPassword)) > 2 {
				sqlStr := "insert into admin( userid, password) VALUES (?,?)"
				_, err := db.Exec(sqlStr, jsoninfo.Admin[0].UserID, jsoninfo.Admin[0].AdminPassword)
				if err != nil {
					log.Fatalln(err)
					ctx.JSON(http.StatusUnprocessableEntity, gin.H{
						"code": 422,
						"msg":  "insert sql is error",
					})
					return
				} else {
					sqlStr = "select * from admin where userid = ? and password = ?;"
					row := db.QueryRow(sqlStr, jsoninfo.Admin[0].UserID, jsoninfo.Admin[0].AdminPassword)
					var admin model.Admin
					row.Scan(&admin.AdminNo,
						&admin.UserID,
						&admin.AdminPassword)
					ctx.JSON(http.StatusOK, gin.H{
						"code": 200,
						"msg":  "注册成功",
						"data": admin,
					})
				}
			} else {
				ctx.JSON(http.StatusOK, gin.H{
					"code": 400,
					"msg":  "密码太短或密码为空（大于4位数）",
				})
			}
		} else {
			fmt.Println("已有同名数据")
			ctx.JSON(http.StatusOK, gin.H{
				"code": 400,
				"msg":  "已有同名id，请修改id注册",
			})
		}
	}
}

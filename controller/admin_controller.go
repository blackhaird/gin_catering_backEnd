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
		if err = row.Err(); err != nil {
			fmt.Println("sql:sql:select is error")
			log.Fatalln(err)
		}
		ctx.JSON(http.StatusOK, gin.H{
			"code": 200,
			"data": []interface{}{
				gin.H{
					"admin_no":       admin.AdminNo,
					"uesr_id":        admin.UserID,
					"admin_password": admin.AdminPassword,
				},
			},
		})
	}
}

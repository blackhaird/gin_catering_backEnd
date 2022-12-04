package controller

import (
	"fmt"
	"gin_catering_backEnd/model"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func Order_show(ctx *gin.Context) {
	fmt.Println("running:User_login")
	rows, err := db.Query("select * from order")
	if err != nil {
		fmt.Println("sql:select is error")
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"code": 422,
			"msg":  "select",
		})
		return
	}
	orders := make([]model.Order, 0)
	for rows.Next() {
		var order model.Order
		//rows.Scan(&order.UserID, &order.UserName, &order.UserPassword, &order.UserLevel)
		orders = append(orders, order)
	}
	if err = rows.Err(); err != nil {
		log.Fatalln(err)
	}
	ctx.JSON(
		200,
		gin.H{
			"code": 200,
			"data": orders,
		},
	)

}

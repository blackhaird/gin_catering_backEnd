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
	rows, err := db.Query("SELECT order_no,customer_name,address,`order`.`time`,age,gender FROM customer JOIN customerorder ON customer.`no` = customerorder.customer_no JOIN `order` ON customerorder.order_no = `order`.`no`;")
	if err != nil {
		fmt.Println("sql:select is error")
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"code": 422,
			"msg":  "select",
		})
		return
	}
	Orders := make([]model.Order, 0)
	for rows.Next() {
		var Order model.Order
		rows.Scan(&Order.OrderNo, &Order.OrderCustomerName, &Order.OrderAddress, &Order.OrderTime, &Order.OrderAge, &Order.OrderGender)
		Orders = append(Orders, Order)
	}
	if err = rows.Err(); err != nil {
		log.Fatalln(err)
	}
	ctx.JSON(
		200,
		gin.H{
			"code": 200,
			"data": Orders,
		},
	)

}

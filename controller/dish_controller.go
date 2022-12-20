package controller

import (
	"fmt"
	"gin_catering_backEnd/model"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func Dish_show(ctx *gin.Context) {
	fmt.Println("running:dish_show is running")
	rows, err := db.Query("SELECT * FROM dish")
	if err != nil {
		fmt.Println("sql:dish_show is error")
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"code": 422,
			"msg":  "dish_show is error",
		})
		return
	}

	dishs := make([]model.Dish, 0)

	for rows.Next() {
		var dish model.Dish
		rows.Scan(&dish.DishNo, &dish.DishName, &dish.DishPrice, &dish.DishTaste)
		dishs = append(dishs, dish)
	}
	if err = rows.Err(); err != nil {
		log.Fatalln(err)
	}
	ctx.JSON(
		200,
		gin.H{
			"code": 200,
			"data": dishs,
		},
	)
}

func Dish_chanceName(ctx *gin.Context) {
	//sqlstr = "UPDATE dish SET dish_name = ? WHERE dish.`no`=  ?;"
	var jsoninfo model.PostDish
	if err := ctx.ShouldBind(&jsoninfo); err != nil {
		sql_select_now_data := "select * from dish where dish_name = ?"
		row := db.QueryRow(sql_select_now_data, jsoninfo.Dish[0].DishName)
		var now_dish model.Dish
		row.Scan(&now_dish.DishNo, &now_dish.DishName, &now_dish.DishPrice, &now_dish.DishTaste)
		fmt.Println(now_dish.DishNo)
		if now_dish.DishName != jsoninfo.Dish[0].DishName {
			sqlStr := "UPDATE dish SET dish_name = ? WHERE dish.`no`=  ?;"
			stmt, err := db.Exec(sqlStr, jsoninfo.Dish[0].DishName, jsoninfo.Dish[0].DishNo)
			if err != nil {
				log.Println(err)
				ctx.JSON(http.StatusUnprocessableEntity, gin.H{
					"code": 422,
					"msg":  "Update sql is error",
				})
			} else {
				fmt.Println(stmt)
				ctx.JSON(http.StatusOK, gin.H{
					"code": 200,
					"msg":  "修改成功",
					"data": []interface{}{gin.H{
						"dish_name": jsoninfo.Dish[0].DishName,
						"dish_no":   jsoninfo.Dish[0].DishNo,
					}},
				})
			}
		} else {
			ctx.JSON(http.StatusOK, gin.H{
				"code": 400,
				"msg":  "已经存在同名数据 [" + now_dish.DishName + "] 无需加入这个菜系",
			})
		}
	}
}

func Dish_delete(ctx *gin.Context) {
	var jsoninfo model.PostDish
	if err := ctx.ShouldBind(&jsoninfo); err != nil {
		_, err := db.Query("DELETE FROM dish WHERE dish.`no` = ?;", jsoninfo.Dish[0].DishNo)
		if err != nil {
			log.Fatal(err)
			ctx.JSON(http.StatusUnprocessableEntity, gin.H{
				"code": 422,
				"msg":  "delete sql is error",
			})
		} else {
			ctx.JSON(http.StatusOK, gin.H{
				"code": 200,
				"msg":  "删除成功",
			})
		}
	}
}

func Dish_add(ctx *gin.Context) {
	var jsoninfo model.PostDish
	if err := ctx.ShouldBind(&jsoninfo); err != nil {
		sql_select_now_data := "select * from dish where dish_name = ?"
		row := db.QueryRow(sql_select_now_data, jsoninfo.Dish[0].DishName)
		var now_dish model.Dish
		row.Scan(&now_dish.DishNo, &now_dish.DishName, &now_dish.DishPrice, &now_dish.DishTaste)
		fmt.Println(now_dish.DishNo)
		if now_dish.DishName != jsoninfo.Dish[0].DishName {
			_, err := db.Exec("INSERT INTO dish(dish_name,dish_price,taste)  VALUES (?,?,?)",
				jsoninfo.Dish[0].DishName,
				jsoninfo.Dish[0].DishPrice,
				jsoninfo.Dish[0].DishTaste)
			if err != nil {
				log.Fatalln(err)
				ctx.JSON(http.StatusUnprocessableEntity, gin.H{
					"code": 422,
					"msg":  "insert sql is error",
				})
			} else {
				ctx.JSON(http.StatusOK, gin.H{
					"code": 422,
					"msg":  "新增成功",
					"data": []interface{}{
						gin.H{
							"dish_name":  jsoninfo.Dish[0].DishName,
							"dish_price": jsoninfo.Dish[0].DishPrice,
							"dish_taste": jsoninfo.Dish[0].DishTaste,
						},
					},
				})
			}
		} else {
			ctx.JSON(http.StatusOK, gin.H{
				"code": 400,
				"msg":  "已经存在同名数据 [" + now_dish.DishName + "] 无需加入这个菜系",
			})
		}
	}
}

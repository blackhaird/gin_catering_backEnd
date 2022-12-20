package controller

import (
	"fmt"
	"gin_catering_backEnd/model"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func Employee_show(ctx *gin.Context) {
	rows, err := db.Query("select * from employee;")
	if err != nil {
		fmt.Println("sql: select * from employee is fail")
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"code": 422,
			"msg":  "sql语句执行出现问题，返回数据为空",
		})
		return
	}
	employees := make([]model.Employee, 0)
	for rows.Next() {
		var employee model.Employee
		rows.Scan(&employee.EmployeeNo,
			&employee.EmployeeName,
			&employee.EmployeeAge,
			&employee.EmployeeGender,
			&employee.EmployeeEnterTime,
			&employee.EmployeeAddress,
			&employee.EmployeePost,
			&employee.EmployeeSalary,
			&employee.EmployeePower)
		employees = append(employees, employee)
	}
	if err = rows.Err(); err != nil {
		log.Fatalln(err)
	}
	ctx.JSON(200, gin.H{
		"code": 200,
		"data": employees,
	})
}

func Employee_chancePost(ctx *gin.Context) {
	var jsoninfo model.PostEmployee
	if err := ctx.ShouldBind(&jsoninfo); err != nil {
		sqlStr := "UPDATE Employee SET Employee.post = ?  WHERE `no` = ?;"
		stmt, err := db.Exec(sqlStr, jsoninfo.Employee[0].EmployeePost, jsoninfo.Employee[0].EmployeeNo)
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
					"dish_name": jsoninfo.Employee[0].EmployeePost,
					"dish_no":   jsoninfo.Employee[0].EmployeeNo,
				}},
			})
		}
	}
}

func Employee_delete(ctx *gin.Context) {
	var jsoninfo model.PostEmployee
	if err := ctx.ShouldBind(&jsoninfo); err != nil {

	}
	_, err := db.Query("DELETE FROM Employee WHERE Employee.`no` = ?;", jsoninfo.Employee[0].EmployeeNo)
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

func Employee_add(ctx *gin.Context) {
	var jsoninfo model.PostEmployee
	if err := ctx.ShouldBind(&jsoninfo); err != nil {
		_, err := db.Exec("INSERT INTO Employee(`name`,age,gender,enterTime,address,post,salary,power)  VALUES (?,?,?,?,?,?,?,?)",
			jsoninfo.Employee[0].EmployeeName,
			jsoninfo.Employee[0].EmployeeAge,
			jsoninfo.Employee[0].EmployeeGender,
			jsoninfo.Employee[0].EmployeeEnterTime,
			jsoninfo.Employee[0].EmployeeAddress,
			jsoninfo.Employee[0].EmployeePost,
			jsoninfo.Employee[0].EmployeeSalary,
			jsoninfo.Employee[0].EmployeePower)
		if err != nil {
			log.Fatalln(err)
			ctx.JSON(http.StatusUnprocessableEntity, gin.H{
				"code": 422,
				"msg":  "insert sql is error",
			})
			return
		} else {
			sqlStr := "select * from employee where name = ? and age = ? and gender = ? and enterTime = ?;"
			row := db.QueryRow(sqlStr,
				jsoninfo.Employee[0].EmployeeName,
				jsoninfo.Employee[0].EmployeeAge,
				jsoninfo.Employee[0].EmployeeGender,
				jsoninfo.Employee[0].EmployeeEnterTime)
			var employee model.Employee
			row.Scan(&employee.EmployeeNo,
				&employee.EmployeeName,
				&employee.EmployeeAge,
				&employee.EmployeeGender,
				&employee.EmployeeEnterTime,
				&employee.EmployeeAddress,
				&employee.EmployeePost,
				&employee.EmployeeSalary,
				&employee.EmployeePower)
			ctx.JSON(http.StatusOK, gin.H{
				"code": 422,
				"msg":  "新增成功",
				"data": employee,
			})
		}
	}
}

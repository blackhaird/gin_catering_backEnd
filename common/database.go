package common

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

var db, err = sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/catering?charset=utf8")

func InitDB() {
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("database connect success")
	}
	return
}

func GetDB() *sql.DB {
	return db
}

func GetError() error {
	return err
}

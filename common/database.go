package common

import (
	"database/sql"
	"fmt"
)

var db, err = sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/database_name?charset=utf8")

func initDB() {
	if err != nil {
		fmt.Println(err)
	}
	return
}

func getDB() *sql.DB {
	return db
}

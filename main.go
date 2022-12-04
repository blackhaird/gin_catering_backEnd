package main

import (
	"database/sql"
	"gin_catering_backEnd/common"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

var db = common.InitDB()

func main() {
	gin_server := gin.Default()
	gin_server = CollectRoute(gin_server)
	err := gin_server.Run(":8401")
	if err != nil {
		log.Fatalln(err)
	}

}

func get_db() *sql.DB {
	return db
}

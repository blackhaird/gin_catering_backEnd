package main

import (
	"gin_catering_backEnd/common"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func main() {
	common.InitDB()
	gin_server := gin.Default()
	gin_server = CollectRoute(gin_server)
	err := gin_server.Run(":8401")
	if err != nil {
		log.Fatalln(err)
	}

}

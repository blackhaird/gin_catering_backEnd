package main

import (
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	gin_server := gin.Default()

	gin_server = CollectRoute(gin_server)
	err := gin_server.Run(":8401")
	if err != nil {
		log.Fatalln(err)
	}

}

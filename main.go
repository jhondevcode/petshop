package main

import (
	"fmt"
	"log"
	"net/http"

	"petshop/config"
	"petshop/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	config.LoadEnv()
	db := config.ConnectDB()
	engine := gin.Default()
	engine.Use(config.ConfigCors())

	engine.GET("", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Hello World"})
	})

	routes.AuthRouter(engine, db)

	log.Panic(engine.Run(fmt.Sprintf(":%d", config.ServerPort())))
}

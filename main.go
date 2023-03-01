package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jhondevcode/petshop/config"
	"github.com/jhondevcode/petshop/routes"
)

func main() {
	config.LoadEnv(true)
	db := config.ConnectDB()
	engine := gin.Default()
	engine.Use(config.ConfigCors())

	engine.GET("", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Hello World"})
	})

	routes.AuthRouter(engine, db)

	log.Panic(engine.Run(":8080"))
}

package main

import (
	"drink-api/config"
	"drink-api/database"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	config.Load()
	db := database.New()
	db.Connect()
	defer db.Connection.Close()

	v1 := router.Group("/api/v1")

	router.Run()
}

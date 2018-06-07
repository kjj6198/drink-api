package main

import (
	"drink-api/apis"
	"drink-api/config"
	"drink-api/database"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func main() {
	router := gin.Default()
	config.Load()

	db := gorm.Open("postgres", {
		User: "kalan"
	})

	v1 := router.Group("/api/v1")

	v1.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "hello",
		})
	})

	v1.POST("/menus", apis.Create)
	if os.Getenv("env") == "development" {
		v1.Use(func(c *gin.Context) {
			c.Writer.Header().Set(
				"Access-Control-Allow-Headers",
				"Content-Type",
			)

			c.Writer.Header().Set(
				"Access-Control-Allow-Methods",
				"OPTIONS,GET,POST,DELETE,PUT,PATCH",
			)
		})
	}

	router.Run()
}

package main

import (
	"drink-api/apis"
	"drink-api/config"
	"drink-api/database"
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func main() {

	router := gin.Default()
	config.Load()

	db, err := database.Connect(&database.DataBaseOptions{
		Addr:     os.Getenv("HOST"),
		Port:     os.Getenv("DATABASE_PORT"),
		User:     os.Getenv("DATABASE_USERNAME"),
		Password: os.Getenv("DATABASE_PASSWORD"),
		Database: os.Getenv("DATABASE"),
		SSLMode:  false,
	})

	db.LogMode(true)

	if err != nil {
		fmt.Println("can not connect database.")
		panic(err)
	}

	v1 := router.Group("/api/v1")
	v1.Use(func(c *gin.Context) {
		c.Set("db", db)
	})

	v1.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "hello",
		})
	})

	v1.POST("/menus", apis.Create)
	v1.GET("/menus/:id", apis.Show)
	v1.POST("/upload", apis.GetPresignURL)
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

package main

import (
	"drink-api/apis"
	"drink-api/config"
	"fmt"
	"net/http"
	"os"

	"github.com/jinzhu/gorm"

	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func main() {
	env := os.Getenv("env")
	router := gin.Default()
	config.Load()
	dbstr := fmt.Sprintf(
		"host=%s port=%s user=%s dbname=%s",
		os.Getenv("DATABASE_PORT"),
		os.Getenv("HOST"),
		os.Getenv("DATABASE_USERNAME"),
		os.Getenv("DATABASE"),
	)

	db, err := gorm.Open("postgres", dbstr)

	if env == "development" {
		db.LogMode(true)
	}

	if err != nil {
		fmt.Println("can not connect database.")
		panic(err)
	}

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

package apis

import (
	"drink-api/models"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/go-pg/pg"
)

func Create(c *gin.Context) {
	db := c.MustGet("db").(*pg.DB)
	var menu models.Menu
	c.ShouldBindJSON(&menu)
	result, err := db.Model(&menu).Returning("*").Insert()
	if err != nil {
		fmt.Println(err)
	}
	c.JSON(200, result)
}

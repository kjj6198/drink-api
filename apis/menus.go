package apis

import (
	"drink-api/models"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type menuParams struct {
	Name String
}

func Create(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "ok",
	})
}

func Show(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	result := db.First(&models.Menu{})
	menu := result.Value.(*models.Menu)
	sec, isEnded := menu.GetRemainTime()
	if isEnded {
		c.JSON(200)
	}
	c.JSON(200, result)
}

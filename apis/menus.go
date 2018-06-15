package apis

import (
	"drink-api/models"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func Create(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "ok",
	})
}

func Show(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	id := c.Param("id")
	menu := db.
		Model(&models.Menu{}).
		Where("id = ?", id)

	if menu.RecordNotFound() {
		c.JSON(404, gin.H{
			"message": "error",
		})
		return
	}

	c.JSON(200, menu.Value)
}

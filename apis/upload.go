package apis

import (
	"drink-api/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Params struct {
	Filename string `json:"filename"`
}

func GetPresignURL(c *gin.Context) {
	var params Params
	c.BindJSON(&params)

	url, err := services.GetPresignURL(params.Filename)
	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]string{
			"message": err.Error(),
		})
	}

	if params.Filename != "" {
		c.JSON(http.StatusOK, map[string]string{
			"url": url,
		})
	}
}

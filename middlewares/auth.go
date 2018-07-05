package middlewares

import (
	"drink-api/services"
	"os"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

var (
	signingKey = []byte(os.Getenv("JWT_SECRET_KEY"))
)

type userInfoClaims struct {
	UserInfo services.UserInfo
	jwt.StandardClaims
}

func Auth(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	jwtToken := c.Request.Cookie("")
	token, err := jwt.ParseWithClaims()
}

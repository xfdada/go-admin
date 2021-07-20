package middleware

import (
	"go-admin/utils/jwts"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func Jwt() gin.HandlerFunc {
	return func(c *gin.Context) {
		var (
			token   string
			code    = 200
			message string
		)
		if s, exist := c.GetQuery("token"); exist {
			token = s
		} else {
			token = c.GetHeader("token")
		}
		if token == "" {
			code = 500
			message = "token为空"
		} else {
			_, err := jwts.ParseToken(token)
			if err != nil {
				switch err.(*jwt.ValidationError).Errors {
				case jwt.ValidationErrorExpired:
					code = 500
					message = "token过期"
				default:
					code = 500
					message = "token无效"
				}
			}
		}
		if code != 200 {
			c.JSON(200, gin.H{"code": code, "err": message})
			c.Abort()
			return
		}
		c.Next()
	}
}

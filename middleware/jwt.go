package middleware

import (
	"go-admin/global/errcode"
	"go-admin/global/response"
	"go-admin/utils/jwts"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func Jwt() gin.HandlerFunc {
	return func(c *gin.Context) {
		r := response.NewResponse(c)
		var (
			token string
			eCode = errcode.Success
		)
		if s, exist := c.GetQuery("token"); exist {
			token = s
		} else {
			token = c.GetHeader("token")
		}
		if token == "" {
			eCode = errcode.NoToken
		} else {
			_, err := jwts.ParseToken(token)
			if err != nil {
				switch err.(*jwt.ValidationError).Errors {
				case jwt.ValidationErrorExpired:
					eCode = errcode.TokenTimeout
				default:
					eCode = errcode.TokenError
				}
			}
		}
		if eCode != errcode.Success {
			r.ToError(eCode)
			c.Abort()
			return
		}
		c.Next()
	}
}

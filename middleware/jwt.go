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
			token  string
			e_code = errcode.Success
		)
		if s, exist := c.GetQuery("token"); exist {
			token = s
		} else {
			token = c.GetHeader("token")
		}
		if token == "" {
			e_code = errcode.NoToken
		} else {
			_, err := jwts.ParseToken(token)
			if err != nil {
				switch err.(*jwt.ValidationError).Errors {
				case jwt.ValidationErrorExpired:
					e_code = errcode.TokenTimeout
				default:
					e_code = errcode.TokenError
				}
			}
		}
		if e_code != errcode.Success {
			r.ToError(e_code)
			c.Abort()
			return
		}
		c.Next()
	}
}

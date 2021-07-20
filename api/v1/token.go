package v1

import (
	"fmt"
	"go-admin/utils/jwts"

	"github.com/gin-gonic/gin"
)

func GetToken(c *gin.Context) {
	token, err := jwts.GenerateToken("xfdada", "go-admin")
	if err != nil {
		c.JSON(200, gin.H{"msg": "failed", "err": fmt.Sprintf("GenerateToken err:%v", err)})
		return
	}
	c.JSON(200, gin.H{"msg": "success", "token": token})
}

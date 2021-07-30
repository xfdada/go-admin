package v1

import (
	"fmt"
	"go-admin/utils/captcha"
	"go-admin/utils/jwts"
	"go-admin/utils/loggers"
	"go-admin/utils/redis"

	"github.com/gin-gonic/gin"
)

//@Tags 固定接口
//@Summary 获取token
//@Produce json
//@Success 200   "成功"
//@Failure 400 {object} errcode.Error "未找到"
//@Failure 500 {object} errcode.Error "内部错误"
//@Router /api/get_token [get]
func GetToken(c *gin.Context) {
	token, err := jwts.GenerateToken("xfdada", "go-admin") //这里可以自定义从数据库中取用户的账号和密码生成
	if err != nil {
		c.JSON(200, gin.H{"msg": "failed", "err": fmt.Sprintf("GenerateToken err:%v", err)})
		return
	}
	c.JSON(200, gin.H{"msg": "success", "token": token})
}

func GetCapt(c *gin.Context) {
	id, capt := captcha.GetCaptcha()
	if id != "" {
		c.JSON(200, gin.H{"id": id, "capt": capt})
	}

}

func SetKey(c *gin.Context) {
	key := c.PostForm("key")
	value := c.PostForm("value")
	loggers.Logs(" Key: " + key + " value: " + value)
	err := redis.Set(key, value, 0)
	if err != nil {

		c.JSON(500, gin.H{"msg": "failed"})
		return
	}
	c.JSON(200, gin.H{"msg": "success"})
	return
}

func GetKey(c *gin.Context) {
	key := c.PostForm("key")
	ok, err := redis.Get(key)

	if err != nil {
		loggers.Logs("get Key failed")
		c.JSON(500, gin.H{"msg": "failed"})
		return
	}
	c.JSON(200, gin.H{"msg": "success", "data": ok})
	return
}

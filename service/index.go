package service

import (
	"go-admin/model"
	"go-admin/utils/captcha"
	"go-admin/utils/redis"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {

	c.HTML(http.StatusOK, "index/index.tmpl", gin.H{"title": "欢迎登录系统"})
}

func Login(c *gin.Context) {
	types := c.PostForm("type")
	userName := c.PostForm("username")
	if types == "pwd" {
		passWord := c.PostForm("password")
		hid := c.PostForm("hid")
		capt := c.PostForm("captcha")
		ok := captcha.Verify(hid, capt)
		if !ok {
			c.HTML(http.StatusOK, "index/index.tmpl", gin.H{"title": "欢迎登录系统"})
			return
		}
		video := model.NewVideoUser()
		ok = video.Login(userName, passWord)
		if !ok {
			c.HTML(http.StatusOK, "index/index.tmpl", gin.H{"title": "欢迎登录系统"})
			return
		}

	} else {
		v, _ := redis.Get(userName)
		capt := c.PostForm("captcha")
		if v != capt {
			c.HTML(http.StatusOK, "index/index.tmpl", gin.H{"title": "欢迎登录系统"})
			return
		}
	}
	c.JSON(http.StatusOK, gin.H{"mesg": "验证成功"})
}

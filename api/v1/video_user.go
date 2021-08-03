package v1

import (
	"github.com/gin-gonic/gin"
	"go-admin/global/errcode"
	"go-admin/global/response"
	"go-admin/model"
	"net/http"
	"strconv"
)

func VideoGet(c *gin.Context) {
	res := response.NewResponse(c)
	v := model.NewVideoUser()
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := v.Get(id)
	if err != nil {
		res.ToError(errcode.AddError)
		c.Abort()
		return
	}
	res.ToResponse(data)
	c.Abort()
	return
}

func VideoCreate(c *gin.Context) {
	info := make(map[string]interface{}, 10)
	info["nickName"] = c.PostForm("nickName")
	info["headIcon"] = c.PostForm("headIcon")
	info["phone"] = c.PostForm("phone")
	info["email"] = c.PostForm("email")
	info["password"] = c.PostForm("password")
	info["desc"] = c.PostForm("desc")
	res := response.NewResponse(c)
	v := model.NewVideoUser()
	err := v.Create(info)
	if err != nil {
		res.ToError(errcode.AddError)
		c.Abort()
		return
	}
	res.ToError(errcode.Success)
	c.Abort()
	return
}

func CheckUser(c *gin.Context) {
	uuids := c.Param("uuid")
	v := model.NewVideoUser()
	err := v.OnLive(uuids)
	if err != nil {
		c.HTML(http.StatusOK, "index/check.tmpl", gin.H{"msg": "激活失败"})
		c.Abort()
		return
	}
	c.HTML(http.StatusOK, "index/check.tmpl", gin.H{"msg": "激活成功"})
}

func VideoUserList(c *gin.Context) {
	res := response.NewResponse(c)
	v := model.NewVideoUser()
	data, err := v.List(c)
	if err != nil {
		res.ToError(errcode.NotFoundError)
		c.Abort()
		return
	}
	res.ToResponse(data)
	c.Abort()
	return
}

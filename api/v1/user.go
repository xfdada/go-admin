package v1

import (
	"fmt"
	"go-admin/global/errcode"
	"go-admin/global/response"
	"go-admin/model"
	"go-admin/utils/upload"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type User struct{}

func NewUser() User {
	return User{}
}

//@Summary 获取指定用户信息
//@Produce json
//@Param id path int true "用户ID"
//@Success 200 {object} model.User "成功"
//@Failure 400 {object} errcode.Error "未找到"
//@Failure 500 {object} errcode.Error "内部错误"
//@Router /api/v1/user/{id} [get]
func (u User) Get(c *gin.Context) {
	r := response.NewResponse(c)
	id := c.Param("id")
	ids, _ := strconv.Atoi(id)
	if ids <= 0 {
		r.ToError(errcode.ParamsError)
		c.Abort()
		return
	}
	user := model.NewUser()
	data, row := user.Get(id)
	if row == 0 {
		r.ToError(errcode.NotFoundError)
		c.Abort()
		return
	}
	r.ToResponse(data)
	c.Abort()
	return
}

//@Summary 获取多条用户信息
//@Produce json
//@Success 200  "成功"
//@Failure 400 "请求错误"
//@Failure 500 "内部错误"
//@Router /api/v1/user [get]
func (u User) List(c *gin.Context) {

	user := model.NewUser()
	data, err := user.List()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"msg": "success", "data": data})
		return
	}
	c.JSON(http.StatusOK, gin.H{"msg": "success", "data": data})
}

//@Summary 新增用户信息
//@Produce json
//@Param name body string true "用户名"
//@Param pwd body string true "密码"
//@Param email body string true "邮箱"
//@Success 200  "成功"
//@Failure 400 "请求错误"
//@Failure 500 "内部错误"
//@Router /api/v1/user [post]
func (u User) Create(c *gin.Context) {
	name := c.PostForm("name")
	age := c.PostForm("age")
	address := c.PostForm("address")
	user := model.NewUser()
	err := user.Add(name, address, age)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"msg": "failed"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"msg": "success"})
}

//@Summary 更新指定用户信息
//@Produce json
//@Param id path int true "用户ID"
//@Param name body string true "用户名"
//@Param age body int true "年龄"
//@Param address body string true "地址"
//@Success 200  "成功"
//@Failure 400 "请求错误"
//@Failure 500 "内部错误"
//@Router /api/v1/user/{id} [put]
func (u User) Update(c *gin.Context) {
	id := c.Param("id")
	name := c.PostForm("name")
	age := c.PostForm("age")
	address := c.PostForm("address")
	user := model.NewUser()
	err := user.Update(id, name, address, age)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"msg": "failed"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"msg": "success"})
}

//@Summary 删除指定用户信息
//@Produce json
//@Param id path int true "用户ID"
//@Success 200  "成功"
//@Failure 400 "请求错误"
//@Failure 500 "内部错误"
//@Router /api/v1/user/{id} [delete]
func (u User) Delete(c *gin.Context) {

	id := c.Param("id")
	user := model.NewUser()
	row, err := user.Delete(id)
	if err != nil || row == 0 {
		c.JSON(http.StatusOK, gin.H{"msg": "failed"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"msg": "success"})
}

func (u User) Upload(c *gin.Context) {
	file, fileHeader, err := c.Request.FormFile("file")
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"msg": "文件上传失败1"})
		return
	}
	url, err := upload.UploadFile(file, fileHeader)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"msg": "文件上传失败2" + fmt.Sprintf("err:%v", err)})
		return
	}
	c.JSON(http.StatusOK, gin.H{"msg": "文件上传成功", "url": url})
}

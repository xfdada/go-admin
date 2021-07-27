package v1

import (
	"fmt"
	"go-admin/model"
	_ "go-admin/model"
	"go-admin/utils/upload"
	"net/http"

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
//@Failure 500 {object} User "内部错误"
//@Router /api/v1/user/{id} [get]
func (u User) Get(c *gin.Context) {
	id := c.Param("id")
	user := model.NewUser()
	data,err := user.Get(id)
	if err != nil{
		c.JSON(http.StatusOK, gin.H{"msg": "success","data":data})
		return
	}
	c.JSON(http.StatusOK, gin.H{"msg": "success","data":data})
}

//@Summary 获取多条用户信息
//@Produce json
//@Success 200  "成功"
//@Failure 400 "请求错误"
//@Failure 500 "内部错误"
//@Router /api/v1/user [get]
func (u User) List(c *gin.Context) {

	user := model.NewUser()
	data,err := user.List()
	if err != nil{
		c.JSON(http.StatusOK, gin.H{"msg": "success","data":data})
		return
	}
	c.JSON(http.StatusOK, gin.H{"msg": "success","data":data})
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
	adderss := c.PostForm("address")
	user := model.NewUser()
	err := user.Add(name,adderss,age)
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
//@Param pwd body string true "密码"
//@Param email body string true "邮箱"
//@Success 200  "成功"
//@Failure 400 "请求错误"
//@Failure 500 "内部错误"
//@Router /api/v1/user/{id} [put]
func (u User) Update(c *gin.Context) {

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

package v1

import (
	"github.com/gin-gonic/gin"
	"go-admin/global/errcode"
	"go-admin/global/response"
	"go-admin/model"
	"strconv"
)

func Article_Get(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	article := model.NewArticle()
	res := response.NewResponse(c)
	data, err := article.Get(id)
	if err != nil {
		res.ToError(errcode.NotFoundError)
		c.Abort()
		return
	}
	res.ToResponse(data)
	c.Abort()
	return

}

func Article_List(c *gin.Context) {

}

func Article_Create(c *gin.Context) {
	title := c.PostForm("title")
	desc := c.PostForm("desc")
	url := c.PostForm("url")
	content := c.PostForm("content")
	createdBy := c.PostForm("created_by")
	article := model.NewArticle()
	res := response.NewResponse(c)
	err := article.Create(title, desc, url, content, createdBy)
	if err != nil {
		res.ToError(errcode.AddError)
		c.Abort()
		return
	}
	res.ToResponse(errcode.Success)
	c.Abort()
	return
}
func Article_Update(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	title := c.PostForm("title")
	desc := c.PostForm("desc")
	url := c.PostForm("url")
	content := c.PostForm("content")
	modifiedBy := c.PostForm("modified_by")
	article := model.NewArticle()
	res := response.NewResponse(c)
	err := article.Update(id, title, desc, url, content, modifiedBy)
	if err != nil {
		res.ToError(errcode.UpdateError)
		c.Abort()
		return
	}
	res.ToResponse(errcode.Success)
	c.Abort()
	return
}

func Article_Delete(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	res := response.NewResponse(c)
	if id == 0 {
		res.ToError(errcode.ParamsError)
		c.Abort()
		return
	}
	article := model.NewArticle()
	err := article.Delete(id)
	if err != nil {
		res.ToError(errcode.DeleteError)
		c.Abort()
		return
	}
	res.ToResponse(errcode.Success)
	c.Abort()
	return
}

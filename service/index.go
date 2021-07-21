package service

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {

	c.HTML(http.StatusOK, "index/index.tmpl", gin.H{"title": "hahah"})
}

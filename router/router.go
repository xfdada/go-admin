package router

import (
	v1 "go-admin/api/v1"
	"go-admin/middleware"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.GET("/api/get_token", v1.GetToken)
	user := v1.NewUser()
	apiv1 := r.Group("api/v1")
	apiv1.Use(middleware.Jwt())
	{
		apiv1.GET("/user", user.List)
		apiv1.POST("/user", user.Create)
		apiv1.PUT("/user/:id", user.Update)
		apiv1.DELETE("/user/:id", user.Delete)
		apiv1.GET("/user/:id", user.Get)
	}
	return r
}

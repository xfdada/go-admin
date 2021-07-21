package router

import (
	v1 "go-admin/api/v1"
	_ "go-admin/docs"
	"go-admin/global"
	"go-admin/middleware"
	"go-admin/service"
	"net/http"

	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

func NewRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.GET("/doc/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.StaticFS("/uploads", http.Dir(global.Upload.UploadPath))
	r.GET("/api/get_token", v1.GetToken)
	r.GET("/", service.Index)
	r.LoadHTMLGlob("resource/view/**/*")
	user := v1.NewUser()
	r.POST("/api/upload", user.Upload)
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

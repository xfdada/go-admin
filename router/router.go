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
	r.Use(middleware.Logger())
	r.GET("/doc/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.StaticFS("/uploads", http.Dir(global.Upload.UploadPath))
	r.GET("/api/get_token", v1.GetToken)
	r.GET("/api/get_capt", v1.GetCapt)
	r.GET("/api/check/:uuid", v1.CheckUser)
	r.GET("/", service.Index)
	r.POST("/login", service.Login)
	r.POST("/sendmail", v1.SendEmail)
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
		apiv1.GET("/article/:id", v1.Article_Get)
		apiv1.POST("/article", v1.Article_Create)
		apiv1.DELETE("/article/:id", v1.Article_Delete)
		apiv1.PUT("/article/:id", v1.Article_Update)
		apiv1.GET("/article", v1.Article_List)
		apiv1.GET("/video_user/:id", v1.VideoGet)
		apiv1.GET("/video_user", v1.VideoUserList)
		apiv1.POST("/video_user", v1.VideoCreate)

	}
	return r
}

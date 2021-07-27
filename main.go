package main

import (
	"go-admin/global"
	"go-admin/router"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)


//@title go-admin快速开发示例
//@version	1.0
func main() {
	gin.SetMode(global.Server.Model)
	routers := router.NewRouter()
	s := &http.Server{
		Addr:           ":" + global.Server.Port,
		Handler:        routers,
		ReadTimeout:    global.Server.ReadTimeout * time.Second,
		WriteTimeout:   global.Server.WriteTimeout * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()
}

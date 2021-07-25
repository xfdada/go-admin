package main

import (
	"flag"
	"fmt"
	"go-admin/config"
	"go-admin/global"
	"go-admin/router"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

var (
	port    string
	runMode string
	cfgpath string
)

func init() {
	err := setupFlag()
	if err != nil {
		fmt.Println(err)
	}
	err = initConfig()
	if err != nil {
		fmt.Println(err)
	}

}

//@title go-admin快速开发示例
//@version	1.0
func main() {
	gin.SetMode(global.Server.Model)
	router := router.NewRouter()
	s := &http.Server{
		Addr:           ":" + global.Server.Port,
		Handler:        router,
		ReadTimeout:    global.Server.ReadTimeout * time.Second,
		WriteTimeout:   global.Server.WriteTimeout * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()
}

func initConfig() error {
	cfg, err := config.NewConfig(cfgpath)
	if err != nil {
		return err
	}
	err = cfg.ReadConfig("Server", &global.Server)
	if err != nil {
		return err
	}
	err = cfg.ReadConfig("JWT", &global.JWT)
	if err != nil {
		return err
	}
	err = cfg.ReadConfig("Upload", &global.Upload)
	if err != nil {
		return err
	}
	err = cfg.ReadConfig("Captcha", &global.Captcha)
	if err != nil {
		return err
	}
	if port != "" {
		global.Server.Port = port
	}
	if runMode != "" {
		global.Server.Model = runMode
	}
	return nil
}

func setupFlag() error {
	flag.StringVar(&port, "port", "", "启动端口")
	flag.StringVar(&runMode, "mode", "", "启动模式")
	flag.StringVar(&cfgpath, "config", "./", "配置文件的路径")
	flag.Parse()
	return nil
}

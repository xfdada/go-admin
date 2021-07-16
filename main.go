package main

import (
	"fmt"
	"go-admin/config"
	"go-admin/global"

	"github.com/gin-gonic/gin"
)

func init() {
	err := initConfig()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(global.Server.Model)
}

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"msg": "success"})
	})
	r.Run()
}

func initConfig() error {
	cfg, err := config.NewConfig()
	if err != nil {
		return err
	}
	err = cfg.ReadConfig("Server", &global.Server)
	if err != nil {
		return err
	}
	return nil
}

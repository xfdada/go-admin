package main

import (
	"fmt"
	"go-admin/config"
	"go-admin/global"
	"go-admin/router"
	"net/http"
	"time"
)

func init() {
	err := initConfig()
	if err != nil {
		fmt.Println(err)
	}
}

func main() {
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
	cfg, err := config.NewConfig()
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
	return nil
}

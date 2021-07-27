package global

import (
	"flag"
	"fmt"
	"go-admin/config"
	"go-admin/utils/loggers"
	"gorm.io/gorm"
)

var (
	Server  *config.Server
	Mysqls   *config.Mysql
	Captcha *config.Captcha
	JWT     *config.JWT
	Upload  *config.Upload
	DB		*gorm.DB
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
	err = initDB()
	if err != nil {
		loggers.Logs("连接数据库失败，错误详情是err:"+fmt.Sprintf("%v\n",err))
	}
}


func initConfig() error {
	cfg, err := config.NewConfig(cfgpath)
	if err != nil {
		return err
	}
	err = cfg.ReadConfig("Server", &Server)
	if err != nil {
		return err
	}
	err = cfg.ReadConfig("Mysql", &Mysqls)
	if err != nil {
		return err
	}
	err = cfg.ReadConfig("JWT", &JWT)
	if err != nil {
		return err
	}
	err = cfg.ReadConfig("Upload", &Upload)
	if err != nil {
		return err
	}
	err = cfg.ReadConfig("Captcha", &Captcha)
	if err != nil {
		return err
	}
	if port != "" {
		Server.Port = port
	}
	if runMode != "" {
		Server.Model = runMode
	}
	return nil
}

func initDB() error{
	var err error
	DB , err = NewDB()
	if err != nil {
		return err
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
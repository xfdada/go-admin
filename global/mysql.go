package global

import (
	"fmt"
	"go-admin/utils/loggers"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

func NewDB() (*gorm.DB,error){
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=%s&parseTime=%v&loc=Local",
		Mysqls.Username,Mysqls.Password,Mysqls.Host,Mysqls.DBName,Mysqls.Charset,Mysqls.ParseTime)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil{
		loggers.Logs("初始化连接数据库失败，错误详情是err:"+fmt.Sprintf("%v\n",err))
		return nil, err
	}
	sqlDB, err1 := db.DB()
	if err1 != nil{
		loggers.Logs("初始化连接数据库失败，错误详情是err:"+fmt.Sprintf("%v\n",err1))
		return nil, err1
	}
	sqlDB.SetMaxIdleConns(Mysqls.MaxIdleConns)
	sqlDB.SetMaxOpenConns(Mysqls.MaxOpenConns)
	sqlDB.SetConnMaxLifetime(time.Hour)
	return db,nil
}


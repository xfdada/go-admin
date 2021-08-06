package global

import (
	"fmt"
	"go-admin/utils/loggers"
	"go-admin/utils/logs"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"reflect"
	"time"
)

type Writer struct {
}

func (w Writer) Printf(format string, args ...interface{}) {
	s := fmt.Sprintf(format, args...)
	logs.SqlLog(s)
}

func NewDB() (*gorm.DB, error) {
	newLogger := logger.New(
		Writer{}, // io writer
		logger.Config{
			SlowThreshold: 100 * time.Millisecond, // 慢 SQL 阈值
			LogLevel:      logger.Info,            // Log level
			Colorful:      false,                  // 禁用彩色打印
		},
	)
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=%s&parseTime=%v&loc=Local",
		Mysqls.Username, Mysqls.Password, Mysqls.Host, Mysqls.DBName, Mysqls.Charset, Mysqls.ParseTime)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{Logger: newLogger})
	db.Callback().Create().Register("gorm:sing", createdTimeForCreateCallback)
	if err != nil {
		loggers.Logs("初始化连接数据库失败，错误详情是err:" + fmt.Sprintf("%v\n", err))
		return nil, err
	}
	sqlDB, err1 := db.DB()
	if err1 != nil {
		loggers.Logs("初始化连接数据库失败，错误详情是err:" + fmt.Sprintf("%v\n", err1))
		return nil, err1
	}

	sqlDB.SetMaxIdleConns(Mysqls.MaxIdleConns)
	sqlDB.SetConnMaxIdleTime(Mysqls.ConnMaxIdleTime * time.Second) //连接池空闲的最长时间
	sqlDB.SetConnMaxLifetime(Mysqls.ConnMaxLifetime * time.Second) //可重用连接的最长时间
	sqlDB.SetMaxOpenConns(Mysqls.MaxOpenConns)
	return db, nil
}

func createdTimeForCreateCallback(db *gorm.DB) {

	field := db.Statement.Schema.LookUpField("SingTime")
	_ = field.Set(db.Statement.ReflectValue, time.Now().Format("2006-01-02 15:04:05"))

	if db.Statement.Schema != nil {
		// 伪代码：裁剪图片并上传至 CDN
		for _, field := range db.Statement.Schema.Fields {
			switch db.Statement.ReflectValue.Kind() {
			case reflect.Slice, reflect.Array:
				for i := 0; i < db.Statement.ReflectValue.Len(); i++ {
					// 从字段获取值
					if fieldValue, isZero := field.ValueOf(db.Statement.ReflectValue.Index(i)); !isZero {
						fmt.Println(fieldValue)
					}
				}
			case reflect.Struct:
				// 从字段获取值
				if fieldValue, isZero := field.ValueOf(db.Statement.ReflectValue); isZero {
					if crop, ok := fieldValue.(string); ok {
						fmt.Println(crop)
					}
				}

				// 设置字段值
				_ = field.Set(db.Statement.ReflectValue, "newValue")
			}

		}

	}
}

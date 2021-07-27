package loggers

import (
	"log"
	"os"
)

func Logs(errs string) {
	/*
	   O_RDWR      读写模式打开文件
	   O_APPEND    写操作时将数据附加到文件尾部
	   O_CREATE    如果不存在将创建一个新文件
	*/
	logFile, err := os.OpenFile("./app.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0766)
	if err != nil {
		panic(err.Error())
	} else {
		// 方法二，log输出到文件
		logger := log.New(logFile, "[logger]", log.LstdFlags|log.Lshortfile|log.LUTC)
		// log输出到文件
		logger.Println(errs)
	}
	// 关闭文件
	defer logFile.Close()
}

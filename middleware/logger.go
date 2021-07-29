package middleware

import (
	"github.com/gin-gonic/gin"
	"go-admin/utils/logs"
	"time"
)

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {

		startTime := time.Now()
		// 处理请求
		c.Next()
		// 结束时间
		endTime := time.Now()
		// 执行时间
		latencyTime := endTime.Sub(startTime)
		// 请求方式
		reqMethod := c.Request.Method
		// 请求路由
		reqUri := c.Request.RequestURI
		// 状态码
		statusCode := c.Writer.Status()
		// 请求IP
		clientIP := c.ClientIP()

		logs.AppLog(statusCode, latencyTime, clientIP, reqMethod, reqUri)
	}

}

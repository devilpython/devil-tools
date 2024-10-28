package middle_ware

import (
	logger2 "devil-tools/logger"
	"devil-tools/utils"
	"github.com/gin-gonic/gin"
	"time"
	"devil-tools/goroutine_local"
)

//gin的日志中间件
func GinLogger() gin.HandlerFunc {
	logger := logger2.GetLoggerInstance()
	return func(c *gin.Context) {
		//防止数据溢出
		defer utils.RemoveAllGlobalData()
		//开始时间
		start := time.Now()
		//处理请求
		c.Next()
		// 结束时间
		end := time.Now()
		//执行时间
		latency := end.Sub(start)
		path := c.Request.URL.Path
		clientIP := c.ClientIP()
		method := c.Request.Method
		statusCode := c.Writer.Status()
		logger.Infof(" GOROUTINE_ID[%d] | %3d | %13v | %15s | %s  %s |",
			goroutine_local.GetGoroutineID(),
			statusCode,
			latency,
			clientIP,
			method, path,
		)
	}
}

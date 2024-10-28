package middle_ware

//import (
//	"github.com/gin-gonic/gin"
//	"github.com/zxfonline/buffpool"
//	"net/http"
//	"time"
//)

////超时设置中间件
//func Timeout(t time.Duration) gin.HandlerFunc {
//	return func(c *gin.Context) {
//		// sync.Pool
//		buffer := buffpool.GetBuff()
//
//		blw := &SimplebodyWriter{body: buffer, ResponseWriter: c.Writer}
//		c.Writer = blw
//
//		finish := make(chan struct{})
//		go func() { // 子协程只会将返回数据写入到内存buff中
//			c.Next()
//			finish <- struct{}{}
//		}()
//
//		select {
//		case <-time.After(t):
//			c.Writer.WriteHeader(http.StatusGatewayTimeout)
//			c.Abort()
//			// 1. 主协程超时退出。此时，子协程可能仍在运行
//			// 如果超时的话，buffer无法主动清除，只能等待GC回收
//		case <-finish:
//			// 2. 返回结果只会在主协程中被写入
//			blw.ResponseWriter.Write(buffer.Bytes())
//			buffpool.PutBuff(buffer)
//		}
//	}
//}

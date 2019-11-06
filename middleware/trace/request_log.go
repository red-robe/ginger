package trace

import "github.com/gin-gonic/gin"

// 请求追踪链的日志处理中间件
func RequestLog() gin.HandlerFunc {
	return func(c *gin.Context) {
		RequestStartLog(c)
		defer RequestEndLog(c)
		c.Next()
	}
}

func RequestStartLog(c *gin.Context) {

}

func RequestEndLog(c *gin.Context) {

}

package logger

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"time"
)

var timeFormat = "2019-11-09T23:02:28.844+0800"

func GingerWithZap(logger *zap.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		query := c.Request.URL.RawQuery
		postForm := c.Request.PostForm.Encode()

		// 生成traceId
		// traceId := common.GenTraceId(c.ClientIP())
		// c.Set("traceId", traceId)

		c.Next()

		end := time.Now()
		latency := end.Sub(start)
		if len(c.Errors) > 0 {
			for _, e := range c.Errors.Errors() {
				logger.Error(e)
			}
		} else {
			logger.Info("[Global Request Log]",
				zap.String("ip", c.ClientIP()),
				zap.Time("time", start.UTC()),
				zap.Duration("latency", latency),
				zap.Int("status", c.Writer.Status()),
				zap.String("method", c.Request.Method),
				zap.String("path", path),
				zap.String("query", query),
				zap.String("post", postForm),
				// zap.String("trace_id", traceId),
			)
		}
	}
}

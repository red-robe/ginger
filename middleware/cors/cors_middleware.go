package cors

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/gofuncchan/ginger/config"
	"time"
)

// Cors跨域请求处理中间件
func CORSMiddleware() gin.HandlerFunc {
	if config.CorsConf.AllowAllOrigins {
		// 允许所有跨域请求
		// same as
		// config := cors.DefaultConfig()
		// config.AllowAllOrigins = true
		return cors.Default()
	} else {
		// 无配置则使用默认的配置
		if config.CorsConf == nil {
			return cors.Default()
		}

		// 如Request Header 无携带Origin字段，默认不是跨域请求CORS request
		return cors.New(cors.Config{
			AllowOrigins:     config.CorsConf.AllowOrigins,
			AllowMethods:     config.CorsConf.AllowMethods,
			AllowHeaders:     config.CorsConf.AllowHeaders,
			ExposeHeaders:    config.CorsConf.ExposeHeaders,
			AllowCredentials: config.CorsConf.AllowCredentials,
			MaxAge:           time.Second * time.Duration(config.CorsConf.MaxAge),
		})
	}
}

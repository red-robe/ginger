package router

import (
	"github.com/gofuncchan/ginger/handler"
	"github.com/gofuncchan/ginger/middleware/auth"
	"github.com/gin-gonic/gin"
	"net/http"
)

// 设置router
func Router(r *gin.Engine) {
	// 请求无路由处理时
	r.NoRoute(func(c *gin.Context) {
		c.HTML(http.StatusNotFound,"error.tmpl",gin.H{"error_code":http.StatusNotFound,"error_title":"Not Found","error_content":"您请求的页面不在本星球..."})
	})
	// 请求无方法处理时
	r.NoMethod(func(c *gin.Context) {
		c.HTML(http.StatusMethodNotAllowed,"error.tmpl",gin.H{"error_code":http.StatusMethodNotAllowed,"error_title":"Method Not Allowed","error_content":"请注意您的发射姿势..."})
	})

	// 默认根路径
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"title": "Ginger - A Useful Gin Scaffold",
			"content":"Ginger - A Useful Gin Scaffold.\n Wish you a pleasant use!!!",
		})
	})

	// 设置静态资源访问路径
	r.Static("/static", "./static")
	r.Static("/asset", "./asset")

	// 非鉴权相关
	// 注册
	r.GET("/signup", handler.SignUp)
	r.POST("/signup", handler.SignUp)

	// 登录
	r.GET("/signin", handler.SignIn)
	r.POST("/signin", handler.SignIn)

	// 退出登录
	r.GET("/signout", auth.AuthMiddleware(), handler.SignOut)

	// 三方登录
	r.GET("oauth2/wechat/signin", handler.WechatSignIn)
	r.GET("oauth2/qq/signin", handler.QQSignIn)
	r.GET("oauth2/weibo/signin", handler.WeiboSignIn)



}

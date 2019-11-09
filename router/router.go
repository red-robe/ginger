package router

import (
	"ginger/handlers"
	"ginger/middleware/auth"
	"github.com/gin-gonic/gin"
	"net/http"
)

// 设置router
func Router(r *gin.Engine) {
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
	r.GET("/signup", handlers.SignUp)
	r.POST("/signup", handlers.SignUp)

	// 登录
	r.GET("/signin", handlers.SignIn)
	r.POST("/signin", handlers.SignIn)

	// 退出登录
	r.GET("/signout", auth.AuthMiddleware(), handlers.SignOut)

}

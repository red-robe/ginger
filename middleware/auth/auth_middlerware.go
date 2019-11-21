package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/gofuncchan/ginger/cache"
	"github.com/gofuncchan/ginger/util/jwt"
	"net/http"
	"strconv"
)

// 用户鉴权中间件
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		// 从http头获取token string

		tkStr := c.GetHeader("token")

		if tkStr == "" {
			c.Abort()
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "token string on http header is required!",
			})
			return
		}

		// 解码校验token是否合法
		customerClaim, err := jwt.JwtService.Decode(tkStr)
		if err != nil {
			c.Abort()
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "token string is invalid!",
			})
			return
		}

		// 在Redis查找token是否存在，不存在或过期则返回-1，还存在则返回token值
		key := "user_token_" + strconv.Itoa(int(customerClaim.TokenUserClaim.Id))
		token := cache.GetToken(key)
		// fmt.Println(key)
		// fmt.Println(tkStr)

		// 校验客户端token和服务端缓存的token
		if token != tkStr {
			c.Abort()
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "Token is not exist,please retry to sign in!",
			})
			return
		}

		// 鉴权通过后设置用户信息
		c.Set("tkStr", tkStr)
		c.Set("userClaim", customerClaim.TokenUserClaim)

		c.Next()
	}
}

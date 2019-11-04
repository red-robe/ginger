package handlers

import (
	"fmt"
	"ginger/cache"
	"ginger/common"
	"ginger/model"
	"ginger/util"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

/*
注册处理,默认邮箱注册
*/
type SignUpForm struct {
	Name       string `form:"name" binding:"required"`
	Email      string `form:"email" binding:"required,email"`
	Password   string `form:"password" binding:"required,alphanum"`
	RePassword string `form:"re_password" binding:"required,eqfield=Password"`
}

func SignUp(c *gin.Context) {
	// 请求注册页面
	if c.Request.Method == "GET" {
		c.HTML(http.StatusOK, "register.tmpl", gin.H{
			"title": "User Register Page",
		})

	} else if c.Request.Method == "POST" {
		// 表单参数校验
		form := new(SignUpForm)
		// _ = c.Bind(form) //MustBindWith()
		if err := c.ShouldBind(form); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// 密码哈希计算
		passHash, salt := util.GenPassHash(form.Password)
		// 创建用户
		id := model.CreateUserByEmail(form.Name, form.Email, passHash, salt)
		if id == -1 {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "Register Error,Please Retry!",
			})
		}

		// 创建登录token并返回
		userClaim := util.TokenUserClaim{
			Id:    id,
			Name:  form.Name,
			Email: form.Email,
		}
		tkStr, err := util.JwtService.Encode(userClaim)
		common.Eh("tk.Encode", err)

		// Redis存储token保存登录状态
		userKey := "user_token_" + strconv.Itoa(int(id))
		cache.SetToken(userKey, tkStr)

		c.JSON(http.StatusOK, gin.H{
			"token": tkStr,
		})

	} else {
		c.JSON(http.StatusMethodNotAllowed, gin.H{
			"message": "Only Allow Get Or Post Method",
		})
	}
}

/*
登录处理
*/
type SignInForm struct {
	Email    string `form:"email" binding:"required,email"`
	PassWord string `form:"password" binding:"required,alphanum"`
}

func SignIn(c *gin.Context) {
	if c.Request.Method == "GET" {
		// 请求登录页面
		c.HTML(http.StatusOK, "login.tmpl", gin.H{
			"title": "User Login Page",
		})
	} else if c.Request.Method == "POST" {
		// 请求登录参数验证
		form := new(SignInForm)
		if err := c.ShouldBind(form); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// 与数据库账号密码鉴权
		// 1.根据邮箱获取哈希密码与盐值
		userInfo := model.GetUserInfoByEmail(form.Email)
		// 2.将用户密码与盐值哈希计算并与数据库密码进行比较
		b := util.IsValidPasswd(form.PassWord, userInfo.Salt, userInfo.Password)
		if !b {
			c.JSON(http.StatusForbidden, gin.H{
				"message": "Email Or Password Error,Please Retry!",
			})
			return
		}

		// 校验成功则生成token并返回
		// 创建登录token并返回
		userClaim := util.TokenUserClaim{
			Id:    int64(userInfo.ID),
			Name:  userInfo.Name,
			Email: userInfo.Email,
		}

		tkStr, err := util.JwtService.Encode(userClaim)
		common.Eh("tk.Encode", err)

		// Redis存储token保存登录状态
		userKey := "user_token_" + strconv.Itoa(int(userInfo.ID))
		cache.SetToken(userKey, tkStr)

		c.JSON(http.StatusOK, gin.H{
			"token": tkStr,
		})

	} else {
		c.JSON(http.StatusMethodNotAllowed, gin.H{
			"message": "Only Allow Get Or Post Method",
		})
	}
}

/*
退出登录处理
*/
func SignOut(c *gin.Context) {
	// 从header获取token字段，在redis删除键
	tkStr := c.GetHeader("token")

	fmt.Println(tkStr)

	// 解码获取id
	claim, err := util.JwtService.Decode(tkStr)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "token string is invalid!",
		})
		return
	}

	key := "user_token_" + strconv.Itoa(int(claim.TokenUserClaim.Id))
	delCount := cache.DeleteToken(key)

	if delCount > 0 {
		c.JSON(http.StatusOK, gin.H{
			"message": "Sign Out Successful!",
		})
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Sign Out Error!",
		})
	}
}

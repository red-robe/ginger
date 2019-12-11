package handler

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/gofuncchan/ginger/cache"
	"github.com/gofuncchan/ginger/common"
	"github.com/gofuncchan/ginger/model/userOauthModel"
	"github.com/gofuncchan/ginger/oauth2"
	"github.com/gofuncchan/ginger/util/e"
	"github.com/gofuncchan/ginger/util/jwt"
	"strconv"
)

/*
登录流程:在客户端发起微博登录请求跳转到微博用户授权页面，并设置重定向到指定的前端链接，接收授权码code，和state值，state在前端自行校验，校验完毕发送code值到服务端完成注册或登录，返回token string

前端工作：
请求链接：https://api.weibo.com/oauth2/authorize
请求方式：GET/POST
请求参数：
参数             必选	    类型及范围	    说明
client_id	    true	string	    申请应用时分配的AppKey。
redirect_uri	true	string	    授权回调地址，站外应用需与设置的回调地址一致，站内应用需填写canvas page的地址。
scope	        false	string	    申请scope权限所需参数，可一次申请多个scope权限，用逗号分隔。使用文档
state	        false	string	    用于保持请求和回调的状态，在回调时，会在Query Parameter中回传该参数。开发者可以用这个参数验证请求有效性，也可以记录用户请求授权页前的位置。这个参数可用于防止跨站请求伪造（CSRF）攻击
display	        false	string	    授权页面的终端类型，取值见下面的说明。
forcelogin	    false	boolean	    是否强制用户重新登录，true：是，false：否。默认false。
language	    false	string	    授权页语言，缺省为中文简体版，en为英文版。英文版测试中，开发者任何意见可反馈至 @微博API

display说明：
参数取值	类型说明
default	默认的授权页面，适用于web浏览器。
mobile	移动终端的授权页面，适用于支持html5的手机。注：使用此版授权页请用 https://open.weibo.cn/oauth2/authorize 授权接口
wap	wap版授权页面，适用于非智能手机。
client	客户端版本授权页面，适用于PC桌面应用。
apponweibo	默认的站内应用授权页，授权后不返回access_token，只刷新站内应用父框架。
返回数据
返回值字段	字段类型	字段说明
code	string	用于第二步调用oauth2/access_token接口，获取授权后的access token。
state	string	如果传递参数，会回传该参数。
示例
//请求
https://api.weibo.com/oauth2/authorize?client_id=123050457758183&redirect_uri=http://www.example.com/response&response_type=code

//同意授权后会重定向
http://www.example.com/response&code=CODE
*/

type WeiboSignInParams struct {
	Code string `form:"code" binding:"required,gt=0"`
}

// 获取微博用户授权信息并注册或登录返回Token String
func WeiboSignIn(c *gin.Context) {
	// validate request params
	form := new(WeiboSignInParams)
	if err := c.ShouldBind(form); err != nil {
		common.ResponseInvalidParam(c, err)
		return
	}

	// 使用微博授权码code开始鉴权流程并获取微博用户信息
	userInfo := oauth2.WeiboOAuth2Manager.Authorize(form.Code)

	// 已获取到三方平台用户信息，进入鉴权流程
	info, err := userOauthModel.GetUserOauthInfo(userInfo.UserInfo.OpenId, userInfo.UserInfo.UnionId, userOauthModel.WeiboPlatform)
	if err != nil {
		common.ResponseServerError(c, errors.New("wechat oauth2 login error"))
		return
	}
	var userClaim jwt.TokenUserClaim
	var userKey string
	if info != nil {
		// 该三方账号已经注册过，走登录流程，获取用户信息，生成TokenString返回
		userInfo := userOauthModel.GetUserInfoByUserOauthId(int64(info.ID))
		if userInfo == nil {
			common.ResponseServerError(c, errors.New("wechat oauth2 login error"))
			return
		}

		// 创建登录token并返回
		userClaim = jwt.TokenUserClaim{
			Id:     int64(userInfo.ID),
			Name:   userInfo.Name,
			Avatar: userInfo.Avatar,
		}
		userKey = cache.UserTokenCacheKeyPrefix + strconv.Itoa(int(userInfo.ID))

	} else {
		// 该三方账号未注册，走注册流程，新增用户信息，生成TokenString返回
		userId := userOauthModel.CreateUserByOauth2UserInfo(info.NickName, info.AvatarURL, info.AccessToken, info.OpenID, info.UnionID, int64(info.Gender), userOauthModel.WeiboPlatform)
		if userId < 0 {
			common.ResponseServerError(c, errors.New("register user info error"))
			return
		}

		// 创建登录token并返回
		userClaim = jwt.TokenUserClaim{
			Id:     userId,
			Name:   info.NickName,
			Avatar: info.AvatarURL,
		}
		userKey = cache.UserTokenCacheKeyPrefix + strconv.Itoa(int(userId))

	}

	tkStr, err := jwt.JwtService.Encode(userClaim)
	if !e.Eh(err) {
		common.ResponseServerError(c, errors.New("jwt token string encode error"))
		return
	}
	// Redis存储token保存登录状态
	cache.SetToken(userKey, tkStr)

	common.ResponseOk(c, gin.H{"token": tkStr})
}

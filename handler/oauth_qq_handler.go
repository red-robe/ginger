package handler

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/gofuncchan/ginger/cache/tokenCache"
	"github.com/gofuncchan/ginger/common"
	"github.com/gofuncchan/ginger/model/userOauthModel"
	"github.com/gofuncchan/ginger/oauth2"
	"github.com/gofuncchan/ginger/util/e"
	"github.com/gofuncchan/ginger/util/jwt"
	"strconv"
)

/*
登录流程:在客户端发起QQ登录请求跳转到qq用户授权页面，并设置重定向到指定的前端链接，接收授权码code，和state值，state在前端自行校验，校验完毕发送code值到服务端完成注册或登录，返回token string

前端工作：获取Authorization Code
请求地址：
PC网站：https://graph.qq.com/oauth2.0/authorize
请求方法：
GET
请求参数：
请求参数请包含如下内容：

参数	            是否必须	含义
response_type	必须	    授权类型，此值固定为“code”。
client_id	    必须	    申请QQ登录成功后，分配给应用的appid。
redirect_uri	必须	    成功授权后的回调地址，必须是注册appid时填写的主域名下的地址，建议设置为网站首页或网站的用户中心。注意需要将url进行URLEncode。
state	        必须	    client端的状态值。用于第三方应用防止CSRF攻击，成功授权后回调时会原样带回。请务必严格按照流程检查用户与state参数状态的绑定。
scope	        可选	    请求用户授权时向用户显示的可进行授权的列表。

可填写的值是API文档中列出的接口，以及一些动作型的授权（目前仅有：do_like），如果要填写多个接口名称，请用逗号隔开。
例如：scope=get_user_info,list_album,upload_pic,do_like
不传则默认请求对接口get_user_info进行授权。
建议控制授权项的数量，只传入必要的接口名称，因为授权项越多，用户越可能拒绝进行任何授权。

display	可选	仅PC网站接入时使用。
用于展示的样式。不传则默认展示为PC下的样式。
如果传入“mobile”，则展示为mobile端下的样式。
返回说明：
1. 如果用户成功登录并授权，则会跳转到指定的回调地址，并在redirect_uri地址后带上Authorization Code和原始的state值。如：
PC网站：http://graph.qq.com/demo/index.jsp?code=9A5F************************06AF&state=test
注意：此code会在10分钟内过期。

*/
type QQSignInParams struct {
	Code string `form:"code" binding:"required,gt=0"`
}

// 获取QQ用户授权信息并注册或登录返回Token String
func QQSignIn(c *gin.Context) {
	// validate request params
	form := new(QQSignInParams)
	if err := c.ShouldBind(form); err != nil {
		common.ResponseInvalidParam(c, err)
		return
	}

	// 使用qq回调授权码code开始鉴权流程并获取QQ用户信息
	userInfo := oauth2.QQOAuth2Manager.Authorize(form.Code)

	// 已获取到三方平台用户信息，进入鉴权流程
	info, err := userOauthModel.GetUserOauthInfo(userInfo.UserInfo.OpenId, userInfo.UserInfo.UnionId, userOauthModel.QQPlatform)
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
		userKey = tokenCache.UserTokenCacheKeyPrefix + strconv.Itoa(int(userInfo.ID))

	} else {
		// 该三方账号未注册，走注册流程，新增用户信息，生成TokenString返回
		userId := userOauthModel.CreateUserByOauth2UserInfo(info.NickName, info.AvatarURL, info.AccessToken, info.OpenID, info.UnionID, int64(info.Gender), userOauthModel.QQPlatform)
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
		userKey = tokenCache.UserTokenCacheKeyPrefix + strconv.Itoa(int(userId))

	}

	tkStr, err := jwt.JwtService.Encode(userClaim)
	if !e.Eh(err) {
		common.ResponseServerError(c, errors.New("jwt token string encode error"))
		return
	}
	// Redis存储token保存登录状态
	tokenCache.SetToken(userKey, tkStr)

	common.ResponseOk(c, gin.H{"token": tkStr})
	return
}

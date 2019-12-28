package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/gofuncchan/ginger/cache/tokenCache"
	"github.com/gofuncchan/ginger/common"
	"github.com/gofuncchan/ginger/model/userOauthModel"
	"github.com/gofuncchan/ginger/oauth2"
	"github.com/gofuncchan/ginger/util/e"
	"github.com/gofuncchan/ginger/util/jwt"
	"github.com/pkg/errors"
	"strconv"
)

/*
登录流程：
1.使用客户端sdk发起微信oauth2请求（如网站端用js sdk在本页面获取二维码），用户微信扫码并确认后在客户端定义的的redirect_uri接收accessTokenCode和state，state可在客户端校验
网站前端步骤
步骤1：在页面中先引入如下JS文件（支持https）：
http://res.wx.qq.com/connect/zh_CN/htmledition/js/wxLogin.js
步骤2：在需要使用微信登录的地方实例以下JS对象：
 var obj = new WxLogin({
 self_redirect:true,
 id:"login_container",
 appid: "",
 scope: "",
 redirect_uri: "",
  state: "",
 style: "",
 href: ""
 });

参数				是否必须	说明
self_redirect	否		true：手机点击确认登录后可以在 iframe 内跳转到 redirect_uri，false：手机点击确认登录后可以在 top window 跳转到 redirect_uri。默认为 false。
id				是		第三方页面显示二维码的容器id
appid			是		应用唯一标识，在微信开放平台提交应用审核通过后获得
scope			是		应用授权作用域，拥有多个作用域用逗号（,）分隔，网页应用目前仅填写snsapi_login即可
redirect_uri	是		重定向地址，需要进行UrlEncode
state			否		用于保持请求和回调的状态，授权请求后原样带回给第三方。该参数可用于防止csrf攻击（跨站请求伪造攻击），建议第三方带上该参数，可设置为简单的随机数加session进行校验
style			否		提供"black"、"white"可选，默认为黑色文字描述。
href			否		自定义样式链接，第三方可根据实际需求覆盖默认样式。

2.客户端使用微信返回的accessTokenCode请求后端登录，由后端做登录或注册处理，处理完成返回JWT TokenString，完成登录流程
*/

type WechatSignInParams struct {
	Code string `form:"code" binding:"required,gt=0"`
}

// 获取微信用户授权信息并注册或登录返回Token String
func WechatSignIn(c *gin.Context) {
	// validate request params
	form := new(WechatSignInParams)
	if err := c.ShouldBind(form); err != nil {
		common.ResponseInvalidParam(c, err)
		return
	}

	// 使用微信授权码code开始鉴权流程并获取微信用户信息
	userInfo := oauth2.WechatOAuth2Manager.Authorize(form.Code)

	// 已获取到三方平台用户信息，进入鉴权流程
	info, err := userOauthModel.GetUserOauthInfo(userInfo.UserInfo.OpenId, userInfo.UserInfo.UnionId, userOauthModel.WechatPlatform)
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
		userId := userOauthModel.CreateUserByOauth2UserInfo(info.NickName, info.AvatarURL, info.AccessToken, info.OpenID, info.UnionID, int64(info.Gender), userOauthModel.WechatPlatform)
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
}

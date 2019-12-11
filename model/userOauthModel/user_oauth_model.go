package userOauthModel

import (
	"github.com/gofuncchan/ginger/dao/mysql"
	"github.com/gofuncchan/ginger/dao/mysql/schema"
	"github.com/gofuncchan/ginger/util/e"
	"time"
)

const (
	WechatPlatform = 1
	QQPlatform     = 2
	WeiboPlatform  = 3
)

/*
通过openid和unionid获取三方账号的用户信息
*/
func GetUserOauthInfo(openId, unionId string, platform int64) *schema.UserOauth2 {
	where := map[string]interface{}{
		"platform": platform,
		"open_id":  openId,
		"union_id": unionId,
	}

	oauth2Result := new(schema.UserOauth2)
	err := mysql.GetOne(schema.UserOauth2TableName, where, oauth2Result)
	if !e.Em(err) {
		return nil
	}
	return oauth2Result
}

/*
TODO 连表获取用户表信息
*/
func GetUserInfoByUserOauthId(oauth_id int64) *schema.User {
	var err error
	// 从关联表获取userId
	where := map[string]interface{}{
		"oauth_user_id": oauth_id,
	}
	bindingResult := new(schema.UserOauth2Binding)
	err = mysql.GetOne(schema.UserOauth2BindingTableName, where, bindingResult)
	if !e.Em(err) {
		return nil
	}
	userId := bindingResult.UserID

	// 根据userId获取用户信息
	where = map[string]interface{}{
		"id": userId,
	}
	userResult := new(schema.User)
	err = mysql.GetOne(schema.UserTableName, where, userResult)
	if !e.Em(err) {
		return nil
	}

	return userResult
}

/*
TODO 关联用户表插入信息
创建三方用户信息和绑定信息
*/
func CreateUserByOauth2(name, avatar, accessToken, openId, unionId string, gender, userId, platform int64) int64 {
	var data []map[string]interface{}
	data = append(data, map[string]interface{}{
		"platform":     platform,
		"nick_name":    name,
		"avatar":       avatar,
		"access_token": accessToken,
		"open_id":      openId,
		"union_id":     unionId,
		"gender":       gender,
		"create_at":    time.Now(),
		"update_at":    time.Now(),
	})

	oAuthId, err := mysql.Insert(schema.UserOauth2TableName, data)
	if !e.Em(err) {
		return -1
	}

	var relationData []map[string]interface{}
	relationData = append(relationData, map[string]interface{}{
		"user_id":       userId,
		"oauth_user_id": oAuthId,
	})

	_, err = mysql.Insert(schema.UserTableName, data)
	if !e.Em(err) {
		return -1
	}

	return oAuthId
}

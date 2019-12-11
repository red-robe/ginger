package userOauthModel

import (
	"github.com/didi/gendry/builder"
	"github.com/didi/gendry/scanner"
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
func GetUserOauthInfo(openId, unionId string, platform int64) (*schema.UserOauth2,error) {
	where := map[string]interface{}{
		"platform": platform,
		"open_id":  openId,
		"union_id": unionId,
	}

	oauth2Result := new(schema.UserOauth2)
	err := mysql.GetOne(schema.UserOauth2TableName, where, nil, oauth2Result)
	if !e.Em(err) {
		return nil,err
	}
	return oauth2Result,nil
}

/*
连表获取用户表信息
*/
func GetUserInfoByUserOauthId(oauthId int64) *schema.User {
	// 联合查询user_oauth_bingding和user表，where equal oauth_id
	// select * form user where id=(select user_id from user_oauth2_binding where oauth_user_id=oauth_id)
	var err error

	condition, values, err := builder.NamedQuery("select * form user where id=(select user_id from user_oauth2_binding where oauth_user_id={{oauthId}})", map[string]interface{}{
		"oauthId": oauthId,
	})

	rows, err := mysql.Db.Query(condition, values)
	if !e.Em(err) || nil == rows {
		return nil
	}
	defer rows.Close()

	var results = make([]*schema.User, 0)

	err = scanner.Scan(rows, results)
	if !e.Em(err) {
		return nil
	}

	return results[0]
}

/*
关联用户表插入信息
创建三方用户信息和绑定信息
*/
func CreateUserByOauth2UserInfo(name, avatar, accessToken, openId, unionId string, gender, platform int64) int64 {

	// 开启事务插入三个用户相关表
	tx, err := mysql.Db.Begin()
	if !e.Em(err) {
		return -1
	}

	// 1.插入三方平台用户信息表
	var oAuthData []map[string]interface{}
	oAuthData = append(oAuthData, map[string]interface{}{
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

	oAuthId, err := mysql.TxInsert(tx, schema.UserOauth2TableName, oAuthData)
	if !e.Em(err) {
		return -1
	}

	// 2.插入user表
	var userData []map[string]interface{}
	userData = append(userData, map[string]interface{}{
		"name":      name,
		"avatar":    avatar,
		"create_at": time.Now(),
		"update_at": time.Now(),
	})

	userId, err := mysql.TxInsert(tx, schema.UserTableName, userData)
	if !e.Em(err) {
		return -1
	}

	// 3.插入关系表
	var relationData []map[string]interface{}
	relationData = append(relationData, map[string]interface{}{
		"user_id":       userId,
		"oauth_user_id": oAuthId,
	})

	_, err = mysql.TxInsert(tx, schema.UserOauth2BindingTableName, relationData)
	if !e.Em(err) {
		return -1
	}

	err = tx.Commit()
	if !e.Em(err) {
		err := tx.Rollback()
		if !e.Em(err) {
			// 回退错误
			return -2
		}
		// 提交事务错误
		return -1
	}

	return userId

}

package userModel

import (
	"github.com/gofuncchan/ginger/dao/mysql"
	"github.com/gofuncchan/ginger/dao/mysql/schema"
	"github.com/gofuncchan/ginger/util/e"
	"time"
)

// 邮箱注册创建用户
func CreateUserByEmail(name, email, passwd, salt string) int64 {
	var data []map[string]interface{}
	data = append(data, map[string]interface{}{
		"name":      name,
		"email":     email,
		"password":  passwd,
		"salt":      salt,
		"create_at": time.Now(),
		"update_at": time.Now(),
	})

	id, err := mysql.Insert(schema.UserTableName, data)
	if !e.Em(err) {
		return -1
	}
	return id
}

// 手机注册创建用户
func CreateUserByPhone(name, phone, passwd, salt string) int64 {
	var data []map[string]interface{}
	data = append(data, map[string]interface{}{
		"name":      name,
		"phone":     phone,
		"passwd":    passwd,
		"salt":      salt,
		"create_at": time.Now(),
		"update_at": time.Now(),
	})

	id, err := mysql.Insert(schema.UserTableName, data)
	if !e.Em(err) {
		return -1
	}
	return id
}


// 更新用户信息
func UpdateUserInfo(id int, data map[string]interface{}) bool {
	where := map[string]interface{}{
		"id": uint(id),
	}

	_, err := mysql.Update(schema.UserOauth2BindingTableName, where, data)

	return e.Em(err)
}

// 根据user_id获取用户信息
func GetUserInfoById(id int64) *schema.User {
	where := map[string]interface{}{
		"id": id,
	}
	userResult := new(schema.User)
	err := mysql.GetOne(schema.UserTableName, where, nil,userResult)
	if !e.Em(err) {
		return nil
	}
	return userResult
}

// 根据邮箱和密码验证用户登录
func GetUserInfoByEmail(email string) *schema.User {
	where := map[string]interface{}{
		"email": email,
	}
	userResult := new(schema.User)
	err := mysql.GetOne(schema.UserTableName, where, nil,userResult)
	if !e.Em(err) {
		return nil
	}
	return userResult
}

// 根据手机和密码验证用户登录
func GetUserInfoByPhone(phone string) *schema.User {
	where := map[string]interface{}{
		"phone": phone,
	}
	userResult := new(schema.User)
	err := mysql.GetOne(schema.UserTableName, where, nil,userResult)
	if !e.Em(err) {
		return nil
	}
	return userResult
}

// 根据邮箱验证用户是否已存在
func IsUserExistByEmail(email string) bool {
	where := map[string]interface{}{
		"email": email,
	}

	count, err := mysql.GetCount(schema.UserTableName, where)
	if !e.Em(err) {
		return false
	}
	return count == 1
}

// 根据手机验证用户是否已存在
func IsUserExistByPhone(phone string) bool {
	where := map[string]interface{}{
		"phone": phone,
	}

	count, err := mysql.GetCount(schema.UserTableName, where)
	if !e.Em(err) {
		return false
	}
	return count == 1
}

// 根据用户昵称验证是否已存在
func IsUserExistByName(name string) bool {
	where := map[string]interface{}{
		"name": name,
	}

	count, err := mysql.GetCount(schema.UserTableName, where)
	if !e.Em(err) {
		return false
	}
	return count == 1
}

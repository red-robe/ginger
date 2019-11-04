package util

import (
	"crypto/md5"
	"crypto/sha1"
	"fmt"
	"io"
	"math/rand"
	"strconv"
	"time"
)

//生成随机40位哈希
func GenHash(src string) string {
	//1.获取当前时间戳
	unix := time.Now().Unix()
	//2.将文件名和时间戳一起计算md5等到前32位十六进制字符
	hash := md5.New()
	io.WriteString(hash,src)
	io.WriteString(hash,strconv.Itoa(int(unix)))
	hb := hash.Sum(nil)

	//获取时间戳前8位字符
	ub := strconv.Itoa(int(unix))[:8]

	//组合输出40位哈希字符
	s := fmt.Sprintf("%x%s", hb, ub)

	return s
}

//用户密码加盐生成哈希
func GenPassHash(src string) (hashStr,salt string)  {
	//获取随机盐值字符串
	salt = getRandomString(4)

	hash := sha1.New()

	io.WriteString(hash,src)
	io.WriteString(hash,salt)

	hashBytes := hash.Sum(nil)
	//组合输出40位哈希字符
	hashStr = fmt.Sprintf("%x", hashBytes)
	return
}

//生成盐值
func  getRandomString(l int) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyz"
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < l; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}

//校验密码
func IsValidPasswd(passStr,salt,passHash string) bool {
	//重新计算密码哈希，与之前的校验
	hash := sha1.New()
	io.WriteString(hash,passStr)
	io.WriteString(hash,salt)
	hashBytes := hash.Sum(nil)
	//组合输出40位哈希字符
	hashStr := fmt.Sprintf("%x", hashBytes)

	return hashStr == passHash
}


package util

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

/*
返回体封装
*/

// 定义用户级别的返回码，并枚举其信息输出
const (
	ResponseCodeOk RespCode  = iota
	ResponseCodeClientError
	ResponseCodeServerError

	ResponseCodeInvalidParam
	ResponseCodeUnAuthorized
	ResponseCodeMethodNotAllowed
)

type RespCode int

func (code RespCode)String() string {
	return []string{
		"ok",
		"client error",
		"server error",
		"invalid request parameters",
		"unauthorized",
		"method not allowed",
	}[code]
}

// 统一返回消息体
type RespStruct struct {
	Code RespCode `json:"code"`
	Msg string `json:"code_msg"`
	Error string `json:"error"`
	Data interface{} `json:"data"`
}

// 通用响应封装
func CommonResponse(c *gin.Context,code RespCode,httpCode int,data interface{},err error)  {
	respStruct := new(RespStruct)
	respStruct.Code = code
	respStruct.Msg = code.String()
	if err != nil {
		respStruct.Error = err.Error()
	}
	respStruct.Data = data

	c.JSON(httpCode,respStruct)
}


// 请求成功响应的封装 （http 200）
func ResponseOk(c *gin.Context,data interface{}) {
	CommonResponse(c,ResponseCodeOk,http.StatusOK,data,nil)
}

// 客户端错误响应的封装 （http 400）
func ResponseClientError(c *gin.Context,err error)  {
	CommonResponse(c,ResponseCodeClientError,http.StatusBadRequest,nil,err)
}

// 服务端错误响应的封装 （http 500）
func ResponseServerError(c *gin.Context,err error)  {
	CommonResponse(c,ResponseCodeServerError,http.StatusInternalServerError,nil,err)
}

// 参数校验错误响应体封装
func ResponseInvalidParam(c *gin.Context,err error)  {
	CommonResponse(c,ResponseCodeInvalidParam,http.StatusBadRequest,nil,err)
}

// 未鉴权访问
func ResponseUnAuthorized(c *gin.Context,err error)  {
	CommonResponse(c,ResponseCodeUnAuthorized,http.StatusUnauthorized,nil,err)
}

func ResponseMethodNotAllowed(c *gin.Context,err error)  {
	CommonResponse(c,ResponseCodeMethodNotAllowed,http.StatusMethodNotAllowed,nil,err)
}
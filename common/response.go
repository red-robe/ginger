package common

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

/*
返回体封装
*/

// 定义用户级别的返回码，并枚举其信息输出
const (
	ResponseCodeOk RespCode = iota
	ResponseCodeClientError
	ResponseCodeServerError

	ResponseCodeInvalidParam
	ResponseCodeUnAuthorized
	ResponseCodeMethodNotAllowed
	ResponseCodeEmptyResult

	ResponseCodeHandlerError
	ResponseCodeModelError
	ResponseCodeRepositoryError
	ResponseCodeCacheError
	ResponseCodeCasbinError
	ResponseCodeAuthMiddlerwareError
	ResponseCodeCorsMiddlerwareError
	ResponseCodeRbacMiddlerwareError
)

type RespCode int

func (code RespCode) String() string {
	return []string{
		"ok",
		"client error",
		"server error",
		"invalid request parameters",
		"unauthorized",
		"method not allowed",
		"empty result or info not found",
		"server error on handler",
		"server error on mysql model",
		"server error on mongo repository",
		"server error on redis cache",
		"server error on casbin api",
		"error on auth middleware",
		"error on cors middleware",
		"error on rbac middkeware",
	}[code]
}

// 统一返回消息体
type RespStruct struct {
	Code  RespCode    `json:"code"`
	Msg   string      `json:"message"`
	Error string      `json:"error"`
	Data  interface{} `json:"data"`
}

// 通用响应封装
func CommonResponse(c *gin.Context, code RespCode, httpCode int, data interface{}, errorString string) {
	respStruct := new(RespStruct)
	respStruct.Code = code
	respStruct.Msg = code.String()
	if errorString != "" {
		respStruct.Error = errorString
	}
	respStruct.Data = data

	c.JSON(httpCode, respStruct)
}


// 请求成功响应的封装 （http 200）
func ResponseOk(c *gin.Context, data interface{}) {
	CommonResponse(c, ResponseCodeOk, http.StatusOK, data, "")
}

// 客户端错误响应的封装 （http 400）
func ResponseClientError(c *gin.Context, errorContent string) {
	CommonResponse(c, ResponseCodeClientError, http.StatusBadRequest, nil, errorContent)
}

// 服务端错误响应的封装 （http 500）
func ResponseServerError(c *gin.Context, errorContent string) {
	CommonResponse(c, ResponseCodeServerError, http.StatusInternalServerError, nil, errorContent)
}

// 参数校验错误响应体封装
func ResponseInvalidParam(c *gin.Context, errorContent string) {
	CommonResponse(c, ResponseCodeInvalidParam, http.StatusBadRequest, nil, errorContent)
}

// 未鉴权访问
func ResponseUnAuthorized(c *gin.Context, errorContent string) {
	CommonResponse(c, ResponseCodeUnAuthorized, http.StatusUnauthorized, nil, errorContent)
}

// 访问方法不被允许
func ResponseMethodNotAllowed(c *gin.Context, errorContent string) {
	CommonResponse(c, ResponseCodeMethodNotAllowed, http.StatusMethodNotAllowed, nil, errorContent)
}

// 返回空结果或查询不到结果
func ResponseEmptyResult(c *gin.Context, errorContent string) {
	CommonResponse(c, ResponseCodeEmptyResult, http.StatusOK, nil, errorContent)
}

// handler error response
func ResponseHandlerError(c *gin.Context, errorContent string) {
	CommonResponse(c, ResponseCodeHandlerError, http.StatusInternalServerError, nil, errorContent)
}

// model error response
func ResponseModelError(c *gin.Context, errorContent string) {
	CommonResponse(c, ResponseCodeModelError, http.StatusInternalServerError, nil, errorContent)
}

// repository error response
func ResponseRepositoryError(c *gin.Context, errorContent string) {
	CommonResponse(c, ResponseCodeRepositoryError, http.StatusInternalServerError, nil, errorContent)
}

// cache error response
func ResponseCacheError(c *gin.Context, errorContent string) {
	CommonResponse(c, ResponseCodeCacheError, http.StatusInternalServerError, nil, errorContent)
}

// casbin api error response
func ResponseCasbinApiError(c *gin.Context, errorContent string) {
	CommonResponse(c, ResponseCodeCasbinError, http.StatusForbidden, nil, errorContent)
}

// auth middleware error response
func ResponseAuthMiddlerwareError(c *gin.Context, errorContent string) {
	c.Abort()
	CommonResponse(c, ResponseCodeAuthMiddlerwareError, http.StatusForbidden, nil, errorContent)
}

// cors middleware error response
func ResponseCorsMiddlewareError(c *gin.Context, errorContent string) {
	c.Abort()
	CommonResponse(c, ResponseCodeCorsMiddlerwareError, http.StatusForbidden, nil, errorContent)
}

// rbac middleware error response
func ResponseRbacMiddlewareError(c *gin.Context, errorContent string) {
	c.Abort()
	CommonResponse(c, ResponseCodeRbacMiddlerwareError, http.StatusForbidden, nil, errorContent)
}

package errcode

import (
	"fmt"
	"net/http"
)

var (
	Success       = NewError(0, "成功")
	ServerError   = NewError(100, "服务器内部错误")
	ParamsError   = NewError(101, "参数错误")
	NotFoundError = NewError(102, "未找到结果")
	NoToken       = NewError(103, "缺少token")
	TokenError    = NewError(104, "鉴权失败，token错误")
	TokenTimeout  = NewError(105, "鉴权失败，token超时无效，请重新生成")
	CodeError     = NewError(108, "验证码错误，请重新输入")
)
var codes = map[int]string{}

type Error struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	//Detail	[]string `json:"detail"`
}

func NewError(code int, msg string) *Error {
	if _, ok := codes[code]; ok {
		panic(fmt.Sprintf("错误码%d 已经存在，请更换一个", code))
	}
	codes[code] = msg
	return &Error{Code: code, Msg: msg}
}

func (e *Error) Error() string {
	return fmt.Sprintf("错误码：%d,错误信息：%s", e.Codes(), e.Msgs())
}

func (e *Error) Codes() int {
	return e.Code
}

func (e *Error) Msgs() string {
	return e.Msg
}

func (e *Error) StatusCode() int {
	switch e.Codes() {
	case Success.Codes():
		return http.StatusOK
	case ServerError.Codes():
		return http.StatusInternalServerError
	case NoToken.Codes():
		fallthrough
	case ParamsError.Codes():
		fallthrough
	case NotFoundError.Codes():
		return http.StatusBadRequest
	case TokenError.Codes():
		fallthrough
	case TokenTimeout.Codes():
		fallthrough
	case CodeError.Codes():
		return http.StatusUnauthorized
	}
	return http.StatusInternalServerError
}

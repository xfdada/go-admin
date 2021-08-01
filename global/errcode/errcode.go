package errcode

import (
	"fmt"
	"net/http"
)

var (
	ServerError   = NewError(100, "服务器内部错误")
	ParamsError   = NewError(101, "参数错误")
	NotFoundError = NewError(102, "未找到结果")
	NoToken       = NewError(103, "缺少token")
	TokenError    = NewError(104, "鉴权失败，token错误")
	TokenTimeout  = NewError(105, "鉴权失败，token超时无效，请重新生成")
	AddError      = NewError(106, "添加失败")
	DeleteError   = NewError(107, "删除失败")
	UpdateError   = NewError(108, "更新失败")
	CodeError     = NewError(109, "验证码错误，请重新输入")
	Success       = NewError(200, "成功")
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
	case AddError.Codes():
		fallthrough
	case DeleteError.Codes():
		fallthrough
	case UpdateError.Codes():
		fallthrough
	case CodeError.Codes():
		return http.StatusUnauthorized
	}
	return http.StatusInternalServerError
}
